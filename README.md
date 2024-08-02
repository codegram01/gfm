# gfm - Golang File Manager
This project build command line tool name gfm for manager file 

## It replace these linux command
- ls
- rm
- pwd
- mv
- mkdir
- cat
- touch
- cp
- echo

# Usage
```sh
gfm [command] [arg1] [arg2]
```

Just use like linux command but have gfm in front 
Like: 
```sh
$ gfm ls
$ gfm pwd
$ gfm cat abc.txt
$ gfm mv abc.txt def.txt
$ gfm rm abc.txt
```

# Install
- install Golang
- Run: `go run .` + [command] + [arg...]
- Build: `go build .`
Command gfm will available in your path


