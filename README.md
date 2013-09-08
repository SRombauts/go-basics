go-basics
---------

![go-basics build status](https://api.travis-ci.org/SRombauts/go-basics.png "go-basics build status")

Testing various aspects of the Go language.

Copyright (c) 2013 Sébastien Rombauts (sebastien.rombauts@gmail.com)

### Building, Running and Testing

Use the Linux Bash scripts:

```bash
./build.sh
./main
./test.sh
```

or the go commands:

```bash
go build -v github.com/srombauts/main/
./main
go test -v github.com/srombauts/multi
```

### Continuous Integration

This project is continuously tested under Ubuntu Linux 64 bits with the go 1.0 and 1.1 compilers
using the Travis CI community service with the following command line:
go build -v ./... && go test -v ./...

Detailed results can be seen online: https://travis-ci.org/SRombauts/go-basics

