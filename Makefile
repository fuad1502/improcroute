.PHONY: build
build: cvwrapper/build/libcvwrapper.a
	go build
	mkdir -p build
	mv improcroute build/improcroute

cvwrapper/build/libcvwrapper.a: cvwrapper/cvwrapper.cpp cvwrapper/cvwrapper.h
	cd cvwrapper && mkdir -p build && cd build && cmake .. && make
	go build -a
	go clean -testcache

.PHONY: clean
clean:
	go clean -testcache
	rm -rf build
	rm -rf cvwrapper/build

.PHONY: test
test: build
	go test -v ./...	
