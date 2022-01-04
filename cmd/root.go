package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mechta-market/jwts/internal/adapters/httpapi/rest"
	"github.com/mechta-market/jwts/internal/adapters/logger/zap"
	"github.com/mechta-market/jwts/internal/domain/core"
	"github.com/mechta-market/jwts/internal/domain/usecases"
	"github.com/spf13/viper"
)

func Execute() {
	var err error

	loadConf()

	debug := viper.GetBool("DEBUG")

	app := struct {
		lg      *zap.St
		core    *core.St
		ucs     *usecases.St
		restApi *rest.St
	}{}

	app.lg, err = zap.New(viper.GetString("LOG_LEVEL"), debug, false)
	if err != nil {
		log.Fatal(err)
	}

	app.core = core.New(app.lg)

	if kid := viper.GetString("KID"); kid != "" {
		var privatePem []byte
		var publicPem []byte

		if privatePemPath := viper.GetString("PRIVATE_PEM"); privatePemPath != "" {
			privatePem, err = ioutil.ReadFile(privatePemPath)
			if err != nil {
				log.Fatal(err)
			}
		}

		if publicPemPath := viper.GetString("PUBLIC_PEM"); publicPemPath != "" {
			publicPem, err = ioutil.ReadFile(publicPemPath)
			if err != nil {
				log.Fatal(err)
			}
		}

		// set keys
		err = app.core.SetKeys(privatePem, publicPem, kid)
		if err != nil {
			log.Fatal(err)
		}
	}

	app.ucs = usecases.New(
		app.lg,
		app.core,
	)

	restApiEChan := make(chan error, 1)

	app.restApi = rest.New(
		app.lg,
		viper.GetString("HTTP_LISTEN"),
		app.ucs,
		restApiEChan,
	)

	app.lg.Infow("Starting")

	app.lg.Infow("http_listen " + viper.GetString("HTTP_LISTEN"))

	app.restApi.Start()

	stopSignalChan := make(chan os.Signal, 1)
	signal.Notify(stopSignalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	var exitCode int

	select {
	case <-stopSignalChan:
	case <-restApiEChan:
		exitCode = 1
	}

	app.lg.Infow("Shutting down...")

	err = app.restApi.Shutdown(20 * time.Second)
	if err != nil {
		app.lg.Errorw("Fail to shutdown http-api", err)
		exitCode = 1
	}

	app.lg.Infow("Exit")

	os.Exit(exitCode)
}

func loadConf() {
	viper.SetDefault("DEBUG", "false")
	viper.SetDefault("HTTP_LISTEN", ":80")
	viper.SetDefault("LOG_LEVEL", "debug")

	confFilePath := os.Getenv("CONF_PATH")
	if confFilePath == "" {
		confFilePath = "conf.yml"
	}
	viper.SetConfigFile(confFilePath)
	_ = viper.ReadInConfig()

	viper.AutomaticEnv()
}
