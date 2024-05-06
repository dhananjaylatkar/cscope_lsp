package cscope_if

import (
	"cscope_lsp/lsp"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	refs = "-0"
	defs  = "-1"
)

func findProjRoot(uri string) string {
	path := uri[7:] // trim "file://" from uri

	for path != "/" {
		path = filepath.Dir(path)

		if _, err := os.Stat(filepath.Join(path, "cscope.out")); err == nil {
			return path
		}
	}

	return ""
}

func getResults(logger *log.Logger, uri string, sym string, op string) []lsp.Location {
	var res []lsp.Location

	projRoot := findProjRoot(uri)
	cs_db := filepath.Join(projRoot, "cscope.out")

	logger.Printf("projRoot: %s", projRoot)
	logger.Printf("op: %s", op)

	out, err := exec.Command("cscope", "-dL", "-f", cs_db, op, sym).Output()

	if err != nil {
		logger.Printf("Err: %s", err)
		return res
	}

	cs_res := strings.Split(string(out), "\n")

	for _, r := range cs_res {
		if r == "" {
			continue
		}

		logger.Printf("%s", r)
		sp := strings.Split(r, " ")
		fname := sp[0]
		lnum, _ := strconv.Atoi(sp[2])

		logger.Printf("fname: %s lnum: %d", fname, lnum)
		res = append(res, lsp.Location{
			URI: fmt.Sprintf("file://%s/%s", projRoot, fname),
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      lnum - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line:      lnum - 1,
					Character: 0,
				},
			},
		})
	}
	return res
}

func GetDefinition(logger *log.Logger, uri string, sym string) []lsp.Location {
	return getResults(logger, uri, sym, defs)
}

func GetReferences(logger *log.Logger, uri string, sym string) []lsp.Location {
	return getResults(logger, uri, sym, refs)
}
