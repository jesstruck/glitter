release:
	./scripts/release.sh


install :
	GOBIN=/Users/jesstruck/go/bin
	# go install 
	go install -ldflags="-X 'github.com/jesstruck/glitter/cmd.version=$(shell git describe --tags --abbrev=8 --dirty --always --long)'"
	
.PHONY: install release

