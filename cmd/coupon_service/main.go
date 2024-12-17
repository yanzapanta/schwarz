package main

import (
	"coupon_service/internal/api"
	"coupon_service/internal/config"
	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/service"
	"fmt"
	"time"

	"github.com/gwthm-in/dotenv"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	cfg  = config.New()
	repo = memdb.New()
)

func main() {
	shouldLogDebug := false
	if shouldLogDebug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  "info.log",
		logrus.DebugLevel: "debug.log",
		logrus.ErrorLevel: "error.log",
		logrus.WarnLevel:  "warn.log",
	}

	logrus.AddHook(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))

	logrus.SetFormatter(&logrus.JSONFormatter{})

	// load env variables
	dotenv.OptLookupMod()
	err := dotenv.Load()
	if err != nil {
		logrus.Errorf("error loading env file")
	}

	svc := service.New(repo)
	本 := api.New(cfg.API, svc)
	本.Start()
	fmt.Println("Starting Coupon service server")
	<-time.After(1 * time.Hour * 24 * 365)
	fmt.Println("Coupon service server alive for a year, closing")
	本.Close()
}
