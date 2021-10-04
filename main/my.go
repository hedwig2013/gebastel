package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"gocv.io/x/gocv"
)

var allpaths []string
var dur1 time.Duration = time.Duration(0)
var dur2 time.Duration = time.Duration(0)

type reportHead struct {
	Reports []report
}

type report struct {
	Extension string
	Bytesize  int64
	Loadtime  int64
	Drawtime  int64
}

var rh reportHead

func main() {

	rh = reportHead{
		Reports: []report{},
	}

	time.Sleep((time.Duration(10) * time.Second))

	if len(os.Args) != 2 {
		fmt.Println("Bitte genau einen Pfad zu einer Bildersammlung angeben.")
		os.Exit(1)
	}

	path := os.Args[1]

	filepath.WalkDir(path, walker2)

	var r report

	for _, p := range allpaths {
		r = walker(p)
		if r.Extension == "---" {
			continue
		}
		rh.Reports = append(rh.Reports, r)
	}

	result, _ := json.Marshal(rh)

	st := string(result)

	fmt.Println(st)
	fmt.Println(dur1)
	fmt.Println(dur2)

}

func walker2(path string, info fs.DirEntry, err error) error {
	if !info.IsDir() {
		allpaths = append(allpaths, path)
	}
	return nil
}

func walker(pathname string) report {

	s1 := time.Now()
	theImage := gocv.IMRead(pathname, gocv.IMReadColor)
	s2 := time.Now()
	isize := theImage.Size()
	pt1 := image.Point{X: 0, Y: 0}
	pt2 := image.Point{X: isize[0], Y: isize[1]}
	gocv.Line(&theImage, pt1, pt2, color.RGBA{R: 255, G: 255, B: 255, A: 255}, 5)
	s3 := time.Now()

	gocv.IMWrite("review.png", theImage)

	dur1 += s2.Sub(s1)
	dur2 += s3.Sub(s2)

	fi, _ := os.Stat(pathname)

	var bs int64
	if fi != nil {
		bs = fi.Size()
	}

	reportdata := report{
		Extension: filepath.Ext(pathname),
		Bytesize:  bs,
		Loadtime:  s2.Sub(s1).Microseconds(),
		Drawtime:  s3.Sub(s2).Microseconds(),
	}

	return reportdata
}
