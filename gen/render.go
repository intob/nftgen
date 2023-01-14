package gen

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"

	"github.com/intob/nftgen/util"
)

func Render(traitMapping map[string]Variant, base image.Image, outputPath string) error {
	result := image.NewRGBA(base.Bounds())
	draw.Src.Draw(result, base.Bounds(), base, image.Point{})
	for _, v := range traitMapping {
		if v.Img == "" {
			continue
		}
		vImage, err := util.DecodeImageFromPath(v.Img)
		if err != nil {
			return err
		}
		draw.Over.Draw(result, vImage.Bounds().Add(image.Pt(v.X, v.Y)), vImage, image.Point{})
	}
	output, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output: %s", err)
	}
	defer output.Close()
	err = jpeg.Encode(output, result, &jpeg.Options{
		Quality: jpeg.DefaultQuality,
	})
	return err
}
