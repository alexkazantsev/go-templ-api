package cmd

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

func MigrationCreate() *cobra.Command {
	var (
		name string
		dir  = "./modules/database/migrations"
	)

	cmd := &cobra.Command{
		Use: "migration:create",
		Run: func(cmd *cobra.Command, args []string) {
			var (
				err error
				out []byte
			)

			if !isToolAvailable() {
				log.Println("migrate command does not exist!")
				if out, err = exec.
					Command("go", "install", "-v", "-tags", "postgres", "github.com/golang-migrate/migrate/v4/cmd/migrate@latest").
					CombinedOutput(); err != nil {
					log.Fatalf("failed to install migrate tool: %v", err)
				}

				log.Println(string(out))
			}

			if out, err = exec.
				Command("migrate", "create", "-ext", "sql", "-dir", dir, "-seq", name).
				CombinedOutput(); err != nil {
				log.Fatalf("failed to create migration: %v", err)
			}

			log.Println(string(out))
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "name of migration file")
	_ = cmd.MarkFlagRequired("name")

	return cmd
}

func isToolAvailable() bool {
	_, err := exec.LookPath("migrate")
	return err == nil
}
