package main

import (
	"log"
	"os"
	"text/template"
	// "fmt"
	"bufio"
	"io"
	"path/filepath"
)

type Posts struct {
	Posts []Post
}

type Post struct {
	Path  string
	Title string
}

func Homepage(dataSourceDir string, homepage Posts, outputDir string) {
	s1, _ := template.ParseFiles(filepath.Join(dataSourceDir, "home.gotemplate"))
	outputfilepath := filepath.Join(outputDir, "index.html")
	outputfile, err := os.Create(outputfilepath)
	if err != nil {
		log.Fatalf("could not create target file %v", err)
	}

	w := bufio.NewWriter(io.Writer(outputfile))

	if err := s1.ExecuteTemplate(w, "body", homepage); err != nil {
		log.Fatalf("template %v execution failed: %v", s1, err)

	}
	log.Printf("%v written\n", outputfilepath)
	w.Flush()
}
