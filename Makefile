.PHONY: build clean deploy gomodgen

aws-sam:
	AWS_REGION=ap-southeast-1
	sam local start-api --template sam-template.yml --docker-network local-network --region ap-southeast-1  --warm-containers LAZY  --debug

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh

build: 
	# gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -o bin/hello -ldflags="-s -w" src/handlers/hello/main.go
	env GOOS=linux go build -o bin/create -ldflags="-s -w" src/handlers/create/main.go

dev-watch: 
	make build
	when-changed -r "src" make build

start-db:
	docker-compose -f src/data/docker-compose.yaml up -d

stop-db:
	docker-compose -f src/data/docker-compose.yaml down

db-admin:
	@if [ -z `which dynamodb-admin 2> /dev/null` ]; then \
			echo "Need to install dynamodb-admin, execute \"npm install dynamodb-admin -g\"";\
			exit 1;\
	fi
	DYNAMO_ENDPOINT=http://localhost:18000 dynamodb-admin

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose

create-table:
	aws dynamodb create-table --cli-input-json file://src/data/create_book_table.json --endpoint-url http://localhost:18000

seed-data:
	aws dynamodb batch-write-item --request-items file://src/data/seed_book_table.json --endpoint-url http://localhost:18000

