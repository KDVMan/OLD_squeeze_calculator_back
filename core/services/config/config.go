package core_services_config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"time"
)

type ConfigService struct {
	Env         string           `yaml:"env" env-required:"true"`
	StoragePath string           `yaml:"storage_path" env-required:"true"`
	HttpServer  HTTPServerConfig `yaml:"http_server"`
	Logger      LoggerConfig     `yaml:"logger"`
	Binance     BinanceConfig    `yaml:"binance"`
}

type HTTPServerConfig struct {
	Address      string
	Host         string        `yaml:"host" env-required:"true"`
	Port         int           `yaml:"port" env-required:"true"`
	TimeoutRead  time.Duration `yaml:"timeout_read" env-required:"true"`
	TimeoutWrite time.Duration `yaml:"timeout_write" env-required:"true"`
	TimeoutIdle  time.Duration `yaml:"timeout_idle" env-required:"true"`
}

type LoggerConfig struct {
	Level string `yaml:"level" env-required:"true"`
}

type BinanceConfig struct {
	SpotLimit         int     `yaml:"spot_limit" env-required:"true"`
	SpotCommission    float64 `yaml:"spot_commission" env-required:"true"`
	FuturesLimit      int     `yaml:"futures_limit" env-required:"true"`
	FuturesCommission float64 `yaml:"futures_commission" env-required:"true"`
}

func New() *ConfigService {
	var configService ConfigService

	if err := cleanenv.ReadConfig("core/configs/config.yml", &configService); err != nil {
		log.Fatalf("Failed to load config file: %s", err)
	}

	configService.HttpServer.Address = fmt.Sprintf("%s:%d", configService.HttpServer.Host, configService.HttpServer.Port)

	return &configService
}
