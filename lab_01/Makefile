GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get

compile:
	$(GOBUILD) -o build/install.exe cmd/lab_01/main.go

install: compile
	./build/install.exe

clean:
	rm -rf build/*.exe