FROM golang:1.23.3-alpine3.20

WORKDIR /app

ENV CGO_ENABLED 0
ENV GOPATH /go
ENV GOCACHE /go-build

COPY . .

RUN go install -v golang.org/x/tools/gopls@latest && \
  go install -v github.com/fatih/gomodifytags@v1.16.0 && \
  go install -v github.com/josharian/impl@v1.1.0 && \
  go install -v github.com/haya14busa/goplay/cmd/goplay@v1.0.0 && \
  go install -v honnef.co/go/tools/cmd/staticcheck@latest && \
  go install -v github.com/go-delve/delve/cmd/dlv@latest && \
  go install -v golang.org/x/tools/cmd/goimports@latest && \
  go install -v github.com/air-verse/air@latest
