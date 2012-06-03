#.INTERMEDIATE: gotangle/gotangle.go goweave/goweave.go
gcflags=-gcflags '-N -l'

all: gotangle goweave

gotangle: gotangle/gotangle.go
	(cd gotangle; go build $(gcflags))

goweave: goweave/goweave.go
	(cd goweave; go build  $(gcflags))

gotangle/gotangle.go: gotangle.w common.w
	-mkdir -p gotangle
	gotangle $< - $@

goweave/goweave.go: goweave.w prod.w common.w
	-mkdir -p goweave
	gotangle $< - $@

install: gotangle goweave
	(cd gotangle; go install)
	(cd goweave; go install)

clean:
	rm -f *.scn *.idx *.tex
	rm -rf gotangle/ goweave/