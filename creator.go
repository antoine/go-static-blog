package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	//"errors"
	//"log"
)

func CreatePost(dataSourceDir string, postTitle string) (*os.File, error) {
	path, err := CreatePostDir(dataSourceDir)
	if err != nil {
		return nil, err
	}

	postfilepath := filepath.Join(path, postTitle+".md")
	if _, err := os.Stat(postfilepath); err == nil {
		return nil, fmt.Errorf("file already exists: %v\n", postfilepath)
	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf("error when checking existence of %v. %v", postfilepath, err)
	}

	postfile, err := os.Create(postfilepath)

	if err != nil {
		return nil, fmt.Errorf("could not create post file. %v", err)
	}

	fmt.Printf("%v created\n", postfile.Name())
	return postfile, nil

}

func CreatePostDir(dataSourceDir string) (string, error) {
	year, month, day := time.Now().Date()
	//using %d since the month overrides toString to show the english name
	path := filepath.Join(dataSourceDir, fmt.Sprintf("%d", year), fmt.Sprintf("%d", month), fmt.Sprintf("%d", day))

	if errFullPath := os.MkdirAll(path, os.ModePerm); errFullPath != nil {
		return path, fmt.Errorf("could not create source directory %v. %v", path, errFullPath)
	}

	return path, nil
}

func CreateHTMLDir(dataSourceDir string, dataSourceFile string, postsourceext string, outputDir string) (string, error) {
	if strings.Index(dataSourceFile, dataSourceDir) == 0 {

		//it is expected that dataSourceFile path starts with dataSourceDir and ends with  postsourceext
		newDir := filepath.Join(outputDir, dataSourceFile[len(dataSourceDir):len(dataSourceFile)-len(postsourceext)])

		if err := os.MkdirAll(newDir, os.ModePerm); err != nil {
			return newDir, fmt.Errorf("could not create target directory %v : %v", newDir, err)
		}
		return newDir, nil
	} else {
		return "", fmt.Errorf("%v doesn't starts with %v", dataSourceFile, dataSourceDir)
	}

}
