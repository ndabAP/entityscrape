default: build_web build_srv build_cli

build_web: 
	cd web; yarn build

build_srv:
	go build -o bin/entityscrape-srv cmd/entityscrape-srv/main.go

build_cli:
	go build -o bin/entityscrape-cli cmd/entityscrape-cli/main.go