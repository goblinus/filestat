.PHONY: clean
clean:
	rm -rf ./bin/*

.PHONY: build
build:
	go build -o ./bin/app ./cmd/.

.PHONY: exec
exec:
	source=dir ./bin/app

.PHONY: run
run: clean build exec