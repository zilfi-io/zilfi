package main

import "testing"

func TestParseLine(t *testing.T) {
	t.Run("Parse a line with a key and value", func(t *testing.T) {
		line := "KEY=VALUE"
		key, value := parseLine(line)
		if key != "KEY" {
			t.Errorf("Expected KEY, got %s", key)
		}
		if value != "VALUE" {
			t.Errorf("Expected VALUE, got %s", value)
		}
	})
}
