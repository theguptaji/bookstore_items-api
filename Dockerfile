# Start from base image 1.12.13
FROM golang:1.12.13

# Configure our repo url so we can configure work directory
ENV REPO_URL=github.com/theguptaji/bookstore_items-api
# Setup our $GOPATH
ENV GOPATH=/app
ENV APP_PATH=$GOPATH/src/$REPO_URL

# final path : /app/src/github.com/theguptaji/bookstore_items-api/src

# Copy the entire source code from current directory to $WORKPATH
ENV WORKPATH=$APP_PATH/src
COPY src $WORKPATH
WORKDIR $WORKPATH

RUN go build -o items-api .

# Expose port 8081
EXPOSE 8081

CMD ["./items-api"]