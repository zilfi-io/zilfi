package parser

import "testing"

func TestParseLine(t *testing.T) {
	t.Run("Parse a line with a key and value", func(t *testing.T) {
		line := "KEY=VALUE"
		parsedLine, err := parseLine(line)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if parsedLine.key != "KEY" {
			t.Errorf("Expected KEY, got %s", parsedLine.key)
		}
		if parsedLine.value != "VALUE" {
			t.Errorf("Expected VALUE, got %s", parsedLine.value)
		}
	})
}
