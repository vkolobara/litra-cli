BINARY := litra
GO     := go

.PHONY: all build clean install run

all: build

build:
	$(GO) build -ldflags "-s -w" -o $(BINARY) .

clean:
	rm -f $(BINARY)

install: build
	install -m 0755 $(BINARY) /usr/local/bin/$(BINARY)

run:
	$(GO) run .
