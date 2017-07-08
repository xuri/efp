# EFP (Excel Formula Parser)

[![Build Status](https://travis-ci.org/Luxurioust/efp.svg?branch=master)](https://travis-ci.org/Luxurioust/efp)
[![Code Coverage](https://codecov.io/gh/Luxurioust/efp/branch/master/graph/badge.svg)](https://codecov.io/gh/Luxurioust/efp)
[![Go Report Card](https://goreportcard.com/badge/github.com/Luxurioust/efp)](https://goreportcard.com/report/github.com/Luxurioust/efp)
[![GoDoc](https://godoc.org/github.com/Luxurioust/efp?status.svg)](https://godoc.org/github.com/Luxurioust/efp)
[![Licenses](https://img.shields.io/badge/license-bsd-orange.svg)](https://opensource.org/licenses/BSD-3-Clause)

Using EFP (Excel Formula Parser) you can get an Abstract Syntax Tree (AST) from Excel formula.

## Installation

```go
go get github.com/Luxurioust/efp
```

## Example

```go
package main

import "github.com/Luxurioust/efp"

func main() {
    ps := efp.ExcelParser()
    ps.Parse("=SUM(A3+B9*2)/2")
    println(ps.PrettyPrint())
}
```

Get AST

```
SUM <Function> <Start>
    A3 <Operand> <Range>
    + <OperatorInfix> <Math>
    B9 <Operand> <Range>
    * <OperatorInfix> <Math>
    2 <Operand> <Number>
 <Function> <Stop>
/ <OperatorInfix> <Math>
2 <Operand> <Number>
```

## Contributing

Contributions are welcome! Open a pull request to fix a bug, or open an issue to discuss a new feature or change.

## Credits

EFP (Excel Formula Parser) is a Golang port of [E. W. Bachtal's](http://ewbi.blogs.com/develops/2004/12/excel_formula_p.html) Excel formula parser.

## Licenses

This program is under the terms of the BSD 3-Clause License. See [https://opensource.org/licenses/BSD-3-Clause](https://opensource.org/licenses/BSD-3-Clause).