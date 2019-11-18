FROM golang:1.13.4

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go get github.com/aws/aws-lambda-go

CMD ["app"]

RUN apt-get update -y &&\
    apt-get install curl -y &&\
    curl -sL https://deb.nodesource.com/setup_10.x | bash &&\
    apt-get install nodejs -y &&\
    node -v &&\
    npm -v &&\
    npm install -g serverless@1.52.2
