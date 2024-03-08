#!/usr/bin/env bash

step="STATIC ANALYSIS STEP"

# - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

echo "[$step] Static Analysis task found. Executing command: 'golangci-lint run --config=.golangci-project.yml ./...'"
golangci-lint run --config=.golangci-project.yml ./...

resultLint=$?

if [[ ${resultLint} -ne 0 ]]; then
    exit 1
fi

# - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

echo "[$step] Static Analysis task found. Executing command: 'godepth -over=4'"
logError=$(find . -name "*.go" -exec godepth -over=4 {} \;)

if [[ ! -z "$logError" ]]; then
	echo $logError
    exit 1
fi

echo "[$step]: Static Code analysis ended successfully!"

# - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

echo "[$step] Static Analysis task found. Executing command: 'arch-go"
arch-go

resultLintArch=$?

if [[ ${resultLintArch} -ne 0 ]]; then
    exit 1
fi

echo "[$step]: Static Arch and Imports analysis ended successfully!"

# - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

echo "[$step] Static Analysis task found. Executing command: 'directory-validator ."
directory-validator .

resultLintDirectories=$?

if [[ ${resultLintDirectories} -ne 0 ]]; then
    exit 1
fi

echo "[$step]: Static directories analysis ended successfully!"
