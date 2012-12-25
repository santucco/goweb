IFILES= \
	$(patsubst %.w, %.idx, $(wildcard *.w) $(wildcard goweave/*.w)) \
	$(patsubst %.w, %.toc, $(wildcard *.w) $(wildcard goweave/*.w)) \
	$(patsubst %.w, %.scn, $(wildcard *.w) $(wildcard goweave/*.w)) \
	$(patsubst %.w, %.log, $(wildcard *.w) $(wildcard goweave/*.w)) \
	$(patsubst %.w, %.tex, $(wildcard *.w) $(wildcard goweave/*.w))

.INTERMEDIATE: $(IFILES)

TEXP?=xetex
gcflags=-gcflags '-N -l'

all: gotangle/gotangle goweave/goweave doc

gotangle/gotangle: gotangle/gotangle.go
	(cd gotangle; go build $(gcflags))

goweave/goweave: goweave/goweave.go
	(cd goweave; go build  $(gcflags))

doc: gotangle.pdf goweave.pdf

gotangle/gotangle.go: gotangle.w common.w
	-mkdir -p gotangle
	gotangle $< - $@

goweave/goweave.go: goweave.w common.w
	-mkdir -p goweave
	gotangle $< - $@

%.pdf %.idx %.toc %.log: %.tex gowebmac.tex
	$(TEXP) -output-directory $(dir $<) $<

%.tex %.scn: %.w common.w
	goweave/goweave $< - $(patsubst %.w, %, $<)

install: gotangle goweave
	(cd gotangle; go install)
	(cd goweave; go install)

clean:
	rm -rf *.pdf gotangle goweave goweave/*.w goweave/*.pdf $(IFILES)

tests: $(patsubst %.w, %.pdf, $(wildcard goweave/*.w))

