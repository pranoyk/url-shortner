# url-shortner

## download source

    cd $GOPATH/src
    git clone git@github.com:pranoyk/url-shortner.git


## install dependecies

    go mod vendor

## run tests

    go test ./...

## run application

    go run .

## build and run application

    go build
    ./main

## to shorten a url
    curl http:localhost:8080/api/url-shortner/v1/shorten --data {"url":"your-url"}

## run using docker image

    docker pull pranoyk/url-shortner:latest
    docker run -p 8080:8080 pranoyk/url-shortner:1.0.0