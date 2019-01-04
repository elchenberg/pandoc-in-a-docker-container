FROM golang:1.11.4-alpine3.8 as builder
SHELL ["/bin/ash", "-o", "pipefail", "-c"]
RUN addgroup -g 1000 -S pandoc && adduser -u 1000 -S pandoc -G pandoc
ARG PANDOC_VERSION=2.5
RUN wget https://github.com/jgm/pandoc/releases/download/${PANDOC_VERSION}/pandoc-${PANDOC_VERSION}-linux.tar.gz --output-document=- | tar --directory=/bin/ --strip-components=2 --extract --gzip --verbose --file=- pandoc-${PANDOC_VERSION}/bin/
COPY cmd/server/ $GOPATH/cmd/server/
WORKDIR $GOPATH/cmd/server/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /bin/server

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /bin/pandoc /bin/pandoc
COPY --from=builder /bin/pandoc-citeproc /bin/pandoc-citeproc
COPY --from=builder /bin/server /bin/server
USER pandoc
ENTRYPOINT [ "server" ]