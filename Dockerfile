FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git 

RUN mkdir /paxos

WORKDIR /paxos

COPY . .

RUN go env -w GO111MODULE=off

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -o /go/bin/paxos


FROM scratch

COPY --from=builder /go/bin/paxos /go/bin/paxos

ENTRYPOINT ["/go/bin/paxos"]

EXPOSE 8080
