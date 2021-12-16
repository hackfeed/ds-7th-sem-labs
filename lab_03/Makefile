GOCMD=go
GOBUILD=$(GOCMD) build

install:
	$(GOBUILD) -o build/des.exe cmd/des/main.go

clean:
	rm -rf build/*.exe *.pem *.sig