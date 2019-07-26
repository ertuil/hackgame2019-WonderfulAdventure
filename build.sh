go-bindata -o=./asset/asset.go -pkg=asset statics/...
go build -o public/wa-darwin WonderfulAdventure
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/wa-linux WonderfulAdventure
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/wa-windows.exe WonderfulAdventure
cd public &&  ./wa-darwin
