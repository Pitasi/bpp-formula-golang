# BPP-Formula, in Go!
Little project for a CS exam at University of Pisa.

This project is composed by:
* **[main.go](main.go) file**, which provides a simple
CLI interface for computing _n-th_ hexadecimal digit of π or binary digit
of _log(2)_
* **[bbpformula/](https://github.com/Pitasi/progetto-esp/tree/master/bppformula)**,
the package you can easily reuse everywhere you need.<br>
It mainly exports three functions:
  1. **Pi(d)**, compute π starting at digit d.
  2. **Log2(d)**, compute log2 starting at digit d.
  3. **ToStringBase(num, base, length)**,
  generate string from a number in the given base.<br/>
  *(you should use __16__ for π and __2__ for log(2))*.

## Usage
To use the CLI interface, you need to build and run the Go code,
you can esily do this by  running:
```
$ git clone https://github.com/Pitasi/progetto-esp
$ cd progetto-esp/
$ go build -o bin/esp
$ bin/esp pi 10000
```

## Testing
You can use the tool provided by Go itself to run the test suite:
```bash
cd bppformula/
go test -v
```
