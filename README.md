# improcroute

This project was made to complete the job test assignment for backend developer
position @ubersnap.

## Build & deploy instructions

### Manual

#### Debian

I assume go tool has been installed on your system and can be called through the `go` command.

```sh
# Install dependencies
sudo apt install libopencv-core-dev libopencv-imgproc-dev libopencv-imgcodecs-dev cmake build-essential

# Get source and build
git clone https://github.com/fuad1502/improcroute.git
cd improcroute
make

# Test
make test

# Run the server
IPR_PORT=8080 IPR_CORS_ORIGIN=\* ./build/improcroute
```

`IPR_PORT` specify which port to install the service to. `IPR_CORS_ORIGIN` specify allowed origins.

### Using Docker

```sh
# Build
podman build -t improcroute .

# Run
podman run -d -p 127.0.0.1:8080:8080 -e IPR_PORT=8080 -e IPR_CORS_ORIGIN=\* improcroute
```

Alternatively, use the docker image that I've pushed to dockerhub:
```sh
podman run -d -p 127.0.0.1:8080:8080 -e IPR_PORT=8080 -e IPR_CORS_ORIGIN=\* fuad1502/improcroute
```

## Demo Client

```sh
cd demo
go run .&
open http://localhost:8081
```

Make sure to run the server too, and ensure the `IPR_CORS_ORIGIN` environment
variable includes the client's origin (http://localhost:8081).

## API

### `POST /PngToJpg`
Converts a PNG file to JPG.
- Request MIME types: image/png
- Response MIME types: image/jpg
- Query string parameters: none

### `POST /ResizeImage`
Resize the given image to the specified width and height. The resulting image
is converted to PNG regardless of the original format.
- Request MIME types: image/png, image/jpg, image/jpeg
- Response MIME types: image/png
- Query string parameters: `width int` REQUIRED, `height int` REQUIRED

### `POST /CompressImage`
Compress the given image using JPEG compression with the specified quality
value. Lower quality value results in lower quality image, but higher
compression ratio. The resulting image is converted to JPG regardless of the
original format.
- Request MIME types: image/png, image/jpg, image/jpeg
- Response MIME types: image/jpg
- Query string parameters: `quality int` [0-100] REQUIRED
