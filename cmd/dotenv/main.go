package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use: "envy",
}

var initCmd = &cobra.Command{
	Use: "init",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented")
	},
}

var parseEnvCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a .env file",
	Long:  "Import a .env file into your environment",
	Run: func(cmd *cobra.Command, args []string) {
		// Grab a .env file and parse it
		location, _ := cmd.Flags().GetString("location")
		parseFile(location)
	},
}

func parseFile(location string) {
	fmt.Printf("Parsing %s\n", location)
}

var helloCmd = &cobra.Command{
	Use: "hello",
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		printFlags(*flags)
	},
}

func printFlags(flags pflag.FlagSet) {
	flags.VisitAll(func(flag *pflag.Flag) {
		fmt.Printf("%s: %s\n", flag.Name, flag.Value)
	})
}

func main() {
	helloCmd.Flags().StringP("env", "e", "dev", "Your environment")
	helloCmd.Flags().StringP("name", "n", "world", "Your name")

	parseEnvCmd.Flags().StringP("location", "l", ".env", "Location of the .env file")

	rootCmd.AddCommand(helloCmd)
	rootCmd.AddCommand(parseEnvCmd)
	rootCmd.Execute()
}
