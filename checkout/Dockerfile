FROM ubuntu:22.04

ADD ./bin/app /app
ADD ./config.yaml /config.yaml

EXPOSE 8080
EXPOSE 8082

CMD ["/app"]

# ## Build

# FROM golang:1.19.7-alpine3.17 AS builder

# ENV CGO_ENABLED 0
# ENV GOOS linux

# RUN apk update --no-cache
# RUN apk add make

# # WORKDIR /build

# COPY go.sum .
# COPY go.mod .
# RUN go mod download
# COPY . .

# RUN make build

# ## Deploy
# FROM scratch AS final

# WORKDIR /app

# COPY --from=builder /bin/app .
# COPY --from=builder /config.yaml .

# EXPOSE 8080
# EXPOSE 8082

# CMD ["/app"]