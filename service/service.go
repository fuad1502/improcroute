package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/fuad1502/improcroute/service/errorreporter"
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
	service.handler.HandleFunc("/PngToJpg", addCorsMiddleware(pngToJpg))
	service.handler.HandleFunc("/ResizeImage", addCorsMiddleware(resizeImage))
	service.handler.HandleFunc("/CompressImage", addCorsMiddleware(compressImage))
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

func addCorsMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", os.Getenv("IPR_CORS_ORIGIN"))
		w.Header().Add("Access-Control-Allow-Headers", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET")
		if r.Method == "GET" || r.Method == "POST" {
			handler(w, r)
		}
	}
}

// checkMimeType is a utility function to check whether `headerMimeTypes` is in
// `validMimeTypes`
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
	reporter := errorreporter.ErrorReporter{FuncName: "pngToJpg"}
	// Check if MIME type valid
	validMimeTypes := map[string]bool{"image/png": true}
	if !checkMimeType(r.Header["Content-Type"], validMimeTypes) {
		reporter.Report(w, http.StatusBadRequest, fmt.Errorf("invalid MIME type"))
		return
	}

	// Process
	body, err := io.ReadAll(r.Body)
	if err != nil {
		reporter.Report(w, http.StatusInternalServerError, err)
		return
	}
	respBody, err := imgproc.ConvertPngToJpg(body)
	if err != nil {
		reporter.Report(w, http.StatusInternalServerError, err)
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
	reporter := errorreporter.ErrorReporter{FuncName: "resizeImage"}
	// Check if MIME type valid
	validMimeTypes := map[string]bool{"image/png": true, "image/jpg": true, "image/jpeg": true}
	if !checkMimeType(r.Header["Content-Type"], validMimeTypes) {
		reporter.Report(w, http.StatusBadRequest, fmt.Errorf("invalid MIME type"))
		return
	}

	// Get width parameter
	width, err := getIntParameter(r.URL.Query(), "width")
	if err != nil {
		reporter.Report(w, http.StatusInternalServerError, err)
		return
	}

	// Get height parameter
	height, err := getIntParameter(r.URL.Query(), "height")
	if err != nil {
		reporter.Report(w, http.StatusInternalServerError, err)
		return
	}

	// Process
	body, err := io.ReadAll(r.Body)
	if err != nil {
		reporter.Report(w, http.StatusInternalServerError, err)
		return
	}
	respBody, err := imgproc.ResizeImage(body, width, height)
	if err != nil {
		reporter.Report(w, http.StatusInternalServerError, err)
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
	reporter := errorreporter.ErrorReporter{FuncName: "compressImage"}
	// Check if MIME type valid
	validMimeTypes := map[string]bool{"image/png": true, "image/jpg": true, "image/jpeg": true}
	if !checkMimeType(r.Header["Content-Type"], validMimeTypes) {
		reporter.Report(w, http.StatusBadRequest, fmt.Errorf("invalid MIME type"))
		return
	}

	// Get quality parameter
	quality, err := getIntParameterWithLimit(r.URL.Query(), "quality", 0, 100)
	if err != nil {
		reporter.Report(w, http.StatusInternalServerError, err)
		return
	}

	// Process
	body, err := io.ReadAll(r.Body)
	if err != nil {
		reporter.Report(w, http.StatusInternalServerError, err)
		return
	}
	respBody, err := imgproc.CompressImage(body, quality)
	if err != nil {
		reporter.Report(w, http.StatusInternalServerError, err)
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
