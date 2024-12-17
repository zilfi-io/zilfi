package main

import (
	"fmt"
	"io"
	"os"
	"strings"

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
		envContents, err := parseFile(location)
		if err != nil {
			fmt.Println("Error parsing file: ", err)
			os.Exit(1)
		}
		writeEnv(envContents)
	},
}

func countLinesWithContent(lines []string) int {
	count := 0
	for _, line := range lines {
		if line != "" {
			count++
		}
	}
	return count
}

func writeEnv(contents string) {
	fmt.Println("Writing to .env")
	file, err := os.Create(".env.json")
	if err != nil {
		fmt.Println("Error creating file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	// Create a JSON file where each line is a key value pair
	lines := strings.Split(contents, "\n")
	var completed int
	// Find the number of lines that are not empty
	linesWithContent := countLinesWithContent(lines)

	file.WriteString("{\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) < 2 {
			continue
		}

		key := strings.Trim(parts[0], "\"")
		value := strings.Trim(parts[1], "\"")

		if completed == linesWithContent-1 {
			file.WriteString(fmt.Sprintf("  \"%s\": \"%s\"\n", key, value))
			break
		}
		file.WriteString(fmt.Sprintf("  \"%s\": \"%s\",\n", key, value))
		completed++
	}
	file.WriteString("}\n")
}

func parseFile(location string) (string, error) {
	fmt.Printf("Parsing %s\n", location)
	file, err := os.Open(location)
	if err != nil {
		fmt.Println("error opening file: ", err)
		return "", err
	}
	defer file.Close()

	// Read the file
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("error reading file: ", err)
		return "", err
	}
	return string(content), nil
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
	parseEnvCmd.Flags().StringP("fileName", "f", ".env.json", "The name of the file to write to")

	rootCmd.AddCommand(helloCmd)
	rootCmd.AddCommand(parseEnvCmd)
	rootCmd.Execute()
}
