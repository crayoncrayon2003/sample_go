# install
```
$ sudo add-apt-repository ppa:longsleep/golang-backports
$ sudo apt update
$ sudo apt  install -y golang-go
```

# version
```
$ go version
```

# setting
```
sudo vim ~/.bash_profile
```

```
export GOPATH=$HOME/go
```

```
source ~/.bash_profile
```

# run
## case1 : run
```
$ go run filename.go
```

## case2 : build and run
```
$ go build filename.go
$ ./filename
