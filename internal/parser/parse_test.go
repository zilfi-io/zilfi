package parser_test

import (
	"testing"

	"github.com/zilfi-io/zilfi/internal/parser"
)

func TestParseLine(t *testing.T) {
	t.Parallel()

	t.Run("Parse a line with a key and value", func(t *testing.T) {
		t.Parallel()

		line := "KEY=VALUE+=something"

		parsedLine, err := parser.ParseLine(line)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		if parsedLine.Key != "KEY" {
			t.Errorf("Expected KEY, got %s", parsedLine.Key)
		}

		if parsedLine.Value != "VALUE+=something" {
			t.Errorf("Expected VALUE, got %s", parsedLine.Value)
		}
	})
}
