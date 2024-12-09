# image for compiling binary
ARG BUILDER_IMAGE="golang:1.23-alpine"
ARG RUNNER_IMAGE="alpine:3.19"
ARG WORKPATH="/go/src/github.com/Reshnyak/innopolis/AssistTodo"
FROM ${BUILDER_IMAGE} AS builder
RUN apk --no-cache --update add git make gcc openssh g++
# coping project files
WORKDIR /go/src/github.com/Reshnyak/innopolis/AssistTodo
# copy go mod
COPY go.mod go.sum ./
# Get dependancies. Also will be cached if we wont't change mod/sum
RUN go mod download
# Build
COPY . .
RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.3.0
RUN go generate
RUN go build -o ./bin/alise-todo ./cmd/main.go
#RUN make build

FROM ${RUNNER_IMAGE}

COPY --from=builder /go/src/github.com/Reshnyak/innopolis/AssistTodo/config ./config
COPY --from=builder /go/src/github.com/Reshnyak/innopolis/AssistTodo/docs ./docs
COPY --from=builder /go/src/github.com/Reshnyak/innopolis/AssistTodo/bin/alise-todo /

CMD ["./alise-todo",""]