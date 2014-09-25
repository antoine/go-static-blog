package main

import (
	"flag"
	//"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var dataSourceDirPath = "data"
var outputDirPath = "output"
var postSourceExt = ".md"

func main() {
	var postTitle = flag.String("post", "", "title of a new post")

	flag.Parse()

	if *postTitle != "" {
		log.Printf("creating new post %v.\n", *postTitle)
		postfile, err := CreatePost(dataSourceDirPath, *postTitle)
		if err != nil {
			log.Fatal(err)
		} else {
			MkdownToFile(postfile)
		}
	} else {
		log.Printf("building site.\n")
		posts := []Post{}
		//process all *.md files in source directory
		filepath.Walk(dataSourceDirPath, func(filePath string, info os.FileInfo, err error) error {
			if err == nil && postSourceExt == strings.ToLower(path.Ext(filePath)) {
				log.Printf("processing %v\n", filePath)
				//creating dir in output dir
				postDir, err := CreateHTMLDir(dataSourceDirPath, filePath, postSourceExt, outputDirPath)
				if err != nil {
					log.Fatal(err)
				} else {
					//processing the md file as html using MarkdownToFile
					postfilepath, err := HTMLToFile(filePath, postDir)
					if err != nil {
						log.Fatal(err)
					} else {
						log.Printf("%v written\n", postfilepath)
						title := path.Base(filePath)

						post := Post{
							Path:  postfilepath[len(outputDirPath)+1:],
							Title: title[0 : len(title)-len(path.Ext(filePath))]}

						posts = append(posts, post)
					}
				}
			} else if err != nil {
				log.Printf("error reading %v : %v", filePath, err)
			}
			return nil
		})

		//create the homepage & archive pages
		Homepage(dataSourceDirPath, Posts{Posts: posts}, outputDirPath)
	}

}
