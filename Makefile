GO=go
GOOPTS=

RM=rm -f

SOURCE=$(wildcard *.go)
TARGET=autobahn-info

INSTALL_PATH=/usr/local/bin

version:
	$(GO) version

clean:
	$(RM) $(TARGET) 

test: $(filter-out %_test.go, $(SOURCE))
	$(GO) test ./...

all: clean version $(SOURCE)
	$(GO) $(GOOPTS) build

install: all
	@[[ -d $(INSTALL_PATH) ]] && cp -v $(TARGET)* $(INSTALL_PATH)/. || echo Directory $(INSTALL_PATH) does not exist, cancelling.
	
