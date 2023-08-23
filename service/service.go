package service

import (
	"flag"
	"fmt"
	"image/png"
	"math"
	"os"
)

type AppOption struct {
	File   *os.File
	Raduis int
}

type Point struct {
	X int
	Y int
}

func GetArgs() (AppOption, error) {
	fileName := flag.String("f", "", "The file name to read")
	raduis := flag.Int("r", 1, "Raduis of blur")
	flag.Parse()
	var errorArgs error
	if *fileName == "" {
		errorArgs = fmt.Errorf("The file name is not specified.")
	}

	file, err := os.Open(*fileName)
	if err != nil {
		errorArgs = err
	}

	return AppOption{
		File:   file,
		Raduis: *raduis,
	}, errorArgs
}

func GaussFilter(file *os.File, raduis int) {
	img, err := png.Decode(file)
	if err != nil {
		fmt.Printf("Error on read file: %s", err)
		panic(err)
	}

	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			points := getPixelsInCercle(Point{X: x, Y: y}, raduis)
			fmt.Println(points)
			return
		}
	}
}

func getPixelsInCercle(point Point, raduis int) []Point {
	var resultPoints []Point
	minX := point.X - raduis
	maxX := point.X + raduis
	minY := point.Y - raduis
	maxY := point.Y + raduis

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			dx := float64(x - point.X)
			dy := float64(y - point.Y)
			distance := math.Sqrt(dx*dx + dy*dy)
			if distance <= float64(raduis) {
				resultPoints = append(resultPoints, Point{X: x, Y: y})
			}
		}
	}

	return resultPoints
}
