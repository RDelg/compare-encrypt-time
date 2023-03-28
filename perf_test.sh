#!/bin/bash

set -e

# Set the number of iterations to run
ITERATIONS=100

# Generate test data file with 1024 rows
rm -f testdata.txt
for i in {1..1024}; do
  head -c 30 </dev/urandom | base64 >> testdata.txt
done

# Measure the time taken by the Python script
echo "Testing Python script..."
time (
  for i in $(seq 1 $ITERATIONS); do
    cat testdata.txt | python script.py >/dev/null
  done
)

# Measure the time taken by the Go program
echo "Building Go program..."
go build encrypt.go
go build decrypt.go
echo "Testing Go program..."
time (
  for i in $(seq 1 $ITERATIONS); do
    cat testdata.txt | ./encrypt | ./decrypt >/dev/null
  done
)