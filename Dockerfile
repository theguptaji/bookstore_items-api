# Start from golang base image
FROM golang:1.14

# Configure our repo url so we can configure work directory
ENV REPO_URL=github.com/theguptaji/bookstore_items-api
# Setup our $GOPATH
ENV GOPATH=/app
ENV APP_PATH=$GOPATH/src/$REPO_URL
ENV GO111MODULE=on
# final path : /app/src/github.com/theguptaji/bookstore_items-api/src

# Copy the entire source code from current directory to $WORKPATH
ENV WORKPATH=$APP_PATH/src
COPY src $WORKPATH
COPY go.mod $APP_PATH
COPY go.sum $APP_PATH

WORKDIR $APP_PATH

RUN go mod download
#RUN go get -u ./...
RUN go build -o items-api ./src/main.go

# Expose port 8081
EXPOSE 8081

CMD ["./items-api"]