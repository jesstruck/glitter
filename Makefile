release:
	./scripts/release.sh


install :
	GOBIN=/Users/jesstruck/go/bin
	# go install 
	go install -ldflags="-X 'github.com/jesstruck/glitter/cmd.version=0.0.2'"
	
.PHONY: install release

