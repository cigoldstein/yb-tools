include global.mk

YBTOOLS = ycrc yugatool yugaware-client yb-connect

all: ${YBTOOLS}

.PHONY: ${YBTOOLS}
${YBTOOLS}:
	${MAKE} -C $@

clean:
	rm -rf bin/
	rm -rf out/
