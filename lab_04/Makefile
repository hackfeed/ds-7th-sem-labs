GOCMD=go
GOBUILD=$(GOCMD) build

install:
	$(GOBUILD) -o build/rsa.exe cmd/rsa/main.go

clean:
	rm -rf build/*.exe *.pem *.sig