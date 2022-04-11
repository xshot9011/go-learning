# FROM golang:1.12-alpine AS build_base

# RUN apk add --no-cache git

# WORKDIR /app/src

# COPY go.mod .
# COPY go.sum .

# RUN go mod download

# COPY . .

# RUN CGO_ENABLED=0 go test -v

# RUN go build -o ./out/go-sample-app .

# FROM alpine:3.9
# RUN apk add ca-certificates

# COPY --from=build_base /tmp/go-sample-app/out/go-sample-app /app/go-sample-app

# EXPOSE 8080
# CMD ["/app/go-sample-app"]
