package main

import (
	"fmt"
	"pictures"
)

func main() {
	path := "examples/images"
	fromFolder := fmt.Sprintf("%s/original", path)
	toFolder := fmt.Sprintf("%s/generated", path)

	picture, err := pictures.Open(fmt.Sprintf("%s/cat.jpg", fromFolder))
	if err != nil {
		panic(err)
	}

	testResize(picture, fmt.Sprintf("%s/resize", toFolder))
	testCrop(picture, fmt.Sprintf("%s/crop", toFolder))
}

func testResize(picture *pictures.Pictures, toFolder string) {
	fmt.Println("Resize images")

	picture.Resize(2000, 1000)

	picture.Save(fmt.Sprintf("%s/cat.png", toFolder), pictures.PNGEncoder())
	picture.Save(fmt.Sprintf("%s/cat.bmp", toFolder), pictures.BMPEncoder())
	picture.Save(fmt.Sprintf("%s/cat.jpeg", toFolder), pictures.JPEGEncoder(100))
}

func testCrop(picture *pictures.Pictures, toFolder string) {
	fmt.Println("Crop images")

	picture.Crop(0, 0, 1000, 1000)

	picture.Save(fmt.Sprintf("%s/cat.png", toFolder), pictures.PNGEncoder())
	picture.Save(fmt.Sprintf("%s/cat.bmp", toFolder), pictures.BMPEncoder())
	picture.Save(fmt.Sprintf("%s/cat.jpeg", toFolder), pictures.JPEGEncoder(100))
}
