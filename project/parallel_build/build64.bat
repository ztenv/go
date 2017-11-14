set TAG=1.1.0.0
go build  -ldflags "-X main._VERSION_='%TAG%'" ./src\BuildACCT.go
move BuildACCT.exe ./bin/
pause