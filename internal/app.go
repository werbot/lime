package app

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"

	"github.com/werbot/lime/internal/config"
	"github.com/werbot/lime/internal/middleware"
	"github.com/werbot/lime/internal/queries"
	"github.com/werbot/lime/internal/routes"
	"github.com/werbot/lime/migrations"
	"github.com/werbot/lime/pkg/fsutil"
	"github.com/werbot/lime/pkg/logging"
)

// NewApp is ...
func NewApp() error {
	log := logging.Log()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Err(err).Send()
		return err
	}

	// generate keys if need
	if !fsutil.IsDir(config.KeyDir) {
		if err := fsutil.MkDirs(0o775, config.KeyDir); err != nil {
			return err
		}
	}

	if !fsutil.IsFile(config.JWTPrivKeyFile) || !fsutil.IsFile(config.JWTPubKeyFile) {
		GenJWTKeys()
	}

	if !fsutil.IsFile(config.LicensePrivKeyFile) || !fsutil.IsFile(config.LicensePubKeyFile) {
		GenLicenseKeys()
	}

	if err := queries.Init(cfg.Database, migrations.Embed()); err != nil {
		log.Err(err).Send()
		return err
	}

	//var id string
	//db := queries.DB()
	//if err := db.QueryRowContext(context.TODO(), `SELECT id FROM setting`).Scan(&id); err != nil {
	//	fmt.Print(err)
	//}
	//fmt.Print(id)

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	middleware.FiberMiddleware(app, log)

	routes.AdminRoutes(app)
	routes.ApiPrivateRoutes(app)
	routes.ApiPublicRoutes(app)
	routes.NotFoundRoute(app)

	fmt.Print("🍋 Lime - lite license server\n")
	fmt.Printf("├─ Public API: http://%s/api/v1/\n", cfg.HTTPAddr)
	fmt.Printf("└─ Admin UI: http://%s/_/\n", cfg.HTTPAddr)

	if cfg.DevMode {
		StartServer(cfg.HTTPAddr, app)
	} else {
		idleConnsClosed := make(chan struct{})

		go func() {
			sigint := make(chan os.Signal, 1)
			signal.Notify(sigint, os.Interrupt)
			<-sigint

			if err := app.Shutdown(); err != nil {
				log.Err(err).Send()
			}

			close(idleConnsClosed)
		}()

		StartServer(cfg.HTTPAddr, app)
		<-idleConnsClosed
	}

	return nil
}

// StartServer is ...
func StartServer(addr string, a *fiber.App) {
	if err := a.Listen(addr); err != nil {
		log.Err(err).Send()
		os.Exit(1)
	}
}

/*
func useStorage(storage string) (*sql.DB, error) {
	switch storage {
	case "sqlite":
		db, err := sqlite.New(&sqlite.Config{})
		if err != nil {
			return nil, err
		}
		return db, nil
	case "postgres":
		db, err := postgres.New(&postgres.Config{})
		if err != nil {
			return nil, err
		}
		return db, nil
	}

	return nil, errors.New("Invalid storage")
}
*/
