![alt text](remotsy.png)goRemotsy		[![GoDoc](https://godoc.org/github.com/jaytarang92/goget?status.svg)](https://godoc.org/github.com/jaytarang92/goremotsy)    [![Build Status](https://travis-ci.org/jaytarang92/goremotsy.svg?branch=master)](https://travis-ci.org/jaytarang92/goremotsy)  [![Go Report Card](https://goreportcard.com/badge/github.com/jaytarang92/goremotsy)](https://goreportcard.com/report/github.com/jaytarang92/goremotsy)
=========
Remotsy library for Go/Golang

Example
--------
Replace `creds.Username = ""` with your acutal username

Replace `creds.Password = ""` with your acutal password

[Click here to download example.go](examples/example.go)

```go
package main

import (
    "fmt"
    "github.com/jaytarang92/goremotsy"
)

func main() {
    remotsyAPI := remotsy.Remotsy{}
    remotsyAPI.Username = ""
    remotsyAPI.Password = ""
    remotsyAPI.GetAPIKey()
    remotes := remotsyAPI.GetRemotes()
    fmt.Println(remotes)
}

```

Credits
--------
Thanks to Jorge Cisneros (Creator of Remotsy) :smiley: [@Github](https://github.com/jorgecis/)
