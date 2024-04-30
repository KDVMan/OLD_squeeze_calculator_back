package services_calculator

import (
	"backend/core/models"
	"backend/core/services/app"
	"backend/core/services/logger"
	"backend/enums"
	"backend/models/control_symbol"
	"backend/models/quote"
	"backend/models/result_calculator"
	"backend/models/symbol"
	"backend/requests/control_calculator"
	"backend/services/calculator_optimization"
	"backend/services/calculator_score"
	"backend/variables/calculator"
	"log"
	"log/slog"
	"sync"
	"time"
)

func Start(appService *core_services_app.AppService, logger *slog.Logger, controlCalculatorModel *models_control_calculator.ControlCalculatorModel,
	symbolModel *models_symbol.SymbolModel, request *requests_control_calculator.StartRequest) {
	var results []*models_result_calculator.ResultCalculatorModel

	calculatorOptimizationService := services_calculator_optimization.New(controlCalculatorModel)
	optimizations := calculatorOptimizationService.Load()

	quoteRange := models_quote.GetRange(int64(appService.ConfigService.Binance.FuturesLimit), request.TimeFrom, request.TimeTo, enums.IntervalMilliseconds(enums.Interval1m))

	progress := &core_models.ProgressChannelModel{
		Count:  0,
		Total:  int64(quoteRange.Iterations+len(optimizations)) + 1,
		Status: enums.WebsocketStatusProgress,
	}

	appService.WebsocketService.ProgressChan <- progress
	quotes, err := appService.QuoteService.LoadRange(request.Symbol, quoteRange, appService.WebsocketService.ProgressChan, progress)
	startTime := time.Now()

	numWorkers := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	taskChan := make(chan *models_result_calculator.ResultCalculatorParam, len(optimizations))
	resultChan := make(chan *models_result_calculator.ResultCalculatorModel, len(optimizations))
	progressChan := make(chan int64, len(optimizations))
	var completedTasks int64 = 0

	for i := 0; i < numWorkers; i++ {
		go worker(taskChan, resultChan, quotes, symbolModel.Limit.TickSize, appService.ConfigService.Binance.FuturesCommission, &wg, appService, progress, progressChan)
	}

	go func() {
		for p := range progressChan {
			completedTasks += p
			progress.Count = completedTasks

			if progress.Count%1000 == 0 {
				appService.WebsocketService.ProgressChan <- progress
			}
		}
	}()

	for _, optimization := range optimizations {
		if variables_calculator.Stop {
			break
		}

		if variables_calculator.Stop {
			progress.Status = enums.WebsocketStatusStop
			appService.WebsocketService.ProgressChan <- progress

			return
		}

		param := &models_result_calculator.ResultCalculatorParam{
			Symbol:         controlCalculatorModel.Symbol,
			TradeDirection: controlCalculatorModel.TradeDirection,
			Interval:       controlCalculatorModel.Interval,
			Bind:           optimization.Bind,
			PercentIn:      optimization.PercentIn,
			PercentOut:     optimization.PercentOut,
			StopTime:       optimization.StopTime,
			StopPercent:    optimization.StopPercent,
			OncePerCandle:  controlCalculatorModel.OncePerCandle,
		}

		taskChan <- param
	}

	close(taskChan)
	wg.Wait()
	close(resultChan)
	close(progressChan)

	for result := range resultChan {
		results = append(results, result)
	}

	log.Printf("Время выполнения кода: %s\n", time.Since(startTime))

	results = services_calculator_score.Results(results)

	if err = appService.ResultCalculatorService.Save(results); err != nil {
		message := "Failed to save result calculator"
		logger.Error(message, core_services_logger.Err(err))
	}

	progress.Count++
	progress.Status = enums.WebsocketStatusDone
	appService.WebsocketService.ProgressChan <- progress
}

func worker(taskChan <-chan *models_result_calculator.ResultCalculatorParam, resultChan chan<- *models_result_calculator.ResultCalculatorModel, quotes []*models_quote.QuoteModel,
	tickSize, commission float64, wg *sync.WaitGroup, appService *core_services_app.AppService, progress *core_models.ProgressChannelModel, progressChan chan<- int64) {
	defer wg.Done()

	for param := range taskChan {
		if variables_calculator.Stop {
			progress.Status = enums.WebsocketStatusStop
			appService.WebsocketService.ProgressChan <- progress
			return
		}

		calculatorService := New(param, quotes, tickSize, commission)
		result := calculatorService.Calculate()
		resultChan <- result
		progressChan <- 1
	}
}
