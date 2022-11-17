# Must be in root dir of project -> i.e. empire/
go build -o ./bin/empire ./cmd/empire/main.go

chmod +x ./bin/empire

sudo mkdir -p /usr/local/empire/bin && sudo cp ./bin/empire /usr/local/empire/bin
sudo mkdir -p /usr/local/empire/stdlib && sudo cp -r ./stdlib /usr/local/empire