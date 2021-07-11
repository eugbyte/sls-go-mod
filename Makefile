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
	env GOOS=linux go build -o bin/scanAll -ldflags="-s -w" src/handlers/scanAll/main.go
	env GOOS=linux go build -o bin/getItem -ldflags="-s -w" src/handlers/getItem/main.go

watch: 
	make build
	when-changed -r "./src" make build

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose

db-start:
	docker-compose -f src/data/seed/docker-compose.yaml up -d

db-stop:
	docker-compose -f src/data/seed/docker-compose.yaml down

db-create-table:
	aws dynamodb create-table --cli-input-json file://src/data/seed/create_book_table.json --endpoint-url http://localhost:18000 >/dev/null 2>&1

db-seed-data:
	aws dynamodb batch-write-item --request-items file://src/data/seed/seed_book_table.json --endpoint-url http://localhost:18000

db-admin:
	@if [ -z `which dynamodb-admin 2> /dev/null` ]; then \
			echo "Need to install dynamodb-admin, execute \"npm install dynamodb-admin -g\"";\
			exit 1;\
	fi
	DYNAMO_ENDPOINT=http://localhost:18000 dynamodb-admin

db:
	make db-stop || echo "db already stopped"
	make db-start
	make db-create-table 
	make db-seed-data
	make db-admin
