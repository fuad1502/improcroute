package imgproc

// #cgo pkg-config: opencv4
// #cgo LDFLAGS: -L../../cvwrapper/build -lcvwrapper -lstdc++
// #cgo CFLAGS: -I../../cvwrapper/
// #include "cvwrapper.h"
import "C"

import (
	"github.com/h2non/bimg"
	"unsafe"
)

func ConvertPngToJpg(inputBuffer []byte) ([]byte, error) {
	var outputSize int
	inputBufferC := (*C.uchar)(&inputBuffer[0])
	inputSizeC := (C.int)(len(inputBuffer))
	outputSizePointerC := (*C.int)(unsafe.Pointer(&outputSize))
	outputBufferC := C.pngToJpg(inputBufferC, inputSizeC, outputSizePointerC)
	return C.GoBytes(unsafe.Pointer(outputBufferC), (C.int)(outputSize)), nil
}

func ResizeImage(inputBuffer []byte, width int, height int) ([]byte, error) {
	outputBuffer, err := bimg.NewImage(inputBuffer).ForceResize(width, height)
	if err != nil {
		return nil, err
	}
	return outputBuffer, nil
}
