osType=$1
appName="gen4gonet"

if [ -z $osType ]; then
    osType="linux"
fi

if [ $osType = mac ]; then
    echo "Choose: mac"
    ./build-mac.sh $appName
elif [ $osType = win ]; then
    echo "Choose: win"
    ./build-win.sh $appName
else
    echo "Choose: linux"
    ./build-linux.sh $appName
fi

# back to mac.
go env -w CGO_ENABLED=1 GOOS=windows GOARCH=amd64
echo "Back to win."

echo "Finish build."
