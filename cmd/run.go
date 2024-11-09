package cmd

import (
	"github.com/alexkazantsev/templ-api/modules/config"
	"github.com/alexkazantsev/templ-api/modules/database"
	"github.com/alexkazantsev/templ-api/modules/user"
	"github.com/alexkazantsev/templ-api/pkg/logger"
	"github.com/alexkazantsev/templ-api/server"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func Run() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "run api gateway",
		Run: func(cmd *cobra.Command, args []string) {
			fx.New(
				config.Module,
				logger.Module,
				database.Module,
				server.Module,

				user.Module,

				logger.WithZapLogger,
			).Run()
		},
	}
}
