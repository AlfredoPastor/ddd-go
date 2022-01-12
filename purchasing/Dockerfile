FROM ddd-go_go-builder AS go_builder

WORKDIR /go/src/github.com/AlfredoPastor/ddd-go/purchasing

ADD cmd cmd
ADD internal internal

RUN CGO_ENABLED=0 go build -o /go/bin/purchasing cmd/main.go

# ---------------------------------------------------------------------------- #
#                                 create image                                 #
# ---------------------------------------------------------------------------- #

FROM scratch

COPY --from=go_builder /go/bin/purchasing /purchasing

CMD ["/purchasing"]
