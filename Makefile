# This file is part of GOWEB Version 0.6 - Mart 2013
# Author Alexander Sychev
# GOWEB is based on program CWEB Version 3.64 - February 2002,
# Copyright (C) 1987, 1990, 1993, 2000 Silvio Levy and Donald E. Knuth
# It is distributed WITHOUT ANY WARRANTY, express or implied.
# Copyright (C) 2013 Alexander Sychev


# Permission is granted to make and distribute verbatim copies of this
# document provided that the copyright notice and this permission notice
# are preserved on all copies.

# Permission is granted to copy and distribute modified versions of this
# document under the conditions for verbatim copying, provided that the
# entire resulting derived work is given a different name and distributed
# under the terms of a permission notice identical to this one.


IFILES= \
	$(patsubst %.w, %.idx, $(wildcard *.w) $(wildcard goweave/*.w)) \
	$(patsubst %.w, %.toc, $(wildcard *.w) $(wildcard goweave/*.w)) \
	$(patsubst %.w, %.scn, $(wildcard *.w) $(wildcard goweave/*.w)) \
	$(patsubst %.w, %.log, $(wildcard *.w) $(wildcard goweave/*.w)) \
	$(patsubst %.w, %.tex, $(wildcard *.w) $(wildcard goweave/*.w)) \
	gowebman.idx gowebman.log gowebman.toc

.INTERMEDIATE: $(IFILES)

TEXP?=xetex
gcflags=-gcflags '-N -l'

all: gotangle/gotangle goweave/goweave doc

gotangle/gotangle: gotangle/gotangle.go
	(cd gotangle; go build $(gcflags))

goweave/goweave: goweave/goweave.go
	(cd goweave; go build  $(gcflags))

doc: gotangle.pdf goweave.pdf gowebman.pdf

gotangle/gotangle.go: gotangle.w gocommon.w
	-mkdir -p gotangle
	gotangle $< - $@

goweave/goweave.go: goweave.w gocommon.w
	-mkdir -p goweave
	gotangle $< - $@

%.pdf %.idx %.toc %.log: %.tex gowebmac.tex
	$(TEXP) -output-directory $(dir $<) $<

%.tex %.scn: %.w gocommon.w
	goweave/goweave $< - $(patsubst %.w, %, $<)

install: gotangle goweave
	(cd gotangle; go install)
	(cd goweave; go install)

clean:
	rm -rf *.pdf gotangle goweave goweave/*.w goweave/*.pdf $(IFILES)

tests: $(patsubst %.w, %.pdf, $(wildcard goweave/*.w))

