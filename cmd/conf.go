package cmd

import (
	"github.com/rendau/dop/dopTools"
	"github.com/spf13/viper"
)

var conf = struct {
	Debug      bool   `mapstructure:"DEBUG"`
	LogLevel   string `mapstructure:"LOG_LEVEL"`
	HttpListen string `mapstructure:"HTTP_LISTEN"`
	HttpCors   bool   `mapstructure:"HTTP_CORS"`
	Kid        string `mapstructure:"KID"`
	PrivatePem string `mapstructure:"PRIVATE_PEM"`
	PublicPem  string `mapstructure:"PUBLIC_PEM"`
}{}

func confLoad() {
	dopTools.SetViperDefaultsFromObj(conf)

	viper.SetDefault("DEBUG", "false")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("HTTP_LISTEN", ":80")

	viper.SetConfigFile("conf.yml")
	_ = viper.ReadInConfig()

	viper.AutomaticEnv()

	_ = viper.Unmarshal(&conf)
}
