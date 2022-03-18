package utils

import (
	"image"
	"os"
)

func GetImageFromFilePath(filePath string) (image.Image, error) {
	// Special thanks:
	// https://stackoverflow.com/a/49595208

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}
