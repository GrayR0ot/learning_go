package task

import (
	"fmt"
	"github.com/disintegration/imaging"
	"imaging/my_image"
	"log"
	"os"
)

func LoadFiles(input string) []my_image.MyImage {
	files, err := os.ReadDir(input)
	if err != nil {
		err = os.Mkdir("input", os.FileMode(777))
		log.Fatal(err)
	}

	images := make([]my_image.MyImage, 0)

	for _, file := range files {
		src, err := imaging.Open(fmt.Sprintf("%s/%s", input, file.Name()))
		if err != nil {
			log.Fatalf("failed to open my_image: %v\n", err)
		}
		images = append(images, my_image.MyImage{
			Image: src,
			Name:  file.Name(),
		})
	}

	return images
}

func SaveFiles(images []my_image.MyImage, output string) {

	err := os.RemoveAll(output)
	if err != nil {
		panic(err)
	}
	_ = os.Mkdir(output, os.FileMode(777))

	for _, img := range images {

		err := imaging.Save(img.Image, fmt.Sprintf("%s/%s", output, img.Name))
		if err != nil {
			log.Fatalf("failed to save my_image: %v\n", err)
		}
	}
}
