package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"

	"github.com/fuad1502/improcroute/service/imgproc"
)

// ImprocrouteService encapsultes all the necessary functionalities of our
// server.
type ImprocrouteService struct {
	handler http.ServeMux
	server  http.Server
}

// addHandlers add all required handlers to the service.
func (service *ImprocrouteService) addHandlers() {
	service.handler.HandleFunc("/PngToJpg", pngToJpg)
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

func checkMimeType(headerMimeTypes []string, validMimeTypes map[string]bool) bool {
	for _, mimeType := range headerMimeTypes {
		if _, ok := validMimeTypes[mimeType]; !ok {
			return false
		}
	}
	return true
}

// pngToJpeg route handler. Convert an image from PNG format to JPG.
//
// Accepted MIME types: image/png
//
// Returned MIME types: image/jpg
//
// Parameters: none
func pngToJpg(w http.ResponseWriter, r *http.Request) {
	// Check if MIME type valid
	validMimeTypes := map[string]bool{"image/png": true}
	if !checkMimeType(r.Header["Content-Type"], validMimeTypes) {
		http.Error(w, "error", http.StatusBadRequest)
		log.Printf("pngToJpeg: invalid mime type\n")
		return
	}

	// Process
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		log.Printf("pngToJpeg: cannot read body: %v\n", err)
		return
	}
	respBody, err := imgproc.ConvertPngToJpg(body)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		log.Printf("pngToJpeg: conversion failed: %v\n", err)
		return
	}

	// Write response
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "image/jpg")
	if sentSize, err := w.Write(respBody); err != nil || sentSize != len(respBody) {
		log.Printf("pngToJpeg: %v\n", err)
	}
}

// resizeImage route handler. Resize an image based on the supplied percentage
// parameter.
//
// Accepted MIME types: image/png, image/jpg, image/jpeg
//
// Returned MIME types: image/png
//
// Parameters: `width` int REQUIRED, `height` int REQUIRED
func resizeImage(w http.ResponseWriter, r *http.Request) {
	// Check if MIME type valid
	validMimeTypes := map[string]bool{"image/png": true, "image/jpg": true, "image/jpeg": true}
	if !checkMimeType(r.Header["Content-Type"], validMimeTypes) {
		http.Error(w, "error", http.StatusBadRequest)
		log.Printf("resizeImage: invalid mime type\n")
		return
	}

	// Get width parameter
	width, err := getIntParameter(r.URL.Query(), "width")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("resizeImage: %v\n", err)
		return
	}

	// Get height parameter
	height, err := getIntParameter(r.URL.Query(), "height")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("resizeImage: %v\n", err)
		return
	}

	// Process
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		log.Printf("resizeImage: cannot read body: %v\n", err)
		return
	}
	respBody, err := imgproc.ResizeImage(body, width, height)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		log.Printf("resizeImage: resize failed: %v\n", err)
		return
	}

	// Write response
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "image/png")
	if sentSize, err := w.Write(respBody); err != nil || sentSize != len(respBody) {
		log.Printf("resizeImage: %v\n", err)
	}
}

// compressImage route handler. Compress an image.
//
// Accepted MIME types: image/*
//
// Returned MIME types: image/jpg, image/jpeg
//
// Parameters: `quality` int [0-100] REQUIRED
func compressImage(w http.ResponseWriter, r *http.Request) {
	// Check if MIME type valid
	validMimeTypes := map[string]bool{"image/png": true, "image/jpg": true, "image/jpeg": true}
	if !checkMimeType(r.Header["Content-Type"], validMimeTypes) {
		http.Error(w, "error", http.StatusBadRequest)
		log.Printf("compressImage: invalid mime type\n")
		return
	}

	// Get quality parameter
	quality, err := getIntParameterWithLimit(r.URL.Query(), "quality", 0, 100)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("compressImage: %v\n", err)
		return
	}

	// Process
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		log.Printf("compressImage: cannot read body: %v\n", err)
		return
	}
	respBody, err := imgproc.CompressImage(body, quality)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		log.Printf("compressImage: resize failed: %v\n", err)
		return
	}

	// Write response
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "image/jpg")
	if sentSize, err := w.Write(respBody); err != nil || sentSize != len(respBody) {
		log.Printf("compressImage: %v\n", err)
	}
}

func getIntParameterWithLimit(parameters url.Values, name string, minLimit int, maxLimit int) (int, error) {
	parameter, ok := parameters[name]
	if !ok {
		return 0, fmt.Errorf("parameter %v not found", name)
	}
	value, err := strconv.Atoi(parameter[0])
	if err != nil {
		return 0, fmt.Errorf("parameter %v value must be an integer", name)
	}
	if value < minLimit || value > maxLimit {
		return 0, fmt.Errorf("parameter %v value exceed limit [%v-%v]", name, minLimit, maxLimit)
	}
	return value, nil
}

func getIntParameter(parameters url.Values, name string) (int, error) {
	return getIntParameterWithLimit(parameters, name, math.MinInt, math.MaxInt)
}
