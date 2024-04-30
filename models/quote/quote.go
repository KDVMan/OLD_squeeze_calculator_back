package models_quote

import (
	"backend/core/services/helper"
	"backend/enums/quote"
	"backend/models"
	"github.com/adshao/go-binance/v2/futures"
	"math"
)

type QuoteModel struct {
	models.BaseModel
	Symbol          string                `gorm:"uniqueIndex:unique_quote;not null" json:"symbol"`
	TimeOpen        int64                 `gorm:"uniqueIndex:unique_quote;not null" json:"timeOpen"`
	TimeClose       int64                 `json:"timeClose"`
	PriceOpen       float64               `json:"priceOpen"`
	PriceHigh       float64               `json:"priceHigh"`
	PriceLow        float64               `json:"priceLow"`
	PriceClose      float64               `json:"priceClose"`
	VolumeLeft      float64               `json:"volumeLeft"`
	VolumeRight     float64               `json:"volumeRight"`
	VolumePrice     float64               `json:"volumePrice"`
	VolumeBuyLeft   float64               `json:"volumeBuyLeft"`
	VolumeBuyRight  float64               `json:"volumeBuyRight"`
	VolumeBuyPrice  float64               `json:"volumeBuyPrice"`
	VolumeSellLeft  float64               `json:"volumeSellLeft"`
	VolumeSellRight float64               `json:"volumeSellRight"`
	VolumeSellPrice float64               `json:"volumeSellPrice"`
	BodySize        float64               `json:"bodySize"`
	StickUpSize     float64               `json:"stickUpSize"`
	StickDownSize   float64               `json:"stickDownSize"`
	StickRatio      float64               `json:"stickRatio"`
	CandleSize      float64               `json:"candleSize"`
	CandleBodyRange float64               `json:"candleBodyRange"`
	Trades          int64                 `json:"trades"`
	Direction       enums_quote.Direction `json:"direction"`
	Percent         QuotePercentModel     `gorm:"embedded" json:"percent"`
	IsClosed        bool                  `json:"isClosed"`
}

func (QuoteModel) TableName() string {
	return "quotes"
}

func KlineToQuote(symbol string, kline *futures.Kline) *QuoteModel {
	priceOpen := core_services_helper.ConvertFloat(kline.Open, 0, 64)
	priceHigh := core_services_helper.ConvertFloat(kline.High, 0, 64)
	priceLow := core_services_helper.ConvertFloat(kline.Low, 0, 64)
	priceClose := core_services_helper.ConvertFloat(kline.Close, 0, 64)
	volumeLeft := core_services_helper.ConvertFloat(kline.Volume, 0, 64)
	volumeRight := core_services_helper.ConvertFloat(kline.QuoteAssetVolume, 0, 64)
	volumeBuyLeft := core_services_helper.ConvertFloat(kline.TakerBuyBaseAssetVolume, 0, 64)
	volumeBuyRight := core_services_helper.ConvertFloat(kline.TakerBuyQuoteAssetVolume, 0, 64)
	volumeSellLeft := volumeLeft - volumeBuyLeft
	volumeSellRight := volumeRight - volumeBuyRight
	bodySize := math.Abs(priceOpen - priceClose)
	stickUpSize := priceHigh - math.Max(priceOpen, priceClose)
	stickDownSize := math.Min(priceOpen, priceClose) - priceLow
	candleSize := math.Abs(priceHigh - priceLow)
	direction := enums_quote.DirectionUp

	if priceOpen > priceClose {
		direction = enums_quote.DirectionDown
	}

	volumePrice := 0.0

	if volumeLeft > 0 {
		volumePrice = core_services_helper.Round(volumeRight/volumeLeft, 8)
	}

	volumeBuyPrice := 0.0

	if volumeBuyLeft > 0 {
		volumeBuyPrice = core_services_helper.Round(volumeBuyRight/volumeBuyLeft, 8)
	}

	volumeSellPrice := 0.0

	if volumeSellLeft > 0 {
		volumeSellPrice = core_services_helper.Round(volumeSellRight/volumeSellLeft, 8)
	}

	stickRatio := 0.0

	if stickDownSize > 0 {
		stickRatio = stickUpSize / stickDownSize
	}

	candleBodyRange := 0.0

	if candleSize > 0 {
		candleBodyRange = bodySize / candleSize
	}

	return &QuoteModel{
		Symbol:          symbol,
		TimeOpen:        kline.OpenTime,
		TimeClose:       kline.CloseTime,
		PriceOpen:       priceOpen,
		PriceHigh:       priceHigh,
		PriceLow:        priceLow,
		PriceClose:      priceClose,
		VolumeLeft:      volumeLeft,
		VolumeRight:     volumeRight,
		VolumePrice:     volumePrice,
		VolumeBuyLeft:   volumeBuyLeft,
		VolumeBuyRight:  volumeBuyRight,
		VolumeBuyPrice:  volumeBuyPrice,
		VolumeSellLeft:  volumeSellLeft,
		VolumeSellRight: volumeSellRight,
		VolumeSellPrice: volumeSellPrice,
		BodySize:        bodySize,
		StickUpSize:     stickUpSize,
		StickDownSize:   stickDownSize,
		StickRatio:      stickRatio,
		CandleSize:      candleSize,
		CandleBodyRange: candleBodyRange,
		Trades:          kline.TradeNum,
		Direction:       direction,
		Percent:         GetPercent(direction, priceOpen, priceHigh, priceLow, priceClose, 2),
		IsClosed:        true,
	}
}
