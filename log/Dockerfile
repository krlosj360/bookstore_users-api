FROM golang:1.14.4

ENV mysql_users_username = root
ENV mysql_users_password = Medic@gile2020
ENV mysql_users_host = 127.0.0.1
ENV mysql_users_schema = medic_agile
ENV REPO_URL=app

# Setup out $GOPATH
ENV GOPATH=/app

ENV APP_PATH=$GOPATH/src/$REPO_URL

ENV APP_PATH=$GOPATH/src/$REPO_URL

# /app/src/github.com/federicoleon/bookstore_items-api/src

# Copy the entire source code from the current directory to $WORKPATH
ENV WORKPATH=$APP_PATH/src
COPY  $WORKPATH
WORKDIR $WORKPATH

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...


RUN go build -o user-api .

# Expose port 8081 to the world:
EXPOSE 3001

CMD ["./user-api"]



