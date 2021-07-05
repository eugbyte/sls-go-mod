.PHONY: build clean deploy gomodgen

aws-sam:
	sam local start-api --template sam-template.yml 

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh

build: 
	# gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -o bin/hello -ldflags="-s -w" src/handlers/hello/main.go
	env GOOS=linux go build -o bin/create -ldflags="-s -w" src/handlers/create/main.go

dev: 
	make build
	when-changed -r "src" make build

start-db:
	docker-compose -f src/data/docker-compose.yaml up -d

stop-db:
	docker-compose -f src/data/docker-compose.yaml down

db-admin:
	DYNAMO_ENDPOINT=http://localhost:18000 dynamodb-admin

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose



