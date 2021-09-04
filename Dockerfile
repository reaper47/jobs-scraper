# STAGE 1: prepare
FROM golang AS prepare

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /source 

COPY go.mod .
COPY go.sum .
RUN go mod download

# STAGE 2: build
FROM prepare AS build

COPY . .

RUN go build -o main .

WORKDIR /dist 

RUN cp /source/main .  

# STAGE 3: run
FROM scratch AS run

COPY --from=build /dist/main /
COPY --from=build /source/.env /

EXPOSE 3001

ENTRYPOINT ["/main"]