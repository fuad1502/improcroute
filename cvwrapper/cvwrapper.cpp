#include "cvwrapper.h"
#include <cstring>
#include <opencv2/core/hal/interface.h>
#include <opencv4/opencv2/core/mat.hpp>
#include <opencv4/opencv2/imgcodecs.hpp>
#include <opencv4/opencv2/imgproc.hpp>

cv::Mat decodeWrapper(unsigned char *fileBytes, int fileSize);
unsigned char *encodeWrapper(const cv::Mat &image, const char *extension,
                             int *size, int quality = 95);

#ifdef __cplusplus
extern "C" {
#endif

unsigned char *toJpg(unsigned char *inputBytes, int inputSize,
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

unsigned char *compressImage(unsigned char *inputBytes, int inputSize,
                             int quality, int *outputSize) {
  auto inputImage = decodeWrapper(inputBytes, inputSize);
  if (inputImage.data == NULL) {
    return nullptr;
  }
  return encodeWrapper(inputImage, ".jpg", outputSize, quality);
}

#ifdef __cplusplus
}
#endif

cv::Mat decodeWrapper(unsigned char *fileBytes, int fileSize) {
  cv::InputArray inputArray{fileBytes, fileSize};
  return cv::imdecode(inputArray, cv::IMREAD_UNCHANGED);
}

unsigned char *encodeWrapper(const cv::Mat &image, const char *extension,
                             int *size, int quality) {
  std::vector<int> flags;
  if (strcmp(extension, ".jpg") == 0 || strcmp(extension, ".jpeg")) {
    flags.push_back(cv::IMWRITE_JPEG_QUALITY);
    flags.push_back(quality);
  }
  std::vector<uchar> vec;
  cv::imencode(extension, image, vec, flags);
  unsigned char *bytes =
      (unsigned char *)malloc(sizeof(unsigned char) * vec.size());
  memcpy((void *)bytes, (void *)vec.data(), sizeof(unsigned char) * vec.size());
  *size = vec.size();
  return bytes;
}
