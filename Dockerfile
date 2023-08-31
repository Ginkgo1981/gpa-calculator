FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY app/gpa/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/gpa app/gpa/gpa.go


FROM scratch

WORKDIR /app
COPY --from=builder /app/gpa /app/gpa
COPY --from=builder /app/etc /app/etc

CMD ["./gpa", "-f", "etc/gpa-api.yaml"]
