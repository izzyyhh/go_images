package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/h2non/bimg"
)

func main() {
	buffer, err := bimg.Read("amk.png")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	newImage := bimg.NewImage(buffer)
	size, err := newImage.Size()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v, %v", size.Height, size.Width)

	DIVIDER := 2

	changedBuff, err := newImage.Resize(size.Width/DIVIDER, size.Height/DIVIDER)

	if err != nil {
		log.Fatal(err)
	}

	bimg.Write("new_amk.png", changedBuff)

	fmt.Printf("%v", newImage.Type())

}

func loadImage(fileLocation string) (image.Image, error) {
	reader, err := os.Open("amk.png")

	if err != nil {
		log.Fatal("Error: Opening File", err)
	}
	defer reader.Close()

	img, format, err := image.Decode(reader)
	fmt.Println("Format:", format)
	if err != nil {
		return nil, err
	}

	return img, err
}
