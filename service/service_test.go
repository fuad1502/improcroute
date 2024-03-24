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

func callApiWithFile(t *testing.T, route string, inputFilePath string, mimeType string) []byte {
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

	// Send request
	resp, err := http.Post("http://localhost:"+port+"/"+route, mimeType, bytes.NewReader(inputBuffer))
	if err != nil {
		t.Fatalf("Failed to issue GET request: %v\n%s", err, string(logBuffer.Bytes()))
	}
	defer resp.Body.Close()

	// Check header
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Request failed: %v\n%s", resp.Status, string(logBuffer.Bytes()))
	}

	// Read body
	outputBuffer, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read reponse body: %v\n%s", err, string(logBuffer.Bytes()))
	}
	return outputBuffer
}

// Tests pngJpeg handler.
func TestPngJpeg(t *testing.T) {
	outputBuffer := callApiWithFile(t, "PngToJpeg", "test_resource/input.png", "image/png")

	refOutputBuffer, err := os.ReadFile("test_resource/output_ref.jpg")
	if err != nil {
		t.Fatalf("Failed to read reference output file: %v\n", err)
	}

	// Compare byte to byte between output and reference output
	// TODO: decode then compare. Use image package
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
	t.Parallel()
	var service ImprocrouteService
	port := strconv.Itoa(portGetter.getUnusedPort())
	go service.Start(":" + port)
	defer service.Shutdown()
	time.Sleep(waitServerStartDurationMilis * time.Millisecond)

	resp, err := http.Get("http://localhost:" + port + "/ResizeImage")
	if err != nil {
		t.Fatalf("Failed to issue GET request\n")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read reponse")
	}

	// Stub test
	expected := "ResizeImage"
	if body := string(body); body != expected {
		t.Fatalf("Expected: %v, Got: %v\n", expected, body)
	}
}

// Tests compressImage handler.
func TestCompressImage(t *testing.T) {
	t.Parallel()
	var service ImprocrouteService
	port := strconv.Itoa(portGetter.getUnusedPort())
	go service.Start(":" + port)
	defer service.Shutdown()
	time.Sleep(waitServerStartDurationMilis * time.Millisecond)

	resp, err := http.Get("http://localhost:" + port + "/CompressImage")
	if err != nil {
		t.Fatalf("Failed to issue GET request\n")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read reponse")
	}

	// Stub test
	expected := "CompressImage"
	if body := string(body); body != expected {
		t.Fatalf("Expected: %v, Got: %v\n", expected, body)
	}
}
