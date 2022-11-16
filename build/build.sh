# Must be in root dir of project -> i.e. empire/
go build -o ./bin/empire ./cmd/empire/main.go

chmod +x ./bin/empire

sudo cp ./bin/empire /usr/local/bin