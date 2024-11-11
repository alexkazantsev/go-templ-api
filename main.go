package main

import (
	"flag"
	"log"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"

	"github.com/alexkazantsev/go-templ-api/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func main() {
	var root = &cobra.Command{Use: "root"}

	root.AddCommand(cmd.Run())

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	if err := root.Execute(); err != nil {
		log.Fatalf("fatal gateway execute error: %v", err)
	}
}
