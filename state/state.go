package state

import (
	"cscope_lsp/cscope_if"
	"cscope_lsp/lsp"
	"log"
	"strings"
)

type State struct {
	// Map of file names to contents
	Documents map[string]string
}

func New() State {
	return State{Documents: map[string]string{}}
}

func (s *State) Update(uri, text string) {
	s.Documents[uri] = text
}

func isWordDelim(c byte) bool {
	return c == ' ' || c == '(' || c == ')' || c == '*' || c == '-' || c == '>' || c == '<' || c == ',' || c == '.'
}

func (s *State) Definition(id int, uri string, logger *log.Logger, position lsp.Position) lsp.DefinitionResponse {
	logger.Printf("URI: %s", uri)
	logger.Printf("Position.line: %d", position.Line)
	logger.Printf("Position.char: %d", position.Character)
	line := strings.Split(s.Documents[uri], "\n")[position.Line]

	start := position.Character
	end := position.Character

	for start > 0 && !isWordDelim(line[start]) {
		start--
	}

	if isWordDelim(line[start]) {
		start++
	}

	for end < len(line) && !isWordDelim(line[end]) {
		end++
	}

	logger.Printf("line: %s", line)
	logger.Printf("word: %s", line[start:end])
	defs := cscope_if.GetDefinition(id, logger, uri, line[start:end])

	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: defs,
	}
}
