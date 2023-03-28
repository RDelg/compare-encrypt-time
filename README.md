# Description
Time performance comparison between encrypting and decripting using Python and Golang.

# How to run it
## Dependencies
### Go
`go get github.com/golang/crypto`

### Python
`pip install cryptography`


## Run
`bash ./perf_test.py`

# Results
```bash
Testing Python script...

real    0m12.728s
user    0m11.881s
sys     0m0.928s
Building Go program...
Testing Go program...

real    0m0.343s
user    0m0.621s
sys     0m0.130s
```