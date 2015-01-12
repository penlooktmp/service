./daemon.sh
go build socket.go
service socket stop
./socket remove
./socket install