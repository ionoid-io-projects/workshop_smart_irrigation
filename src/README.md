# Introduction


# How to
Compile irrigation.go like this
```
go get github.com/stianeikeland/go-rpio
env GOOS=linux GOARCH=arm GOARM=6 go build irrigation.go
```
Copy the generated file to your Raspberry Pi device and execute it with this command

```
./irrigation
```

Happy blinking 