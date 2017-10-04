# tic-tac-go
Go (golang) console version of Tic-Tac-Toe

## Setup Go environment

- Install [Go](https://golang.org/doc/install)
    - On mac `brew install go` should work
- Setup workspace as described [here](https://golang.org/doc/code.html#Workspaces)
    - The typical setup is to create a `go` directory in `$HOME` with subdirectories for 
    `bin`, `pkg`, and `src`
  ```markdown
  $HOME/
       go/
         bin/
         pkg/
         src/
  ```
- The `GOPATH` enviroment variable defaults to `$HOME/go`, so if you put your workspace in a different location update your `GOPATH` as described [here](https://golang.org/doc/code.html#Workspaces)
- For convenience, add the workspace's bin subdirectory to your `PATH`
  ```markdown
  export PATH=$PATH:$(go env GOPATH)/bin
  ```
  
## Installation 
To install use `go get`
```markdown
go get github.com/sc2nomore/tic-tac-go
```

## Play
```markdown
tic-tac-go
```
* This assumes that the go `bin` directory was added to `PATH`. If this wasn't done the binary can be found in `$GOPATH/go/bin`.

## Run tests
From the project directory `tic-tac-go` 
```markdown
go get ./...
go test ./...
```
