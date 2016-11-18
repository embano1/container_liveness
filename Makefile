REPO=embano1
BINARY=container_liveness
VERSION=1.1

all: image

build:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o ${BINARY}

image: build
	docker build -t ${REPO}/${BINARY}:${VERSION} .
