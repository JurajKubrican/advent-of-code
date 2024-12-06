## run
just run the included `./day6` binary 

## install go 1.23.4

with asdf or from the website
```
asdf install golang 1.23.4
asdf global golang 1.23.4
```


## build 
```
cd 2024/day6
go mod tidy
go build -o day6 .
```



## local benchmark
- initial
```./day6 in.txt  2.43s user 0.57s system 204% cpu 1.464 total```
- using runes instead of strings
```go run . in.txt  0.92s user 0.61s system 139% cpu 1.094 total```
- not copying memory 
```go run . in.txt  0.74s user 0.60s system 136% cpu 0.989 total```
- decrease memeory in hit tracking map
```go run . in.txt  0.70s user 0.51s system 133% cpu 0.907 total```
- byte instead of rune 
```go run . in.txt  0.67s user 0.54s system 136% cpu 0.887 total```
