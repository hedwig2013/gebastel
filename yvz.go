package main

import (
	"image"
	"log"
	"os"
	"path/filepath"
)

func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func WalkMatch(root string) ([]string, error) {
	const pattern = "*.jpg"
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func Walkdir(path string) {
	files, _ := WalkMatch(path)

	for _, s := range files {
		log.Println(s)
	}

}

// func main() {
// 	fmt.Println("Hello World")

// }

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}
