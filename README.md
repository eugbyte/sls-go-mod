# About
serverless lambda in golang with aws dynamodb local 

# Installation  
<ins>make</ins>: `choco install make` | `apt-get install make`  
<ins>file-watcher</ins>: `pip install https://github.com/joh/when-changed/archive/master.zip`  
<ins>aws-sam</ins>: `aws-sam: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html`  
<ins>docker</ins>: `choco install docker-desktop` | `sudo apt-get install docker-ce docker-ce-cli containerd.io`  
<ins>dynamodb-admin</ins>: `npm install dynamodb-admin -g`

# Development
Full list of commands are listed in Makefile 
If you are on windows, you need to have `git bash` cli installed to run the commands1

## start aws-sam development server  
`make aws-sam`  

## watch files in ./src directory and recompile whenever they change  
Open another terminal  
`make watch`

### note
Only file changes in the src directory is detected.  
Also note that if you change the sam-template.yml file, you will have to restart the aws-sam development server too

## start local dynamodb server
Open another terminal  
`make db`
### Note: Connecting aws-sam to dynamodb local
1. Need to ensure that aws-sam and dynamodb local run on the [same docker network](https://stackoverflow.com/questions/48926260/connecting-aws-sam-local-with-dynamodb-in-docker)
```
sam local start-api --docker-network local-network
```

```
# docker-compose.yml for amazon/dynamodb-local
networks:
  backend:
    name: local-network

services:
  dynamodb:
    image: amazon/dynamodb-local
    ports:
      - 1234:8000
    networks: 
      - backend
```

2. Need to configure the aws sdk client to point to the local dynamodb endpoint, e.g.
```
var config = aws.NewConfig().WithEndpoint("http://localhost:1234").
var Client = dynamodb.New(..., config)
``` 
If running docker on windows, need to speicfy the endpoint as `http://host.docker.internal:1234` instead of `http://localhost:1234`. This is due to the fact that the host has a [changing IP address](https://docs.docker.com/docker-for-windows/networking/). 

### Inspect dynamodb with dynamodb-admin
After `make db`, open the browser, and go to http://localhost:8001

### stop dynamodb local
`make stop-db`

# Starting with an empty project template
```
npm install -g serverless
serverless create --template aws-go-mod --path my-folder
```
Then replace the `serverless.yml` with the `sam-template.yml` 
serverless framework does not yet fully support golang development, e.g. no sls offline dev server  
Regardless, the generated template is helpful
