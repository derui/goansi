goansi
======

The goansi is a go pkg to control ANSI escape sequense, and provide some funcitons for it.

Installation
============

`go get -u github.com/derui/goansi`


Features
========
Features of goansi package are as follows.

- Simple functions appending display attributes for some texts.
- Functions to control Standard ANSI escape seqense are able to use to be parallel.

Usage
=====

    import (
      "fmt"
      "github.com/derui/goansi"
    )
    
    func main() {
      fmt.Println("%s, %s\n", goansi.Red("Hello"), goansi.Blue("World!"))
    }

