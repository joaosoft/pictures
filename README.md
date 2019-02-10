# pictures
[![Build Status](https://travis-ci.org/joaosoft/pictures.svg?branch=master)](https://travis-ci.org/joaosoft/pictures) | [![codecov](https://codecov.io/gh/joaosoft/pictures/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/pictures) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/pictures)](https://goreportcard.com/report/github.com/joaosoft/pictures) | [![GoDoc](https://godoc.org/github.com/joaosoft/pictures?status.svg)](https://godoc.org/github.com/joaosoft/pictures)

A image tool that allows you to make simple tasks on images.

## Support for 
* Resize
* Crop

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## Dependecy Management
>### Dependency

Project dependencies are managed using Dep. Read more about [Dep](https://github.com/golang/dep).
* Get dependency manager: `go get github.com/joaosoft/dependency`
* Install dependencies: `dependency get`


>### Go
```
go get github.com/joaosoft/pictures
```

## Usage 
This examples are available in the project at [pictures/examples](https://github.com/joaosoft/pictures/tree/master/examples)
```go
func main() {
	path := "examples/images"
	fromFolder := fmt.Sprintf("%s/original", path)
	toFolder := fmt.Sprintf("%s/generated", path)

	picture, err := pictures.FromPath(fmt.Sprintf("%s/cat.jpg", fromFolder))
	if err != nil {
		panic(err)
	}

	testResize(picture, fmt.Sprintf("%s/resize", toFolder))
	testCrop(picture, fmt.Sprintf("%s/crop", toFolder))
}

func testResize(picture *pictures.Pictures, toFolder string) {
	fmt.Println("Resize images")

	picture.Resize(2000, 1000)

	picture.ToFile(fmt.Sprintf("%s/cat.png", toFolder), pictures.PNGEncoder())
	picture.ToFile(fmt.Sprintf("%s/cat.bmp", toFolder), pictures.BMPEncoder())
	picture.ToFile(fmt.Sprintf("%s/cat.jpeg", toFolder), pictures.JPEGEncoder(100))
}

func testCrop(picture *pictures.Pictures, toFolder string) {
	fmt.Println("Crop images")

	picture.Crop(0, 0, 1000, 1000)

	picture.ToFile(fmt.Sprintf("%s/cat.png", toFolder), pictures.PNGEncoder())
	picture.ToFile(fmt.Sprintf("%s/cat.bmp", toFolder), pictures.BMPEncoder())
	picture.ToFile(fmt.Sprintf("%s/cat.jpeg", toFolder), pictures.JPEGEncoder(100))
}
```

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
