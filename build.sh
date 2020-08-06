COMMIT_SHA1=$(git rev-parse --short HEAD || echo "0.0.0")
BUILD_TIME=$(date "+%F %T")
GO_VERSION=$(go version)
Version=$1

go build -o ango -ldflags "-X github.com/oldthreefeng/ango/cmd.Version=$1 -X  'github.com/oldthreefeng/ango/cmd.Goversion=${GO_VERSION}' -X github.com/oldthreefeng/ango/cmd.Githash=${COMMIT_SHA1} -X 'github.com/oldthreefeng/ango/cmd.Buildstamp=${BUILD_TIME}'" main.go