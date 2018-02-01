[![Build Status](https://travis-ci.org/cthulhu/jsonpath.svg?branch=master)](https://travis-ci.org/cthulhu/jsonpath)  [![Goreport](https://goreportcard.com/badge/github.com/cthulhu/jsonpath)](https://goreportcard.com/badge/github.com/cthulhu/jsonpath) [![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/cthulhu/jsonpath/master/LICENSE)


# jsonpath - Json generation by dot notation

This library is very close to [mongodb's dot notation](https://docs.mongodb.com/manual/core/document/#dot-notation) with a bit of extensions.

Supported syntaxes

Standard dot notation

|input hash value | output json result      |
|-----------------|-------------------------|
|{"0.value":"100"}| [{"value":"100"}]       |
|{"1.value":"100"}| [null, {"value":"100"}] |
|{"value":"100"}  |  {"value":"100"}        |
|{"value.1":"100"}|  {"value":[null, "100"]}|
|{"v.0.k":"100"}  |  {"v":[{"k":100}]}      |

Extended

|input hash value | output json result      |
|-----------------|-------------------------|
|{"v.num()":"1.0"}  |  {"v":1.0}      |
|{"v.bool()":"true"}  |  {"v":true}      |


# Benchmarks

Run benchmarks

    go test -bench=.

Results

v0.0.3

    BenchmarkComplexJSONPathArray-8             	  100000	     12480 ns/op
    BenchmarkSimpleJSONPathArrayWithNum-8       	  500000	      2654 ns/op
    BenchmarkSimpleJSONPathArrayWithBool-8      	 1000000	      2319 ns/op
    BenchmarkSimpleJSONPathArrayInsideArray-8   	  500000	      3066 ns/op
    BenchmarkSimpleJSONPathArrays-8             	  500000	      2636 ns/op
    BenchmarkSimpleJSONPathSimple-8             	 1000000	      1626 ns/op
    BenchmarkJSONNative-8                       	 2000000	       959 ns/op

# Installation

go get github.com/cthulhu/jsonpath

# Usage

    in := map[string]string{"0.value": "100.00"}
    jsobBytes := jsonpath.Marshal(in)

For more examples check jsonpath_test.go file

# LICENSE

See LICENSE file
