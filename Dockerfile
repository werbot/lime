# Build Stage
FROM golang:1.14 AS build-stage

LABEL app="build-lime"
LABEL REPO="https://github.com/werbot/lime"

ENV PROJPATH=/go/src/github.com/lacion/lime

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/lacion/lime
WORKDIR /go/src/github.com/lacion/lime

RUN make build-alpine



# Final Stage
FROM alpine:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/werbot/lime"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/lime/bin

WORKDIR /opt/lime/bin

COPY --from=build-stage /go/src/github.com/werbot/lime/bin/lime /opt/werbot/bin/
RUN chmod +x /opt/lime/bin/lime

# Create appuser
RUN adduser -D -g '' lime
USER lime

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/lime/bin/lime"]
