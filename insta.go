package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	xdraw "golang.org/x/image/draw"
)

// Base square size length
const HEIGHT int = 1080
const WIDTH int = 1080

// Parcentage of margin
const defaultMarginPercent int = 80

func main() {
	var percentOpt = flag.Int("p", defaultMarginPercent, "set the parcent of margin")
	var dirOpt = flag.Bool("d", false, "resize all images inside the specified directory")
	flag.Parse()

	marginPercent := *percentOpt
	if *dirOpt {
		inputPath := flag.Arg(0)
		if inputPath[len(inputPath)-1:] != "/" {
			inputPath += "/"
		}
		files, _ := ioutil.ReadDir(inputPath)
		for _, f := range files {
			inputImagePath := inputPath + f.Name()
			resize(inputImagePath, marginPercent)
		}
	} else {
		inputImagePath := flag.Arg(0)
		resize(inputImagePath, marginPercent)
	}
}

func resize(inputImagePath string, marginPercent int) {
	outName := "output-" + suffix(6) + ".png"

	inputImageFile, err := os.Open(inputImagePath)
	if err != nil {
		fmt.Println(err)
	}
	defer inputImageFile.Close()

	inputImage, _, err := image.Decode(inputImageFile)
	if err != nil {
		fmt.Println(err)
	}

	// Base image creation
	baseImage := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			baseImage.Set(j, i, color.RGBA{255, 255, 255, 255})
		}
	}

	// Resize the input image according as the longest side
	inputRect := inputImage.Bounds()
	var newX, newY int

	if inputRect.Dx() > inputRect.Dy() {
		newX = WIDTH * marginPercent / 100
		newY = inputRect.Dy() * newX / inputRect.Dx()
	} else {
		newY = HEIGHT * marginPercent / 100
		newX = inputRect.Dx() * newY / inputRect.Dy()
	}

	newSize := image.Rect(0, 0, newX, newY)
	resizedImage := image.NewRGBA(newSize)

	// Scale and paste
	xdraw.CatmullRom.Scale(resizedImage, newSize, inputImage, inputRect, draw.Over, nil)
	draw.Draw(baseImage, baseImage.Rect, resizedImage, image.Point{-1 * (WIDTH - newX) / 2, -1 * (HEIGHT - newY) / 2}, draw.Over)

	f, _ := os.OpenFile(outName, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, baseImage)
}

// Output file has random base62 suffix
func suffix(n int) string {
	rand.Seed(time.Now().UnixNano())
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
