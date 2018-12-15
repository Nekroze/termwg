FROM golang:1 AS build

ARG PROJECT=termwg
ENV PROJECT="${PROJECT}"

WORKDIR "$GOPATH/src/github.com/Nekroze/${PROJECT}"

# Tools
RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b "$GOPATH/bin" v1.12.2 \
 && go get -u github.com/kyoh86/richgo
# Go 1.11+ modules
ENV GO111MODULE=on

# Deps
COPY go.mod go.sum ./
RUN go get -v -d ./...

# Check and compile everything
COPY main.go ./main.go
COPY cmd ./cmd
RUN richgo test -v ./... \
 && golangci-lint run --deadline '2m' --enable-all --disable gochecknoglobals,gochecknoinits \
 && CGO_ENABLED=0 GOOS=linux GOARCH=386 go build \
    -a -installsuffix cgo -ldflags='-w -s' -o "/usr/bin/${PROJECT}" -v \
    .

FROM nekroze/containaruba:alpine AS test
CMD ["--order=random"]

COPY --from=build "/usr/bin/${PROJECT}" "/usr/bin/${PROJECT}"
COPY ./tests/features /usr/src/app/features


FROM scratch AS final

ARG PROJECT=termwg
ENV PROJECT="${PROJECT}"

COPY --from=build "/usr/bin/${PROJECT}" /app
ENTRYPOINT ["/app"]
