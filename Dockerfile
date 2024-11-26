#==================================================================================
#   Copyright (c) 2023 Capgemini Intellectual Property.
#
#==================================================================================
#
#
#      Abstract:       Builds a container to compile ConflictMgr code
#      Date:           12 Feb 2023
#
###########################################################
FROM nexus3.o-ran-sc.org:10002/o-ran-sc/bldr-ubuntu20-c-go:1.0.0 as conflictMgrbuild
ARG GOVERSION="1.19"
RUN wget -nv https://dl.google.com/go/go${GOVERSION}.linux-amd64.tar.gz \
     && tar -xf go${GOVERSION}.linux-amd64.tar.gz \
     && mv go /opt/go/${GOVERSION} \
     && rm -f go*.gz

# Install RMr shared library
ARG RMRVERSION=4.9.4
RUN wget --content-disposition https://packagecloud.io/o-ran-sc/release/packages/debian/stretch/rmr_${RMRVERSION}_amd64.deb/download.deb && dpkg -i rmr_${RMRVERSION}_amd64.deb && rm -rf rmr_${RMRVERSION}_amd64.deb
# Install RMr development header files
RUN wget --content-disposition https://packagecloud.io/o-ran-sc/release/packages/debian/stretch/rmr-dev_${RMRVERSION}_amd64.deb/download.deb && dpkg -i rmr-dev_${RMRVERSION}_amd64.deb && rm -rf rmr-dev_${RMRVERSION}_amd64.deb


ENV DEFAULTPATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
ENV PATH=$DEFAULTPATH:/usr/local/go/bin:/opt/go/${GOVERSION}/bin:/root/go/bin

RUN mkdir -p /go/src/conflictmgr
WORKDIR /go/src/conflictmgr
COPY config config
COPY constants constants
COPY main main
COPY procedures procedures
COPY go.sum go.sum
COPY go.mod go.mod
COPY conflictCache conflictCache
RUN go mod vendor
RUN go build -o conflictmgr  main/main.go

FROM ubuntu:20.04
RUN mkdir /conflict
COPY  --from=conflictMgrbuild /go/src/conflictmgr/conflictmgr /conflict
WORKDIR "/conflict"
ENV PLT_NAMESPACE="ricplt"
COPY --from=conflictMgrbuild /usr/local/include /usr/local/include
COPY --from=conflictMgrbuild /usr/local/lib /usr/local/lib
RUN ldconfig
ENTRYPOINT ["./conflictmgr","-f","../cfg/config.yaml"]
