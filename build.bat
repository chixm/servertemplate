set GOPATH=%cd%
cd src\service
go get -v
cd ..
cd ..
cd src\main
go get -v
go build -o atago.exe