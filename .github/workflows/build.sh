#!/bin/bash

platforms=("linux/amd64" "linux/arm" "linux/arm64" "darwin/amd64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "-X github.com/alecbcs/lookout/cli.appVersion=$1" -o lookout-$1-${GOOS}-${GOARCH} .

done