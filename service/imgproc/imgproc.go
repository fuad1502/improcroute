package imgproc

// #cgo LDFLAGS: -L../../cvwrapper/build -lcvwrapper -lopencv_core -lopencv_imgproc -lopencv_imgcodecs -lstdc++
// #cgo CFLAGS: -I../../cvwrapper/
// #include "cvwrapper.h"
import "C"

import (
	"unsafe"
)

func ConvertPngToJpg(inputBuffer []byte) ([]byte, error) {
	var outputSize int
	inputBufferC := (*C.uchar)(&inputBuffer[0])
	inputSizeC := (C.int)(len(inputBuffer))
	outputSizePointerC := (*C.int)(unsafe.Pointer(&outputSize))
	outputBufferC := C.toJpg(inputBufferC, inputSizeC, outputSizePointerC)
	return C.GoBytes(unsafe.Pointer(outputBufferC), (C.int)(outputSize)), nil
}

func ResizeImage(inputBuffer []byte, width int, height int) ([]byte, error) {
	inputBufferC := (*C.uchar)(&inputBuffer[0])
	inputSizeC := (C.int)(len(inputBuffer))
	widthC := (C.int)(width)
	heightC := (C.int)(height)
	var outputSize int
	outputSizePointerC := (*C.int)(unsafe.Pointer(&outputSize))

	outputBufferC := C.resizeImage(inputBufferC, inputSizeC, widthC, heightC, outputSizePointerC)

	return C.GoBytes(unsafe.Pointer(outputBufferC), (C.int)(outputSize)), nil
}

func CompressImage(inputBuffer []byte, quality int) ([]byte, error) {
	var outputSize int
	inputBufferC := (*C.uchar)(&inputBuffer[0])
	inputSizeC := (C.int)(len(inputBuffer))
	outputSizePointerC := (*C.int)(unsafe.Pointer(&outputSize))
	qualityC := (C.int)(quality)
	outputBufferC := C.compressImage(inputBufferC, inputSizeC, qualityC, outputSizePointerC)
	return C.GoBytes(unsafe.Pointer(outputBufferC), (C.int)(outputSize)), nil
}
