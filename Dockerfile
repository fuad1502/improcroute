FROM ubuntu:latest AS install
RUN apt update
RUN apt install -y libopencv-core-dev libopencv-imgproc-dev libopencv-imgcodecs-dev
RUN apt install -y cmake
RUN apt install -y build-essential
RUN apt install -y wget
RUN wget https://go.dev/dl/go1.22.1.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.22.1.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

FROM install AS build
WORKDIR /go/src
COPY go.mod go.sum .
RUN --mount=type=cache,target=/root/.cache/go-build go mod download
COPY Makefile .
COPY main.go .
COPY service/ service/
COPY cvwrapper/CMakeLists.txt cvwrapper/
COPY cvwrapper/cvwrapper.cpp cvwrapper/
COPY cvwrapper/cvwrapper.h cvwrapper/
RUN --mount=type=cache,target=/root/.cache/go-build make

FROM install
COPY --from=build /go/src/build /bin
ENTRYPOINT ["improcroute"]
