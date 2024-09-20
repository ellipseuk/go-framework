# Binary file name
BINARY_NAME=myframework
 
# Folder for the command application
CMD_DIR=cmd/app/main.go
 
# Commands
.PHONY: all clean build run
 
# Builds the project and runs it
all: build run
 
# Builds binary file
build:
	go build -o $(BINARY_NAME) $(CMD_DIR)

# Runs the binary file
run:
	./$(BINARY_NAME)

# Removes the binary file
clean:
	rm -f $(BINARY_NAME)