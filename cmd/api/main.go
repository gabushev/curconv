package main

import (
	"curconv/cmd/api/migrations"
	"curconv/config"
	_ "curconv/docs"
	"curconv/internal/handlers"
	"curconv/internal/models"
	"curconv/internal/repository"
	"curconv/internal/services"
	"curconv/internal/updater"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

//	@title			Currency exchange API
//	@version		1.0
//	@description	This is a sample server for a currency exchange API.
//	@termsOfService	http://swagger.io/terms/

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:3000
// @BasePath	/
func main() {
	cfg := config.LoadConfig()
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	dsn := "host=" + cfg.DbHost + " user=" + cfg.DbUser + " password=" + cfg.DbPassword + " dbname=" + cfg.DbName + " port=" + cfg.DbPort + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	if err := models.Migrate(db); err != nil {
		log.Fatal().Err(err).Msg("Failed to migrate database")
	}
	migrations.SeedCurrencyPairs(db)

	forexClient := services.NewAPIClient(cfg.ForexAPI, cfg.ForexApiKey)
	currencyRepo := repository.NewCurrencyRepository(db)
	rateUpdater := updater.NewRateUpdater(currencyRepo, &forexClient)
	rateUpdater.Start()

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger,
	}))
	exchangeHandler := handlers.NewExchangeHandler(rateUpdater)
	app.Get("/convert", exchangeHandler.Convert)
	// Swagger route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	log.Info().Msg("Starting server on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
