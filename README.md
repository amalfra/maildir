maildir
=======
[![GitHub release](https://img.shields.io/github/release/amalfra/maildir.svg)](https://github.com/amalfra/maildir/releases)
[![Build Status](https://travis-ci.org/amalfra/maildir.svg?branch=master)](https://travis-ci.org/amalfra/maildir)
[![GoDoc](https://godoc.org/github.com/amalfra/maildir?status.svg)](https://godoc.org/github.com/amalfra/maildir)
[![Go Report Card](https://goreportcard.com/badge/github.com/amalfra/maildir)](https://goreportcard.com/report/github.com/amalfra/maildir)

A go package for reading and writing messages in the maildir format.

> The Maildir e-mail format is a common way of storing e-mail messages, where each message is kept in a separate file with a unique name, and each folder is a directory. The local filesystem handles file locking as messages are added, moved and deleted. A major design goal of Maildir is to eliminate program code having to handle locking, which is often difficult.

Refer http://cr.yp.to/proto/maildir.html and http://en.wikipedia.org/wiki/Maildir

## Installation

You can download the package using

``` go
go get github.com/amalfra/maildir
```

## Usage

Next, import the package

``` go
import (
  "github.com/amalfra/maildir"
)
```

#### Create a maildir in /home/amal/mail
``` go
myMaildir := maildir.NewMaildir("/home/amal/mail")
```

This command automatically creates the standard Maildir directories - `cur`,
`new`, and `tmp` - if they do not exist.

#### Add a new message
This creates a new file with the contents "foo"; returns the Message struct reference. Messages are written to the tmp dir then moved to new.
``` go
message, err := myMaildir.Add("foo")
```

#### List new messages
``` go
mailList, err := myMaildir.List("new")
```
This will return a map of messages by key, sorted by key

#### List current messages
``` go
mailList, err := myMaildir.List("cur")
```
This will return a map of messages by key, sorted by key

#### Find the message using key
``` go
message := maildir.Get(key)
```

#### Delete the message from disk by key
``` go
err := maildir.Delete(key)
```

**Below are the methods that are available on Message instance**

#### Get the key used to uniquely identify the message
``` go
key := message.Key()
```

#### Load the message content from file
``` go
data, err := message.GetData()
```

#### Process message - move the message from "new" to "cur"
This is usaully done to indicate that some process has retrieved the message.
``` go
key, err := message.Process("new")
```

## Development

Questions, problems or suggestions? Please post them on the [issue tracker](https://github.com/amalfra/maildir/issues).

You can contribute changes by forking the project and submitting a pull request. You can ensure the tests are passing by running ```make test```. Feel free to contribute :heart_eyes:

## UNDER MIT LICENSE

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
