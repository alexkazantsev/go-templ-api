package cmd

import (
	"github.com/alexkazantsev/go-templ-api/modules/config"
	"github.com/alexkazantsev/go-templ-api/modules/core"
	"github.com/alexkazantsev/go-templ-api/modules/database"
	"github.com/alexkazantsev/go-templ-api/modules/user"
	"github.com/alexkazantsev/go-templ-api/pkg/logger"
	"github.com/alexkazantsev/go-templ-api/server"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func Run() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "run api gateway",
		Run: func(cmd *cobra.Command, args []string) {
			fx.New(
				fx.Provide(logger.NewLogger),

				config.Module,
				database.Module,
				server.Module,

				core.Module,
				user.Module,

				logger.WithZapLogger,
			).Run()
		},
	}
}
