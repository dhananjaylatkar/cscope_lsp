package state

import (
	"github.com/dhananjaylatkar/cscope_lsp/cscope_if"
	"github.com/dhananjaylatkar/cscope_lsp/lsp"
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

func isWordChar(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_'
}

func extractWord(line string, pos int) string {
	start := pos
	end := pos

	for start > 0 && isWordChar(line[start]) {
		start--
	}

	if !isWordChar(line[start]) {
		start++
	}

	for end < len(line) && isWordChar(line[end]) {
		end++
	}

	if start >= end {
		return ""
	}

	return line[start:end]
}

func (s *State) Definition(id int, uri string, logger *log.Logger, position lsp.Position) lsp.DefinitionResponse {
	logger.Printf("uri: %s", uri)
	logger.Printf("position.Line: %d", position.Line)
	logger.Printf("position.Char: %d", position.Character)
	line := strings.Split(s.Documents[uri], "\n")[position.Line]
	word := extractWord(line, position.Character)

	logger.Printf("line: %s", line)
	logger.Printf("word: %s", word)

	defs := cscope_if.GetDefinition(logger, uri, word)

	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: defs,
	}
}

func (s *State) References(id int, uri string, logger *log.Logger, position lsp.Position) lsp.ReferencesResponse {
	logger.Printf("uri: %s", uri)
	logger.Printf("position.Line: %d", position.Line)
	logger.Printf("position.Char: %d", position.Character)

	line := strings.Split(s.Documents[uri], "\n")[position.Line]
	word := extractWord(line, position.Character)
	logger.Printf("line: %s", line)
	logger.Printf("word: %s", word)

	defs := cscope_if.GetReferences(logger, uri, word)

	return lsp.ReferencesResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: defs,
	}
}
