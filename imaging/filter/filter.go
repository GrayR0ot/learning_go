package filter

import (
	"fmt"
	"github.com/disintegration/imaging"
	"imaging/my_image"
	"time"
)

func blur(img *my_image.MyImage) {
	img.Image = imaging.Blur(img.Image, 5)
}

func grayScale(img *my_image.MyImage) {
	img.Image = imaging.Grayscale(img.Image)
}

func Filter(img *my_image.MyImage) my_image.MyImage {
	time.Sleep(250 * time.Millisecond)
	blur(img)
	grayScale(img)
	img.Name = fmt.Sprintf("filtered_%s", img.Name)
	return *img
}
