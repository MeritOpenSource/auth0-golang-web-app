FROM golang:1.23-alpine3.20 AS builder

# Define current working directory
WORKDIR /01-Login

# Download modules to local cache so we can skip re-
# downloading on consecutive docker build commands
COPY go.mod .
COPY go.sum .
RUN go mod download

# Add sources
COPY . .

RUN go build -o out/auth0-go-web-app .

FROM alpine:3.20 AS app

ENV HOME=/

# Copy the binary from the builder stage
COPY --from=builder /01-Login/out/auth0-go-web-app /auth0-go-web-app

# Copy the website files
COPY --from=builder /01-Login/web/static /web/static
COPY --from=builder /01-Login/web/template /web/template

# Expose port 3000 for our web app binary
EXPOSE 3000

CMD ["/auth0-go-web-app"]
