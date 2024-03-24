package imgproc

import (
	"github.com/h2non/bimg"
)

func ConvertPngToJpg(inputBuffer []byte) ([]byte, error) {
	outputBuffer, err := bimg.NewImage(inputBuffer).Convert(bimg.JPEG)
	if err != nil {
		return nil, err
	}
	return outputBuffer, nil
}
