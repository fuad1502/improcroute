cmake_minimum_required(VERSION 2.8)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

project( cvwrapper )
include_directories( /usr/include/opencv4 )
add_library( cvwrapper cvwrapper.cpp )
target_link_libraries( cvwrapper -lopencv_core -lopencv_imgproc -lopencv_imgcodecs )
