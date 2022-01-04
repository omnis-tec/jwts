package tests

import (
	_ "embed"
	"log"
	"os"
	"testing"

	"github.com/mechta-market/jwts/internal/adapters/logger/zap"
	"github.com/mechta-market/jwts/internal/domain/core"
	"github.com/mechta-market/jwts/internal/domain/usecases"
	"github.com/spf13/viper"
)

//go:embed private.pem
var privatePem []byte

//go:embed public.pem
var publicPem []byte

func TestMain(m *testing.M) {
	var err error

	viper.SetConfigFile("test_conf.yml")
	_ = viper.ReadInConfig()

	viper.AutomaticEnv()

	app.lg, err = zap.New(
		"info",
		true,
		false,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer app.lg.Sync()

	app.core = core.New(app.lg)

	err = app.core.SetKeys(privatePem, publicPem, "key1")
	if err != nil {
		log.Fatal(err)
	}

	app.ucs = usecases.New(
		app.lg,
		app.core,
	)

	// Start tests
	code := m.Run()

	os.Exit(code)
}
