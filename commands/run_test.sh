#!/usr/bin/env bash

export GO_TEST_FLAGS='-tags=integration -p=1'

/commands/test.sh
