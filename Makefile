GO := /usr/bin/go

all:
	@$(GO) build -o elibot-apiserver
	cp -f elibot-apiserver /usr/bin