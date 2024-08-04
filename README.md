# gfm - Golang File Manager

gfm is a command-line file manager written in Go. It provides a simple and unified interface to commonly used file management functionalities, eliminating the need to remember individual Linux commands.

[Video when i create gfm](https://youtu.be/UolQMQBzOGs)

## Features
- Files manger: info, create, delete, move, copy, read, write file
- Replaces core Linux commands: ls, rm, pwd, mv, mkdir, cat, touch, cp, and echo.

## Usage
```sh
gfm [command] [arg1] [arg2]
```

Use the commands like their Linux counterparts, with gfm preceding them.

Examples: 
```sh
# List files in the current directory
gfm ls

# Get the current working directory
gfm pwd

# View the contents of a file
gfm cat abc.txt

# Move a file
gfm mv abc.txt def.txt 

# Delete a file
gfm rm abc.txt
```

## Install
Prerequisites: [Install Golang](https://go.dev/doc/install) 

### Option 1: Install Binary
```sh
go install github.com/codegram01/gfm@latest
```

### Option 2: Install and build from source
```sh
git clone https://github.com/codegram01/gfm.git
cd gfm

# Run code: 
# Use go run . + command
go run . 

# Build 
go build .

# Install command 
go install 
```


