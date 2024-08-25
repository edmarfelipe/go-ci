# Golang CI/CD with Github Actions


This repository is a simple example of how to use Github Actions to build and test a Golang application.

## Features

- Build with Cache to speed up the process
- Run tests with coverage
- Push code coverage to PR comments
- Push linting errors to PR code review

## Tools used

- For show linting errors: [golangci/golangci-lint-action](https://github.com/golangci/golangci-lint-action)
- For code coverage: [irongut/CodeCoverageSummary](https://github.com/irongut/CodeCoverageSummary)
- For convert go coverage to cobertura: [gocover-cobertura](github.com/boumenot/gocover-cobertura)
- For better test output: [gotestfmt]( github.com/gotesttools/gotestfmt)
- For comments coverage on PR: [marocchino/sticky-pull-request-comment](https://github.com/marocchino/sticky-pull-request-comment)
