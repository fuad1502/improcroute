#ifndef CV_WRAPPER_H
#define CV_WRAPPER_H

#ifdef __cplusplus
extern "C" {
#endif

unsigned char *toJpg(unsigned char *inputBytes, int inputSize, int *outputSize);
unsigned char *resizeImage(unsigned char *inputBytes, int inputSize, int width,
                           int height, int *outputSize);
unsigned char *compressImage(unsigned char *inputBytes, int inputSize,
                             int quality, int *outputSize);

#ifdef __cplusplus
}
#endif

#endif // !CV_WRAPPER_H
