FROM golang:1.12.13 as build
WORKDIR /app

# Set environment for build
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Copy files and pull vendor
COPY . .

RUN go mod download

# Build to binrary
RUN go build -a -ldflags "-s -w" -v -o main .

# Optimize docker image after build
FROM alpine:3.10

# Add non root user for security context
RUN addgroup -S app && adduser -S -g app app 

WORKDIR /app

COPY --from=build /app/main .

RUN chown -R app /app

# Use with app instead root
USER app

EXPOSE 8888

CMD ["./main"]