SHELL = /bin/sh

.PHONY: clobber docfiles

JSONFILE = ../codeplug/codeplugs.json
BINDIR = $(GOPATH)/bin
SOURCES = *.go

docfiles: $(BINDIR)/docCodeplug $(JSONFILE)
	$(BINDIR)/docCodeplug $(JSONFILE)

$(BINDIR)/docCodeplug: $(SOURCES)
	go install

clobber:
	rm -f $(BINDIR)/docCodeplug
