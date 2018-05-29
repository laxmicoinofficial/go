FROM golang:1.9
RUN bash -c "curl https://glide.sh/get | sh"
WORKDIR /go/src/github.com/laxmicoinofficial/go

COPY glide.lock /go/src/github.com/laxmicoinofficial/go
COPY glide.yaml /go/src/github.com/laxmicoinofficial/go
RUN glide install

COPY . .
RUN go install github.com/laxmicoinofficial/go/tools/...
RUN go install github.com/laxmicoinofficial/go/services/...