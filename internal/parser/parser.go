package parser

import (
	"errors"
	"strings"
)

type parsedLine struct {
	key   string
	value string
}

func parseLine(line string) (*parsedLine, error) {
	parts := strings.Split(line, "=")
	if len(parts) != 2 {
		return nil, errors.New("line does not contain a key and value")
	}
	return &parsedLine{
		key:   parts[0],
		value: parts[1],
	}, nil
}
