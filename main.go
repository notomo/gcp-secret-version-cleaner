package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/notomo/gcp-secret-version-cleaner/app"
	"github.com/urfave/cli/v2"
)

const (
	paramProjectName = "project"
	paramSecretName  = "secret-name"
	paramLogDir      = "log-dir"

	paramFilter     = "filter"
	paramDryRun     = "dry-run"
	paramKeepRecent = "keep-recent-count"
)

func main() {

	app := &cli.App{
		Name: "gcp-secret-version-cleaner",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     paramProjectName,
				Required: true,
				Usage:    "project name",
			},
			&cli.StringFlag{
				Name:     paramSecretName,
				Required: true,
				Usage:    "secret name",
			},
			&cli.StringFlag{
				Name:     paramLogDir,
				Required: false,
				Usage:    "log directory (output log if not empty)",
			},
		},

		Commands: cli.Commands{
			{
				Name: "destroy",
				Action: func(c *cli.Context) error {
					baseTransport := app.LogTransport(c.String(paramLogDir), http.DefaultTransport)

					return app.Destroy(
						c.Context,
						c.String(paramProjectName),
						c.String(paramSecretName),
						c.String(paramFilter),
						c.Uint(paramKeepRecent),
						baseTransport,
						c.Bool(paramDryRun),
					)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     paramFilter,
						Required: false,
						Value:    "state:DISABLED",
						Usage:    "secret version filter : https://cloud.google.com/secret-manager/docs/filtering ",
					},
					&cli.UintFlag{
						Name:     paramKeepRecent,
						Required: false,
						Value:    0,
						Usage:    "keep recent count (applied after filter option)",
					},
					&cli.BoolFlag{
						Name:     paramDryRun,
						Required: false,
						Value:    false,
						Usage:    "dry run",
					},
				},
			},

			{
				Name: "disable",
				Action: func(c *cli.Context) error {
					baseTransport := app.LogTransport(c.String(paramLogDir), http.DefaultTransport)

					return app.Disable(
						c.Context,
						c.String(paramProjectName),
						c.String(paramSecretName),
						c.String(paramFilter),
						c.Uint(paramKeepRecent),
						baseTransport,
						c.Bool(paramDryRun),
					)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     paramFilter,
						Required: false,
						Value:    "state:ENABLED",
						Usage:    "secret version filter : https://cloud.google.com/secret-manager/docs/filtering ",
					},
					&cli.UintFlag{
						Name:     paramKeepRecent,
						Required: false,
						Value:    0,
						Usage:    "keep recent count (applied after filter option)",
					},
					&cli.BoolFlag{
						Name:     paramDryRun,
						Required: false,
						Value:    false,
						Usage:    "dry run",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		slog.Default().Error(err.Error())
		os.Exit(1)
	}
}
