FROM golang:alpine AS builder
RUN apk update && apk add --no-cache 'git=~2'

# Install dependencies
ENV GO111MODULE=on
WORKDIR $GOPATH/src/eztrust


# Fetch dependencies.
# Using go get.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/main cmd/main.go

############################
# STEP 2 build a small image
############################
FROM alpine:3

RUN apk add tzdata
RUN ln -s /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime

WORKDIR /

# Copy our static executable.
COPY --from=builder /go/main /go/main
COPY --from=builder /go/src/eztrust/.env /go/.env
COPY --from=builder /go/src/eztrust/template /go/template



ENV PORT 9090
ENV GIN_MODE release
EXPOSE 9090

WORKDIR /go

# Run the Go Gin binary.
ENTRYPOINT ["/go/main"]
