FROM ubuntu:22.04 as builder

RUN apt-get update
RUN apt-get install -y wget

RUN wget \
  -P /tmp \
  --no-verbose --show-progress \
  --progress=bar:force:noscroll \
  https://go.dev/dl/go1.19.linux-amd64.tar.gz 2>&1
RUN tar -C /usr/local -xzf /tmp/go1.19.linux-amd64.tar.gz
RUN rm /tmp/go1.19.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

RUN apt-get install -y python3 gcc gcc-mingw-w64

FROM builder

RUN mkdir /build

WORKDIR /builder
COPY . /builder

CMD python3 tools/make.py
