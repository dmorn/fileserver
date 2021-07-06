export BINDIR ?= $(abspath bin)

PREFIX :=
TARGETS := fileserver
SRC := $(addsuffix /main.go,$(addprefix cmd/,$(TARGETS)))

BINNAMES := $(addprefix $(PREFIX), $(TARGETS))
BINS := $(addprefix $(BINDIR)/, $(BINNAMES))

all : $(BINS)

test : $(SRC)
	go test ./...

image : $(SRC)
	# TAG is not in a variable as this file is also used inside the Docker
	# builder in which we do not copy the .git stuff, hence `git describe`
	# fails there.
	docker build \
		-f Dockerfile \
		-t $(notdir $(abspath .)):$(shell git describe --tags --always) \
		.

clean :
	rm -rf $(BINDIR)/$(PREFIX)*

$(BINDIR)/$(PREFIX)% : $(SRC)
	go build -o $@ cmd/$*/main.go

$(BINS) : | $(BINDIR)

$(BINDIR) :
	mkdir -p $(BINDIR)
