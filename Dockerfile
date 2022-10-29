

# Alpine is chosen for its small footprint
# compared to Ubuntu

FROM golang:1.19-alpine

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...



# # Add Maintainer info
# LABEL maintainer="Arvid Wedtstein"

# WORKDIR /app

# RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# COPY go.mod ./
# COPY go.sum ./

# # RUN go mod download<

# COPY *.go ./

# RUN go install
# # Expose port 8080 to the outside world
EXPOSE 8080

RUN go build -o /githubembedapi

CMD [ "/githubembedapi" ]


