appName=$1

echo "build to mac."
go env -w CGO_ENABLED=1 GOOS=darwin GOARCH=amd64
go build -o $appName
chmod 755 $appName