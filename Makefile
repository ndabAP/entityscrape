default: build_web build_srv build_cli

dev_web:
	cd web; yarn serve

dev_srv:
	PORT=3000 go run cmd/entityscrape-srv/main.go

dev_cli:
	go run cmd/entityscrape-cli/main.go

build_web: 
	cd web; yarn build

build_srv:
	go build -o bin/entityscrape-srv cmd/entityscrape-srv/main.go

build_cli:
	go build -o bin/entityscrape-cli cmd/entityscrape-cli/main.go