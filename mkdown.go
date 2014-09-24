package main

import (
	"bufio"
	"github.com/knieriem/markdown"
	"io"
	"os"
	"fmt"
	"path/filepath"
)

func Mkdown() {
	p := markdown.NewParser(&markdown.Extensions{Smart: true})
	w := bufio.NewWriter(os.Stdout)
	p.Markdown(os.Stdin, markdown.ToHTML(w))
	w.Flush()
}

func MkdownToFile(postfile *os.File) {
	w := bufio.NewWriter(io.Writer(postfile))
	stdin := bufio.NewReader(os.Stdin)
	w.ReadFrom(stdin)
	w.Flush()
}

func HTMLToFile(postfilepath string, outputfiledirpath string) (string, error) {
	postfile, err := os.Open(postfilepath)
	if err != nil {
		return "", fmt.Errorf("could not open post file %v. %v", postfile, err)
	}

  outputfilepath := filepath.Join(outputfiledirpath, "index.html")
	outputfile, err := os.Create(outputfilepath)
	if err != nil {
		return outputfilepath, fmt.Errorf("could not create target file %v. %v", outputfile, err)
	}

	p := markdown.NewParser(&markdown.Extensions{Smart: true})
	w := bufio.NewWriter(io.Writer(outputfile))
	r := bufio.NewReader(io.Reader(postfile))
	p.Markdown(r, markdown.ToHTML(w))
	w.Flush()

	return outputfilepath, nil
}
