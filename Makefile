# Define the Go command
GO := go

# Define the main package
MAIN_PACKAGE := ./main.go

# Build the Go project
build:
	$(GO) build -o sentinel $(MAIN_PACKAGE)

# Run the Go project
run:
	$(GO) run $(MAIN_PACKAGE)

# Clean the build
clean:
	rm -f sentinel

# Define the default target
.PHONY: build run clean
