# exopulse strbuf package
Package strbuf contains string buffer type which allows multiple operations on the same buffer while not forcing user to 
check for errors after each operation.

[![CircleCI](https://circleci.com/gh/exopulse/strbuf.svg?style=svg)](https://circleci.com/gh/exopulse/strbuf)
[![Build Status](https://travis-ci.org/exopulse/strbuf.svg?branch=master)](https://travis-ci.org/exopulse/strbuf)
[![GitHub license](https://img.shields.io/github/license/exopulse/strbuf.svg)](https://github.com/exopulse/strbuf/blob/master/LICENSE)

# Overview

Package strbuf contains string buffer type which allows multiple operations on the same buffer while not forcing user to 
check for errors after each operation.

Each line is separated using new-line string provided during buffer creation.

Once done with the buffer, user can interrogate error state to see if there were errors during string build.

## Features

### Buffer 

Buffer serves as intermediate storage for string based operations.

# Using strbuf package

## Installing package

Use go get to install the latest version of the library.

    $ go get github.com/exopulse/strbuf
 
Include strbuf in your application.
```go
import "github.com/exopulse/strbuf"
```

## Use Buffer to build a complex string with automatic new line separation.
```go
b := NewBuffer("\n")
b.Append("Line1")
b.EnsureEmptyLine()
b.EnsureEmptyLine()
b.AppendTitle("Test", "----")
b.AppendAll([]string{"row1", "row2"})
b.NewLine()
b.Appendf("value=%d", 1234)
```
## Before serving a string, check if there were errors.

```go
if b.Error() != nil {
	t.Fatal("unexpected error:", b.Error())
}

return s.String()
```

# About the project

## Contributors

* [exopulse](https://github.com/exopulse)

## License

Strbuf package is released under the MIT license. See
[LICENSE](https://github.com/exopulse/strbuf/blob/master/LICENSE)
