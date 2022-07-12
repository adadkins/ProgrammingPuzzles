FROM golang:1.18-alpine as builder

# RUN go install golang.org/dl/go1.18@latest \
#     && go1.18 download

WORKDIR /build

COPY go.mod .
COPY go.sum .
COPY ./url_shortener .

RUN go mod tidy

RUN go build -o ./app .

ENTRYPOINT ["/build/app"]