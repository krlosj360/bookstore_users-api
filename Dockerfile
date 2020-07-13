FROM golang:1.14.4

ENV REPO_URL=bookstore_users-api

# SETUP OUT $GOPATH


ENV GOPATH=/go

ENV APP_PATH =$GOPATH/src/$REPO_URL

# Copy the entire source code from the current directory to $WORKPATH
ENV WORKPATH=$APP_PATH/src
COPY src $WORKPATH
WORKDIR $WORKPATH

#WORKDIR /go/src/bookstore_users-api
#COPY . .

#RUN go get -d -v ./...
#RUN go install -v ./...

RUN go build -o user-api .

# Expose port 8081 to the world:
EXPOSE 3001

CMD ["./user-api"]



