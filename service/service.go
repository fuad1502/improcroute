package service

import (
	"context"
	"fmt"
	"net/http"
)

// ImprocrouteService encapsultes all the necessary functionalities of our
// server.
type ImprocrouteService struct {
	handler http.ServeMux
	server  http.Server
}

// addHandlers add all required handlers to the service.
func (service *ImprocrouteService) addHandlers() {
	service.handler.HandleFunc("/PngToJpeg", pngToJpeg)
	service.handler.HandleFunc("/ResizeImage", resizeImage)
	service.handler.HandleFunc("/CompressImage", compressImage)
}

// Start adds all required handlers and starts the service at the specified
// address.
func (service *ImprocrouteService) Start(addr string) {
	service.addHandlers()
	service.server.Handler = &service.handler
	service.server.Addr = addr
	service.server.ListenAndServe()
}

// Shutdown shuts down the service.
func (service *ImprocrouteService) Shutdown() {
	service.server.Shutdown(context.Background())
}

// pngToJpeg route handler. Convert an image from PNG format to JPG.
//
// Accepted MIME types: image/png
//
// Returned MIME types: image/jpg
//
// Parameters: none
func pngToJpeg(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PngToJpeg");
}

// resizeImage route handler. Resize an image based on the supplied percentage
// parameter.
//
// Accepted MIME types: image/* 
//
// Returned MIME types: image/*
//
// Parameters: `percentage` float
func resizeImage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ResizeImage");
}

// compressImage route handler. Compress an image.
//
// Accepted MIME types: image/* 
//
// Returned MIME types: image/*
func compressImage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CompressImage");
}
