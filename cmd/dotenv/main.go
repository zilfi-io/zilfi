package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/zilfi-io/zilfi/internal/parser"
)

func createRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "envy",
	}

	return rootCmd
}

func createHelloCmd() *cobra.Command {
	helloCmd := &cobra.Command{
		Use: "hello",
		Run: func(cmd *cobra.Command, _ []string) {
			flags := cmd.Flags()
			printFlags(*flags)
		},
	}
	helloCmd.Flags().StringP("env", "e", "dev", "Your environment")
	helloCmd.Flags().StringP("name", "n", "world", "Your name")

	return helloCmd
}

func printFlags(flags pflag.FlagSet) {
	flags.VisitAll(func(flag *pflag.Flag) {
		log.Printf("%s: %s\n", flag.Name, flag.Value)
	})
}

func main() {
	rootCmd := createRootCmd()
	helloCmd := createHelloCmd()
	rootCmd.AddCommand(helloCmd)

	parseEnvCmd := parser.CreateParseEnvCmd()
	rootCmd.AddCommand(parseEnvCmd)

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
