# imgprocroute

This project is made to complete the job test assignment for backend developer
position @ubersnap.

## Build & deploy instructions

### Manual

#### Debian

```sh
# Install dependencies
sudo apt install libopencv-dev cmake

# Get source and build
git clone https://github.com/fuad1502/improcroute.git
cd improcroute
make

# Test
make test

# Run the server
IPR_PORT=8080 IPR_CORS_ORIGIN=http://localhost:8081 ./build/improcroute
```

### Using Docker

## Demo Client

```sh
cd demo
go run .&
open http://localhost:8081
```

Make sure to run the server too, and set the `IPR_CORS_ORIGIN` environment
variable to client's origin.
