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

RUN apt-get install -y python3
RUN apt-get install -y gcc
RUN apt-get install -y gcc-multilib
RUN apt-get install -y gcc-mingw-w64-x86-64
RUN apt-get install -y gcc-arm-linux-gnueabi
RUN apt-get install -y binutils-arm-linux-gnueabi
RUN apt-get install -y gcc-aarch64-linux-gnu
RUN apt-get install -y binutils-aarch64-linux-gnu

# TODO: Fix arm-i386 conflict + darwin cross compile

FROM builder

RUN mkdir /build

WORKDIR /builder
COPY . /builder

CMD python3 tools/make.py
