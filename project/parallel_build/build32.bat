set GOOS=windows
set GOARCH=386
set CGO_ENABLED=1
set TAG=1.0.9.0
go build  -ldflags "-X main._VERSION_='%TAG%'" ./src\BuildACCT.go
move BuildACCT.exe ./bin/
pause