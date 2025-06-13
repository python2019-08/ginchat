FROM alpine:latest

WORKDIR /xxapp
COPY build/gotest .

ENTRYPOINT ["/bin/sh", "-c", "ls -l ${WORKDIR} && exec sh"]
CMD []

