#include "cvwrapper.h"
#include <cstring>
#include <opencv2/core/hal/interface.h>
#include <opencv4/opencv2/core/mat.hpp>
#include <opencv4/opencv2/imgcodecs.hpp>
#include <opencv4/opencv2/imgproc.hpp>

cv::Mat decodeWrapper(unsigned char *fileBytes, int fileSize);
unsigned char *encodeWrapper(const cv::Mat &image, const char* extension, int* size);

#ifdef __cplusplus
extern "C" {
#endif

unsigned char *pngToJpg(unsigned char *inputBytes, int inputSize,
                        int *outputSize) {
  auto inputImage = decodeWrapper(inputBytes, inputSize);
  if (inputImage.data == NULL) {
    return nullptr;
  }
  return encodeWrapper(inputImage, ".jpg", outputSize);
}

unsigned char *resizeImage(unsigned char *inputBytes, int inputSize, int width,
                           int height, int *outputSize) {
  auto inputImage = decodeWrapper(inputBytes, inputSize);
  if (inputImage.data == NULL) {
    return nullptr;
  }
  cv::Mat outputImage{};
  cv::resize(inputImage, outputImage, cv::Size(width, height));
  return encodeWrapper(outputImage, ".png", outputSize);
}

#ifdef __cplusplus
}
#endif

cv::Mat decodeWrapper(unsigned char *fileBytes, int fileSize) {
  cv::InputArray inputArray{fileBytes, fileSize};
  return cv::imdecode(inputArray, cv::IMREAD_UNCHANGED);
}

unsigned char *encodeWrapper(const cv::Mat &image, const char* extension, int* size) {
  std::vector<uchar> vec;
  cv::imencode(extension, image, vec);
  unsigned char *bytes =
      (unsigned char *)malloc(sizeof(unsigned char) * vec.size());
  memcpy((void *)bytes, (void *)vec.data(), sizeof(unsigned char) * vec.size());
  *size = vec.size();
  return bytes;
}
