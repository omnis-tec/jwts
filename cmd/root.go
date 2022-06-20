package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	dopLoggerZap "github.com/rendau/dop/adapters/logger/zap"
	dopServerHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/dop/dopTools"
	"github.com/rendau/jwts/internal/adapters/server/rest"
	"github.com/rendau/jwts/internal/domain/core"
)

func Execute() {
	var err error

	app := struct {
		lg         *dopLoggerZap.St
		core       *core.St
		restApiSrv *dopServerHttps.St
	}{}

	confLoad()

	app.lg = dopLoggerZap.New(conf.LogLevel, conf.Debug)

	app.core = core.New(app.lg)

	if conf.Kid != "" {
		var privatePem []byte
		var publicPem []byte

		if privatePemPath := conf.PrivatePem; privatePemPath != "" {
			privatePem, err = ioutil.ReadFile(privatePemPath)
			if err != nil {
				log.Fatal(err)
			}
		}

		if publicPemPath := conf.PublicPem; publicPemPath != "" {
			publicPem, err = ioutil.ReadFile(publicPemPath)
			if err != nil {
				log.Fatal(err)
			}
		}

		// set keys
		err = app.core.SetKeys(privatePem, publicPem, conf.Kid)
		if err != nil {
			log.Fatal(err)
		}
	}

	// START

	app.lg.Infow("Starting")

	app.restApiSrv = dopServerHttps.Start(
		conf.HttpListen,
		rest.GetHandler(
			app.lg,
			app.core,
			conf.HttpCors,
		),
		app.lg,
	)

	var exitCode int

	select {
	case <-dopTools.StopSignal():
	case <-app.restApiSrv.Wait():
		exitCode = 1
	}

	// STOP

	app.lg.Infow("Shutting down...")

	if !app.restApiSrv.Shutdown(20 * time.Second) {
		exitCode = 1
	}

	app.lg.Infow("Exit")

	os.Exit(exitCode)
}
