package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/fogleman/gg"
)

var allpaths []string
var dur1 time.Duration = time.Duration(0)
var dur2 time.Duration = time.Duration(0)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Bitte genau einen Pfad zu einer Bildersammlung angeben.")
		os.Exit(1)
	}

	path := os.Args[1]

	filepath.WalkDir(path, walker2)

	for _, p := range allpaths {
		walker(p)
	}

	fmt.Println(dur1)
	fmt.Println(dur2)

}

func walker2(path string, info fs.DirEntry, err error) error {
	if !info.IsDir() {
		allpaths = append(allpaths, path)
	}
	return nil
}

func walker(pathname string) {

	s1 := time.Now()
	i, err := gg.LoadImage(pathname)
	if err != nil {
		return
	}
	s2 := time.Now()
	dc := gg.NewContextForImage(i)
	height := float64(dc.Height())
	width := float64(dc.Width())
	dc.SetHexColor("ffffff")
	dc.SetLineWidth(5)
	dc.DrawLine(0, 0, width, height)
	dc.DrawLine(width, 0, 0, height)
	dc.Stroke()
	s3 := time.Now()

	// dc.SavePNG("review.png")

	dur1 += s2.Sub(s1)
	dur2 += s3.Sub(s2)

}
