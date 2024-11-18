

install :
	GOBIN=/Users/jesstruck/go/bin
	BUILD := `git rev-parse HEAD` 
	LDFLAGS=-ldflags "-X=$(GIT)build.Build=$(BUILD)"
	go install -$(LDFLAGS) -o glitter .
