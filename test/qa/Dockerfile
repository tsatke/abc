FROM docker:18.06-dind

############
## GOLANG ##
############

RUN apk add --no-cache git make musl-dev curl && \
    curl -O https://storage.googleapis.com/golang/go1.11.2.linux-amd64.tar.gz && \
    tar -xvf go1.11.2.linux-amd64.tar.gz  &&\
    mv go /

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

WORKDIR $GOPATH

##############
## FINALIZE ##
##############

CMD []
