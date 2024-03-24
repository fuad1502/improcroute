#ifndef CV_WRAPPER_H
#define CV_WRAPPER_H

#ifdef __cplusplus
extern "C" {
#endif

// toJpg reads the given image file bytes (`inputBytes`) of `inputSize` size,
// and returns JPG encoded file bytes of the image. The output file
// size is written to `outputSize` pointer.
unsigned char *toJpg(unsigned char *inputBytes, int inputSize, int *outputSize);

// resizeImage reads the given image file bytes (`inputBytes`) of `intputSize`
// size, and returns the PNG encoded file bytes of the image after it's resized
// to (`width` px, `height` px). The output file size is written to `outputSize`
// pointer.
unsigned char *resizeImage(unsigned char *inputBytes, int inputSize, int width,
                           int height, int *outputSize);

// compressImage reads the given image file bytes (`inputBytes`) of `intputSize`
// size, and returns the JPG encoded file bytes of the image with compression
// quality `quality`. The output file size is written to `outputSize`
// pointer.
unsigned char *compressImage(unsigned char *inputBytes, int inputSize,
                             int quality, int *outputSize);

#ifdef __cplusplus
}
#endif

#endif // !CV_WRAPPER_H
