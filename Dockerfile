# Generic builder:
FROM ubuntu:22.04 as builder-base

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

RUN mkdir /build

WORKDIR /builder

CMD python3 tools/make.py

# Platform-specific builders:
FROM builder-base as builder-linux-amd64
ENV OS linux
ENV ARCH amd64
COPY . /builder

FROM builder-linux-amd64 as builder-linux-i386
ENV ARCH i386

FROM builder-base as builder-windows-amd64
RUN apt-get install -y gcc-mingw-w64-x86-64
ENV OS windows
ENV ARCH amd64
COPY . /builder

FROM builder-base as builder-windows-i386
RUN apt-get install -y gcc-mingw-w64-i686
ENV OS windows
ENV ARCH i386
COPY . /builder

FROM builder-base as builder-linux-arm
RUN apt-get install -y gcc-arm-linux-gnueabi
RUN apt-get install -y binutils-arm-linux-gnueabi
ENV OS linux
ENV ARCH arm
COPY . /builder

FROM builder-base as builder-linux-arm64
RUN apt-get install -y gcc-aarch64-linux-gnu
RUN apt-get install -y binutils-aarch64-linux-gnu
ENV OS linux
ENV ARCH arm64
COPY . /builder

# TODO: Fix darwin cross compile
