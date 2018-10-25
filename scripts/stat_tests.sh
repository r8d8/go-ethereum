#!/usr/bin/env bash

# TEST_LOG = "test.out"
# go test -v ./tests >> TEST_LOG

sort %1 && sort %2
comm -12 %1 %2
