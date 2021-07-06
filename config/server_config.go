package config

import (
	"fmt"
)

type HTTPServerConfig struct {
	Host             string `mapstructure:"HTTP_SERVER_HOST"`
	Port             string `mapstructure:"HTTP_SERVER_PORT"`
	ReadTimoutInSec  int    `mapstructure:"HTTP_SERVER_READ_TIMEOUT_IN_SEC"`
	WriteTimoutInSec int    `mapstructure:"HTTP_SERVER_WRITE_TIMEOUT_IN_SEC"`
}

func (sc HTTPServerConfig) GetAddress() string {
	return fmt.Sprintf(":%s", sc.Port)
}

func (sc HTTPServerConfig) GetReadTimeout() int {
	return sc.ReadTimoutInSec
}

func (sc HTTPServerConfig) GetWriteTimeout() int {
	return sc.WriteTimoutInSec
}

func InitConfig() HTTPServerConfig {
	return HTTPServerConfig{
		Host:             "localhost",
		Port:             "8080",
		ReadTimoutInSec:  2,
		WriteTimoutInSec: 5,
	}
}
