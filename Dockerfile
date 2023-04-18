#TODO add healthcheck

# Stage 1: build the project

FROM golang:1.20 as build-stage

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download

COPY . ./


RUN CGO_ENABLED=0 GOOS=linux go build -o /goversion

# Stage 2: build the image

FROM alpine:latest  
RUN apk --no-cache add ca-certificates libc6-compat
WORKDIR /app/
COPY --from=build-stage /goversion .
EXPOSE 8000
CMD ["./goversion"]  


# # old version

# FROM golang:1.20

# WORKDIR /app


# COPY go.mod go.sum ./
# RUN go mod download

# COPY . ./


# RUN go build -o /goversion

# EXPOSE 8000


# CMD ["/goversion"]