#!/usr/bin/env bash

go build main.go
./main -i unsorted.dat -o sorted.dat -a qsort
