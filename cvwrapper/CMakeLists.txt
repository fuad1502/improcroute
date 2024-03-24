cmake_minimum_required(VERSION 2.8)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

project( cvwrapper )
find_package( OpenCV REQUIRED )
include_directories( ${OpenCV_INCLUDE_DIRS} )
add_library( cvwrapper cvwrapper.cpp )
target_link_libraries( cvwrapper ${OpenCV_LIBS} )
