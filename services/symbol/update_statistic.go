package services_symbol

import (
	"backend/core/models"
	"backend/core/services/helper"
	"backend/enums"
	"backend/models/symbol"
	"errors"
	"fmt"
	"github.com/adshao/go-binance/v2/futures"
	"gorm.io/gorm"
)

func (symbolService *SymbolService) UpdateStatistic(tickets []*futures.WsMarketTickerEvent) error {
	const label = "services.symbol.UpdateStatistic"

	err := symbolService.storageService.DB.Transaction(func(tx *gorm.DB) error {
		for _, ticket := range tickets {
			var symbolModel models_symbol.SymbolModel

			if err := tx.Where("symbol = ?", ticket.Symbol).First(&symbolModel).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					continue
				}

				return fmt.Errorf("%s: %w", label, err)
			}

			if err := tx.Model(&symbolModel).Updates(models_symbol.SymbolStatisticModel{
				Price:        core_services_helper.ConvertFloat(ticket.ClosePrice, 0, 64),
				PriceLow:     core_services_helper.ConvertFloat(ticket.LowPrice, 0, 64),
				PriceHigh:    core_services_helper.ConvertFloat(ticket.HighPrice, 0, 64),
				PricePercent: core_services_helper.ConvertFloat(ticket.PriceChangePercent, 0, 64),
				Volume:       core_services_helper.ConvertFloat(ticket.QuoteVolume, 0, 64),
				Trades:       ticket.TradeCount,
			}).Error; err != nil {
				return fmt.Errorf("%s: %w", label, err)
			}
		}

		return nil
	})

	if err == nil {
		symbols, err := symbolService.LoadAll()

		if err != nil {
			return fmt.Errorf("%s: %w", label, err)
		}

		broadcastModel := core_models.BroadcastChannelModel{
			Event: enums.WebsocketEventSymbolCalculatorSymbols,
			Data:  symbols,
		}

		symbolService.broadcastChan <- &broadcastModel
	}

	return err
}
