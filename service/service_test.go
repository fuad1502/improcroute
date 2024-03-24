package service

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"strconv"
	"sync"
	"testing"
)

type PortGetter struct {
	mtx              sync.Mutex
	lowestUnusedPort int
}

func (portGetter *PortGetter) getUnusedPort() int {
	portGetter.mtx.Lock()
	defer portGetter.mtx.Unlock()
	portGetter.lowestUnusedPort += 1
	return portGetter.lowestUnusedPort - 1
}

var portGetter = PortGetter{lowestUnusedPort: 8080}
var waitServerStartDurationMilis time.Duration = 500

func callApiWithFile(t *testing.T, route string, inputFilePath string, mimeType string, parameters map[string]string) []byte {
	t.Parallel()
	var logBuffer bytes.Buffer
	log.SetOutput(&logBuffer)

	// Start service
	var service ImprocrouteService
	port := strconv.Itoa(portGetter.getUnusedPort())
	go service.Start(":" + port)
	defer service.Shutdown()
	time.Sleep(waitServerStartDurationMilis * time.Millisecond)

	// Read input file
	inputBuffer, err := os.ReadFile(inputFilePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v\n%s", err, string(logBuffer.Bytes()))
	}

	// Create query string
	var queryString string
	start := true
	for key := range parameters {
		if start {
			queryString += key + "=" + parameters[key]
			start = false
		} else {
			queryString += "&" + key + "=" + parameters[key]
		}
	}

	// Send request
	resp, err := http.Post("http://localhost:"+port+"/"+route+"?"+queryString, mimeType, bytes.NewReader(inputBuffer))
	if err != nil {
		t.Fatalf("Failed to issue GET request: %v\n%s", err, string(logBuffer.Bytes()))
	}
	defer resp.Body.Close()

	// Read body
	outputBuffer, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read reponse body: %v\n%s", err, string(logBuffer.Bytes()))
	}

	// Check header
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Request failed: %v\n%v\n%s", resp.Status, string(outputBuffer), string(logBuffer.Bytes()))
	}

	return outputBuffer
}

// Tests pngJpeg handler.
func TestPngJpeg(t *testing.T) {
	outputBuffer := callApiWithFile(t, "PngToJpg", "test_resource/input.png", "image/png", map[string]string{})

	refOutputBuffer, err := os.ReadFile("test_resource/ref_png_to_jpg.jpg")
	if err != nil {
		t.Fatalf("Failed to read reference output file: %v\n", err)
	}

	// Compare byte to byte between output and reference output
	if len(outputBuffer) != len(refOutputBuffer) {
		t.Fatalf("Output length (%v) differs from reference output (%v) length.\n", len(outputBuffer), len(refOutputBuffer))
	}
	for i := 0; i < len(outputBuffer); i++ {
		if outputBuffer[i] != refOutputBuffer[i] {
			t.Fatalf("Byte (%v) differs from reference output.\n", i)
		}
	}
}

// Tests resizeImage handler.
func TestResizeImage(t *testing.T) {
	parameters := map[string]string{"width": "500", "height": "500"}
	outputBuffer := callApiWithFile(t, "ResizeImage", "test_resource/input.png", "image/png", parameters)

	refOutputBuffer, err := os.ReadFile("test_resource/ref_resize_image.png")
	if err != nil {
		t.Fatalf("Failed to read reference output file: %v\n", err)
	}

	// Compare byte to byte between output and reference output
	if len(outputBuffer) != len(refOutputBuffer) {
		t.Fatalf("Output length (%v) differs from reference output (%v) length.\n", len(outputBuffer), len(refOutputBuffer))
	}
	for i := 0; i < len(outputBuffer); i++ {
		if outputBuffer[i] != refOutputBuffer[i] {
			t.Fatalf("Byte (%v) differs from reference output.\n", i)
		}
	}
}

// Tests compressImage handler.
func TestCompressImage(t *testing.T) {
	parameters := map[string]string{"quality": "50"}
	outputBuffer := callApiWithFile(t, "CompressImage", "test_resource/input.png", "image/png", parameters)

	refOutputBuffer, err := os.ReadFile("test_resource/ref_compress_image.jpg")
	if err != nil {
		t.Fatalf("Failed to read reference output file: %v\n", err)
	}

	// Compare byte to byte between output and reference output
	if len(outputBuffer) != len(refOutputBuffer) {
		t.Fatalf("Output length (%v) differs from reference output (%v) length.\n", len(outputBuffer), len(refOutputBuffer))
	}
	for i := 0; i < len(outputBuffer); i++ {
		if outputBuffer[i] != refOutputBuffer[i] {
			t.Fatalf("Byte (%v) differs from reference output.\n", i)
		}
	}
}
