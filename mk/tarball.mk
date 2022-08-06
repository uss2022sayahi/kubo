

ifeq (,$(wildcard .tarball))
tarball-is:=0
else
tarball-is:=1
# override git hash
git-hash:=$(shell cat .tarball)
endif

GOCC ?= go1.17.13

go-ipfs-source.tar.gz: distclean
	GOCC=$(GOCC) bin/maketarball.sh $@

kubo-source.tar.gz: distclean
	GOCC=$(GOCC) bin/maketarball.sh $@
