package service

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
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
	rgbaImage := img.(*image.RGBA)

	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			points := getPixelsInCercle(Point{X: x, Y: y}, raduis)
			color := getAvarageColorInCircle(points, Point{X: x, Y: y}, img, width, height)
			rgbaImage.SetRGBA(x, y, color)
		}
	}

	newFile, err := os.Create("new_img.png")
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(newFile, rgbaImage)
	if err != nil {
		log.Fatal(err)
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

func getAvarageColorInCircle(cirle []Point, currentPoint Point, img image.Image, maxX int, maxY int) color.RGBA {
	var r, g, b, a, count uint32
	r, g, b, a = 0, 0, 0, 0
	count = 0
	for _, value := range cirle {
		var rN, gN, bN, aN uint32
		if value.X < 0 || value.X > maxX || value.Y < 0 || value.Y > maxY {
			continue
		} else {
			color := img.At(value.X, value.Y)
			rN, gN, bN, aN = color.RGBA()
		}
		r, g, b, a = r+rN, g+gN, b+bN, a+aN
		count++
	}
	r = r / count / 256
	b = b / count / 256
	g = g / count / 256
	a = a / count / 256
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}
