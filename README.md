## Arithmetic Expression evaluator

### Abstract

This project implements a simple arithmetic expression evaluator.

It supports:

- +, -
- *, /
- parenthesis, ()
- respects operator precedence
- ignore whitespaces

#### Examples

```bash
# see parser_test.go

# or use cmd-tool

./evaluator
(1 + 1) * 2
result: 4
```

### Development

This project ships with a *Makefile**.

```bash
# build project
make

# start cmd-tool
make start

# run tests
make test

# run test with caching
make test-no-cache
```
