# Define the Go command
GO := go

# Define the main package
MAIN_PACKAGE := ./cmd/main.go

# Build the Go project
build:
	$(GO) build -o bin/sentinel $(MAIN_PACKAGE)

# Run the Go project
run:
	$(GO) run $(MAIN_PACKAGE)

# Clean the build
clean:
	rm -f bin/sentinel

# Define the default target
.PHONY: build run clean
