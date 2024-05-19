FROM golang:1.22.3 as build
WORKDIR /workdir
COPY go.mod go.sum /workdir/
RUN go mod download

COPY . /workdir
RUN go build -ldflags='-s -w' -o app .

RUN cp /workdir/app /app
ENTRYPOINT ["/app"]
