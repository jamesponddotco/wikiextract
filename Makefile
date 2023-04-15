.POSIX:
.SUFFIXES:

PREFIX=/usr/local
BINDIR=bin
MANDIR=share/man
GO=go
GOFUMPT=gofumpt
GOLINT=golangci-lint
GOVULN=govulncheck
RM = rm
INSTALL = install
SCDOC = scdoc
GOBUILD_OPTS=-trimpath

all: build doc

init: # Downloads and verifies project dependencies and tooling.
	$(GO) get
	$(GO) install mvdan.cc/gofumpt@v0.5.0
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2
	$(GO) install golang.org/x/vuln/cmd/govulncheck@latest

build: # Builds an application binary.
	$(GO) build $(GOBUILD_OPTS)

doc: # Builds the manpage.
	$(SCDOC) <doc/wikiextract.1.scd >doc/wikiextract.1

install: # Installs the release binary.
	$(INSTALL) -d \
		$(DESTDIR)$(PREFIX)/$(BINDIR)/ \
		$(DESTDIR)$(PREFIX)/$(MANDIR)/man1/
	$(INSTALL) -pm 0755 wikiextract $(DESTDIR)$(PREFIX)/$(BINDIR)/
	$(INSTALL) -pm 0644 doc/wikiextract.1 $(DESTDIR)$(PREFIX)/$(MANDIR)/man1/

fmt: # Formats Go source files in this repository.
	$(GOFUMPT) -e -extra -w .

lint: # Runs golangci-lint using the config at the root of the repository.
	$(GOLINT) run ./...

vulnerabilities: # Analyzes the codebase and looks for vulnerabilities affecting it.
	$(GO) install golang.org/x/vuln/cmd/govulncheck@latest
	$(GOVULN) ./...

test: # Runs unit tests.
	$(GO) test -cover -race -vet all -mod readonly ./...

test/coverage: # Generates a coverage profile and open it in a browser.
	$(GO) test -coverprofile cover.out
	$(GO) tool cover -html=cover.out

clean: # Cleans cache files from tests and deletes any build output.
	$(RM) -f wikiextract doc/wikiextract.1

.PHONY: all init build doc install fmt lint vulnerabilities test test/coverage clean
