package main

import (
	"fmt"
	"log"
	"mail-api/docs"
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

	swagger "github.com/arsmn/fiber-swagger/v2"
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

// @title Mail API Services
// @version 1.0
// @description Mail API for ICFIN company.
// @termsOfService http://swagger.io/terms/
// @contact.name K.Apiwattanawong
// @contact.url http://www.swagger.io/support
// @contact.email k.apiwattanawong@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9090
// @BasePath /email
// @schemes http https
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

	swag := app.Group("/swagger")
	swag.Use(middle.BasicAuthenicationMiddleware())
	registerSwaggerRoute(swag)

	e := app.Group(viper.GetString("app.context"))

	e.Use(middle.JSONMiddleware())
	e.Use(middle.ContextLocaleMiddleware())
	e.Use(middle.LoggingMiddleware())

	emailHandler := email.NewMailHandler()

	e.Post("/verification", handler.Helper(emailHandler.SendVerifyEmail, logger))
	e.Post("/otp", handler.Helper(emailHandler.SendOtp, logger))
	e.Post("/margin-call", handler.Helper(emailHandler.SendMarginCall, logger))
	e.Post("/liquidation", handler.Helper(emailHandler.SendLiquidateFund, logger))

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

	viper.SetDefault("swagger.host", "localhost:8080")
	viper.SetDefault("swagger.user", "admin")
	viper.SetDefault("swagger.password", "password")

	viper.SetDefault("log.level", "debug")
	viper.SetDefault("log.env", "dev")

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

func registerSwaggerRoute(swag fiber.Router) {
	swag.Get("/*", swagger.New(swagger.Config{
		URL:         fmt.Sprintf("http://%s/swagger/doc.json", viper.GetString("swagger.host")),
		DeepLinking: false,
	}))
	docs.SwaggerInfo.Host = viper.GetString("swagger.host")
	docs.SwaggerInfo.BasePath = viper.GetString("app.context")
}
