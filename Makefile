.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh

sam-local:
	sam local start-api --template sam-template.yml

invoke-hello:
	make build	
	sls invoke local --function hello --data '{"body":{"message":"Hi!"}}'

