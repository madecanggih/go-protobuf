# Protobuf in Golang

Tutorial is written by [@novalagung](https://dasarpemrogramangolang.novalagung.com/C-29-golang-protobuf-implementation.html)

1. Install protoc

```
brew install protobuf
```

2. Install go protobuf library

```
go get -u github.com/golang/protobuf/protoc-gen-go
```

3. Initialize go module

```
go mod init github.com/madecanggih/protobuf
```

4. Create proto model

5. Generate model from proto file

```
protoc --go_out=. *.proto
```

6. Build main file

```
go build main.go
```

7. Run to see the result

```
./main
```