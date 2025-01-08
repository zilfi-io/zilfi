package parser

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	ErrInvalidKeyValuePair = errors.New("invalid key value pair")
	ErrEmptyInput          = errors.New("input is empty")
	ErrMissingDelimiter    = errors.New("missing delimiter")
)

type ParsedLine struct {
	Key   string
	Value string
}

// Should read one line at a time.
func getFileContents(location string) (string, error) {
	log.Printf("Parsing %s\n", location)

	file, err := os.Open(location)
	if err != nil {
		log.Println("error opening file: ", err)
		err = fmt.Errorf("error opening file: %w", err)

		return "", err
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Println("error reading file: ", err)
		err = fmt.Errorf("error reading file: %w", err)

		return "", err
	}

	return string(content), nil
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

func ParseLine(line string) (*ParsedLine, error) {
	// Check if the line contains a key value pair
	if line == "" {
		return nil, ErrEmptyInput
	}

	if !strings.Contains(line, "=") {
		return nil, ErrMissingDelimiter
	}

	// Return only 2 substrings, so = can be a valid part of the value
	SubstringsToReturn := 2

	parts := strings.SplitN(line, "=", SubstringsToReturn)

	if len(parts) != SubstringsToReturn {
		return nil, ErrInvalidKeyValuePair
	}

	// Trim the quotes from the key
	key := strings.Trim(parts[0], "\"")

	// Concatenate the remaining parts of the line and then trim the quotes
	value := strings.Trim(parts[1], "\"")

	return &ParsedLine{
		Key:   key,
		Value: value,
	}, nil
}

func writeKeyValuePairs(file *os.File, lines []string) error {
	var completed int
	// Find the number of lines that are not empty
	linesWithContent := countLinesWithContent(lines)

	for _, line := range lines {
		if line == "" {
			continue
		}

		parsedLine, err := ParseLine(line)
		if err != nil {
			log.Println("Error parsing line: ", err)

			continue
		}

		key := parsedLine.Key
		value := parsedLine.Value

		if completed == linesWithContent-1 {
			_, err := file.WriteString(fmt.Sprintf("  \"%s\": \"%s\"\n", key, value))
			if err != nil {
				err = fmt.Errorf("error writing to file: %w", err)
				log.Println("Error writing to file: ", err)

				return err
			}

			break
		}

		_, err = file.WriteString(fmt.Sprintf("  \"%s\": \"%s\",\n", key, value))
		if err != nil {
			err = fmt.Errorf("error writing to file: %w", err)
			log.Println("Error writing to file: ", err)

			return err
		}

		completed++
	}

	return nil
}

func writeEnv(contents string) {
	log.Println("Writing to .env")

	file, err := os.Create(".env.json")
	if err != nil {
		log.Fatalln("Error creating file: ", err)
	}

	// Create a JSON file where each line is a key value pair
	lines := strings.Split(contents, "\n")

	// Start the JSON object
	_, err = file.WriteString("{\n")
	if err != nil {
		file.Close()
		log.Fatalln("Error writing starting curly braces to file: ", err)
	}

	// Write the key value pairs
	err = writeKeyValuePairs(file, lines)
	if err != nil {
		file.Close()
		log.Fatalln("Error writing key vlue pairs to file: ", err)
	}

	// End the JSON object
	_, err = file.WriteString("}\n")
	if err != nil {
		file.Close()
		log.Fatalln("Error writing end curly braces to file: ", err)
	}

	file.Close()
}

func CreateParseEnvCmd() *cobra.Command {
	parseEnvCmd := &cobra.Command{
		Use:   "import",
		Short: "Import a .env file",
		Long:  "Import a .env file into your environment",
		Run: func(cmd *cobra.Command, _ []string) {
			// Grab a .env file and parse it
			location, _ := cmd.Flags().GetString("location")

			// OPTIMIZE: Let's read this one line at a time
			// scanner := bufio.NewScanner(file)
			// for scanner.Scan() {
			//   line := scanner.Text()
			// }
			envContents, err := getFileContents(location)
			if err != nil {
				log.Println("Error parsing file: ", err)
				os.Exit(1)
			}
			writeEnv(envContents)
		},
	}

	parseEnvCmd.Flags().StringP("location", "l", ".env", "Location of the .env file")
	parseEnvCmd.Flags().StringP("fileName", "f", ".env.json", "The name of the file to write to")

	return parseEnvCmd
}
