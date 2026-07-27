.PHONY: build install run flash

build:
	go build -o bin/flash ./cmd/flash

install:
	go install ./cmd/flash

run:
	go run main.go

flash:
	go run ./cmd/flash $(filter-out $@,$(MAKECMDGOALS))

%:
	@true
