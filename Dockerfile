# Builder builds the service as a static binary.
FROM siuyin/go:dev AS builder

COPY . /src
WORKDIR /src
RUN go mod download

# root privilege needed as siuyin/go:dev runs as a normal user.
USER root
RUN CGO_ENABLED=0 go build -o /app cmd/hello/main.go

# Deployment image is built from scatch, an empty image.
FROM scratch
COPY --from=builder /app .
COPY --from=builder /src/testdata /testdata
EXPOSE 8080
CMD ["/app"]
