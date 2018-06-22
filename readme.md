# Check Site Every 5 Min

A `go` project that check website every 5min.

## Description

User can enter a web url and expect a ✅ or ❌ base on the health of the site.

The site will also be scheduled to check every 5 min and will automatically update the status.

## Getting Started

To run this project, `golang` should already be installed. If not, visit [here](https://golang.org/doc/install) to follow the instruction.

#### Step 1: Clone the project in `go` workspace directory

```bash
$ cd $HOME/go/src # This path can varies if you use a different directory when installed
$ git clone ...
```

#### Step 2: Build the project

```bash
$ cd goweb
$ go build web.go
```

If there are error with cannot find `package`, install the package first with `go get`. Example:

```bash
$ go build web.go
web.go:4:3: cannot find package "github.com/gorilla/websocket" in any of:
  /usr/local/Cellar/go/1.9.2/libexec/src/github.com/gorilla/websocket (from $GOROOT)
  /Users/zac/go/src/github.com/gorilla/websocket (from $GOPATH)
$ go get github.com/gorilla/websocket
```

#### Step 3: Run the project

```bash
$ ./web
```

Visit [here](http://localhost:8080/) to start playing with the app.

#### Running Test

The project also include tests. Run it with:

```bash
$ go test --cover
```
