#include "cvwrapper.h"
#include <cstring>
#include <opencv2/core/hal/interface.h>
#include <opencv4/opencv2/core/mat.hpp>
#include <opencv4/opencv2/imgcodecs.hpp>
#include <opencv4/opencv2/imgproc.hpp>

#ifdef __cplusplus 
extern "C" {
#endif

unsigned char* pngToJpg(unsigned char* inputBytes, int inputSize, int *outputSize) {
	cv::InputArray inputArray{inputBytes, inputSize};
	auto inputImage = cv::imdecode(inputArray, cv::IMREAD_UNCHANGED);
	if (inputImage.data == NULL) {
		return nullptr;
	}
	std::vector<uchar> outputVec;
	cv::imencode(".jpg", inputImage, outputVec);
	*outputSize = outputVec.size();
	unsigned char *outputBytes = (unsigned char*) malloc(sizeof(unsigned char) * *outputSize);
	memcpy((void*)outputBytes, (void*)outputVec.data(), sizeof(unsigned char) * *outputSize);
	return outputBytes;
}

unsigned char* resizeImage(unsigned char* inputBytes, int inputSize, int width, int height, int *outputSize) {
	cv::InputArray inputArray{inputBytes, inputSize};
	auto inputImage = cv::imdecode(inputArray, cv::IMREAD_UNCHANGED);
	if (inputImage.data == NULL) {
		return nullptr;
	}
	cv::Mat outputImage{};
	cv::resize(inputImage, outputImage, cv::Size(width, height));
	std::vector<uchar> outputVec;
	cv::imencode(".jpg", outputImage, outputVec);
	*outputSize = outputVec.size();
	unsigned char *outputBytes = (unsigned char*) malloc(sizeof(unsigned char) * *outputSize);
	memcpy((void*)outputBytes, (void*)outputVec.data(), sizeof(unsigned char) * *outputSize);
	return outputBytes;
}

#ifdef __cplusplus 
}
#endif

