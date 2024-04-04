package app

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/gofiber/fiber/v2"

	"github.com/werbot/lime/internal/config"
	"github.com/werbot/lime/internal/middleware"
	"github.com/werbot/lime/internal/queries"
	"github.com/werbot/lime/internal/routes"
	"github.com/werbot/lime/migrations"
	"github.com/werbot/lime/pkg/fsutil"
	"github.com/werbot/lime/pkg/jwtutil"
	"github.com/werbot/lime/pkg/logging"
)

var log *logging.Logger

// NewApp is ...
func NewApp() error {
	fmt.Print("🍋 Lime - lite license server\n")

	log = logging.New()

	if err := config.LoadConfig(); err != nil {
		log.Err(err).Send()
		return err
	}
	cfg := config.Data()

	// generate keys if need
	if !fsutil.IsDir(cfg.Keys.KeyDir) {
		if err := fsutil.MkDirs(0o775, cfg.Keys.KeyDir); err != nil {
			return err
		}
	}

	jwtPubKeyFile := filepath.Join(cfg.Keys.KeyDir, cfg.Keys.JWT.PublicKey)
	jwtPrivKeyFile := filepath.Join(cfg.Keys.KeyDir, cfg.Keys.JWT.PrivateKey)
	if !fsutil.IsFile(jwtPrivKeyFile) || !fsutil.IsFile(jwtPubKeyFile) {
		fmt.Print("├─[🔑] A new JWT key pair has been created.\n")
		if err := GenJWTKeys(); err != nil {
			log.Err(err).Send()
			return err
		}
	}

	if err := jwtutil.LoadKeys(jwtPubKeyFile, jwtPrivKeyFile); err != nil {
		log.Err(err).Send()
		return err
	}

	licensePubKeyFile := filepath.Join(cfg.Keys.KeyDir, cfg.Keys.License.PublicKey)
	licensePrivKeyFile := filepath.Join(cfg.Keys.KeyDir, cfg.Keys.License.PrivateKey)
	if !fsutil.IsFile(licensePrivKeyFile) || !fsutil.IsFile(licensePubKeyFile) {
		fmt.Print("├─[🔑] A new License key pair has been created.\n")
		if err := GenLicenseKeys(); err != nil {
			log.Err(err).Send()
			return err
		}
	}

	if cfg.GeoDatabase.Check() {
		fmt.Print("├─[🌏] Downloading the geo database for country identification.\n")
		if err := fsutil.MkDirs(0o775, cfg.GeoDatabase.DBPath); err != nil {
			log.Err(err).Send()
			return err
		}

		if err := cfg.GeoDatabase.Download(); err != nil {
			log.Err(err).Send()
			return err
		}
	}

	if err := queries.Init(cfg.Database, migrations.Embed()); err != nil {
		log.Err(err).Send()
		return err
	}

	/*
		var installed bool
		db := queries.DB()
		if err := db.QueryRowContext(context.TODO(), `SELECT value FROM setting WHERE key = 'installed'`).Scan(&installed); err != nil {
			fmt.Print(err)
		}
		fmt.Print(installed)
	*/

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	middleware.FiberMiddleware(app)

	routes.ApiAdminRoutes(app)
	routes.ApiManagerRoutes(app)
	routes.ApiPublicRoutes(app)
	routes.NotFoundRoute(app)
	routes.UIRoutes(app)

	fmt.Printf("├─[🚀] Admin UI: http://%s/_/\n", cfg.HTTPAddr)
	fmt.Printf("├─[🚀] Public UI: http://%s/\n", cfg.HTTPAddr)
	fmt.Printf("└─[🚀] Public API: http://%s/api/\n", cfg.HTTPAddr)

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
