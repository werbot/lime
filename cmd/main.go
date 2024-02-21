package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	app "github.com/werbot/lime/internal"
)

var (
	version   = "v0.0.1"
	gitCommit = "00000000"
	buildDate = "24.05.2020"
)

var rootCmd = &cobra.Command{
	Use:                "lime",
	Short:              "Lime CLI",
	Long:               "🍋 Lime - lite license server",
	Version:            fmt.Sprintf("%s (%s) from %s", version, gitCommit, buildDate),
	FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
	CompletionOptions:  cobra.CompletionOptions{DisableDefaultCmd: true},
}

func main() {
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	rootCmd.AddCommand(cmdServe())
	rootCmd.AddCommand(cmdGen())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func cmdServe() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve [flags]",
		Short: "Starts the web server (default to 0.0.0.0:8088 and use sqlite database)",
		Run: func(serveCmd *cobra.Command, args []string) {
			if err := app.NewApp(); err != nil {
				os.Exit(1)
			}
		},
	}

	return cmd
}

func cmdGen() *cobra.Command {
	var configFile, keyJWT, keyLicense bool
	cmd := &cobra.Command{
		Use:   "gen [flags]",
		Short: "Generate keys and config files",
		Run: func(serveCmd *cobra.Command, args []string) {
			if !configFile && !keyJWT && !keyLicense {
				serveCmd.Help()
			}
			if configFile {
				if err := app.GenConfigFile(); err != nil {
					fmt.Print("Config file generated")
					os.Exit(1)
				}
			}
			if keyJWT {
				if err := app.GenJWTKeys(); err != nil {
					fmt.Print("JWT key files generated")
					os.Exit(1)
				}
			}
			if keyLicense {
				if err := app.GenLicenseKeys(); err != nil {
					fmt.Print("Root license key files generated")
					os.Exit(1)
				}
			}
		},
	}

	cmd.PersistentFlags().BoolVar(&configFile, "config", false, "config file")
	cmd.PersistentFlags().BoolVar(&keyJWT, "jwt", false, "jwt keys")
	cmd.PersistentFlags().BoolVar(&keyLicense, "license", false, "license keys")
	return cmd
}
