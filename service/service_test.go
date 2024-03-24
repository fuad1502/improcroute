package service

import (
	"io"
	"net/http"
	"time"

	"strconv"
	"sync"
	"testing"
)

type PortGetter struct {
	mtx sync.Mutex
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

// Tests pngJpeg handler.
func TestPngJpeg(t *testing.T) {
	t.Parallel()
	var service ImprocrouteService
	port := strconv.Itoa(portGetter.getUnusedPort())
	go service.Start(":" + port)
	defer service.Shutdown()
	time.Sleep(waitServerStartDurationMilis * time.Millisecond)

	resp, err := http.Get("http://localhost:" + port + "/PngToJpeg")
	if err != nil {
		t.Fatalf("Failed to issue GET request\n")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read reponse")
	}

	// Stub test
	expected := "PngToJpeg"
	if body := string(body); body != expected {
		t.Fatalf("Expected: %v, Got: %v\n", expected, body)
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
