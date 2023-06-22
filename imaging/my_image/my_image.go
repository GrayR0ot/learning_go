package my_image

import "image"

type MyImage struct {
	Image image.Image `json:"my_image,omitempty"`
	Name  string      `json:"name,omitempty"`
}
