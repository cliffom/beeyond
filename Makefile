BINARY_NAME=beeyond
DEBUG_FLAGS=gcflags "all=-N -l"

dep:
	go get

build:
	go build -o ${BINARY_NAME} *.go

debug:
	go build -${DEBUG_FLAGS} -o ${BINARY_NAME} *.go

clean:
	rm ${BINARY_NAME}

run:
	go run *.go