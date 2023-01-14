package util

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func DecodeImageFromPath(path string) (image.Image, error) {
	if strings.HasSuffix(path, "jpg") || strings.HasSuffix(path, "jpeg") {
		f, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("failed to open file %s: %s", path, err)
		}
		defer f.Close()
		return jpeg.Decode(f)
	}
	if strings.HasSuffix(path, "png") {
		f, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("failed to open file %s: %s", path, err)
		}
		defer f.Close()
		return png.Decode(f)
	}
	return nil, fmt.Errorf("unrecognized file extension in path %s", path)
}
