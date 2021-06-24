## Installation  
```choco install make ``` | ```apt-get install make```  
```pip install https://github.com/joh/when-changed/archive/master.zip```  
```aws-sam: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html```

## Development
Full list of commands are listed in Makefile 

### start aws-sam development server  
```make aws-sam```  

### watch files in ./src directory and recompile whenever they change  
Open another terminal
```make dev```

