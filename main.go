package main

import (
	"fmt"
	"log"
	"mail-api/email"
	"mail-api/internal/handler"
	"mail-api/logz"
	"mail-api/middleware"
	"mail-api/version"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	_ "time/tzdata"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var isReady bool

func init() {
	runtime.GOMAXPROCS(1)
	initTimezone()
	initViper()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New(fiber.Config{
		StrictRouting: true,
		CaseSensitive: true,
		Immutable:     true,
		ReadTimeout:   viper.GetDuration("app.timeout"),
		WriteTimeout:  viper.GetDuration("app.timeout"),
		IdleTimeout:   viper.GetDuration("app.timeout"),
	})

	logger, err := logz.NewLogConfig()
	if err != nil {
		log.Fatal(err)
	}

	middle := middleware.NewMiddleware(logger)

	e := app.Group(viper.GetString("app.context"))

	e.Use(middle.JSONMiddleware())
	e.Use(middle.ContextLocaleMiddleware())
	e.Use(middle.LoggingMiddleware())

	emailHandler := email.NewMailHandler()

	e.Post("/verification", handler.Helper(emailHandler.SendVerifyEmail, logger))
	e.Post("/otp", handler.Helper(emailHandler.SendOtp, logger))
	e.Post("/margin-call", handler.Helper(emailHandler.SendMarginCall, logger))

	app.Get("/version", version.VersionHandler)
	app.Get("/liveness", func(c *fiber.Ctx) error { return c.Status(http.StatusOK).SendString("Live!") })
	app.Get("/readiness", func(c *fiber.Ctx) error {
		if isReady {
			return c.Status(http.StatusOK).SendString("Read!")
		}
		return c.Status(http.StatusServiceUnavailable).SendString("Unavailable!")
	})

	logger.Info(fmt.Sprintf("⇨ http server started on [::]:%s", viper.GetString("app.port")))

	go func() {
		if err := app.Listen(fmt.Sprintf(":%s", viper.GetString("app.port"))); err != nil {
			logger.Info(err.Error())
		}
	}()

	isReady = true

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	select {
	case <-c:
		logger.Info("terminating: by signal")
	}

	app.Shutdown()

	logger.Info("shutting down")
	os.Exit(0)
}

func initViper() {
	viper.SetDefault("app.name", "mail-api")
	viper.SetDefault("app.port", "8080")
	viper.SetDefault("app.timeout", "60s")
	viper.SetDefault("app.context", "/email")

	viper.SetDefault("log.level", "debug")
	viper.SetDefault("log.env", "dev")

	viper.SetDefault("mail.username", "icfin999@gmail.com")
	viper.SetDefault("mail.smtp-host", "smtp.gmail.com")
	viper.SetDefault("mail.smtp-port", "587")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func initTimezone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Printf("error loading location 'Asia/Bangkok': %v\n", err)
	}
	time.Local = ict
}
