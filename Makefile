.PHONY: build clean deploy gomodgen

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh

aws-sam:
	sam local start-api --template sam-template.yml --docker-network local-network --region ap-southeast-1  --debug

build: 
	export GO111MODULE=on
	env GOOS=linux go build -o bin/hello -ldflags="-s -w" src/handlers/hello/main.go
	env GOOS=linux go build -o bin/create -ldflags="-s -w" src/handlers/create/main.go
	env GOOS=linux go build -o bin/scan -ldflags="-s -w" src/handlers/scan/main.go

watch: 
	make build
	when-changed -r "./src" make build

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose

start-db:
	docker-compose -f src/data/docker-compose.yaml up -d

stop-db:
	docker-compose -f src/data/docker-compose.yaml down

create-table:
	aws dynamodb create-table --cli-input-json file://src/data/create_book_table.json --endpoint-url http://localhost:18000 >/dev/null 2>&1

seed-data:
	aws dynamodb batch-write-item --request-items file://src/data/seed_book_table.json --endpoint-url http://localhost:18000

db-admin:
	@if [ -z `which dynamodb-admin 2> /dev/null` ]; then \
			echo "Need to install dynamodb-admin, execute \"npm install dynamodb-admin -g\"";\
			exit 1;\
	fi
	DYNAMO_ENDPOINT=http://localhost:18000 dynamodb-admin

db:
	make stop-db || echo "db already stopped"
	make start-db
	make create-table 
	make seed-data
	make db-admin
