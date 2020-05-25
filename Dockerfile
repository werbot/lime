# Build Stage
FROM golang:1.14 AS build-stage

LABEL REPO="https://github.com/werbot/lime"

ENV PROJPATH=/go/src/github.com/werbot/lime

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/werbot/lime
WORKDIR /go/src/github.com/werbot/lime

RUN make build-alpine



# Final Stage
FROM alpine:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/werbot/lime"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

WORKDIR /opt/bin

COPY --from=build-stage /go/src/github.com/werbot/lime/bin/lime /opt/bin/
RUN chmod +x /opt/bin/lime

# Create appuser
RUN adduser -D -g '' lime
USER lime

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/bin/lime"]
