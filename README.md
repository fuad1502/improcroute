# improcroute

This project is made to complete the job test assignment for backend developer
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
IPR_PORT=8080 IPR_CORS_ORIGIN=http://localhost:8081 ./build/improcroute
```

`IPR_PORT` specify which port to install the service to. `IPR_CORS_ORIGIN` specify allowed origins.

### Using Docker

```sh
# Build
podman build -t improcroute .

# Run
podman run -d -p 127.0.0.1:8080:8080 -e IPR_PORT=8080 -e IPR_CORS_ORIGIN=http://localhost:8081 improcroute
```

Alternatively, use the docker image that I've pushed to dockerhub:
```sh
podman run -d -p 127.0.0.1:8080:8080 -e IPR_PORT=8080 -e IPR_CORS_ORIGIN=http://localhost:8081 fuad1502/improcroute
```

## Demo Client

```sh
cd demo
go run .&
open http://localhost:8081
```

Make sure to run the server too, and ensure the `IPR_CORS_ORIGIN` environment
variable is set to client's origin (http://localhost:8081).
