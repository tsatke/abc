FROM docker:18.06-dind

############
## GOLANG ##
############

RUN apk add --no-cache git make musl-dev curl
RUN curl -O https://storage.googleapis.com/golang/go1.11.2.linux-amd64.tar.gz && \
    tar -C /usr/local -xvf go1.11.2.linux-amd64.tar.gz
RUN echo export PATH=$PATH:/usr/local/go/bin >> $HOME/.profile

# Configure Go
ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH ${GOROOT}/bin:${GOPATH}/bin:$PATH

##############
## FINALIZE ##
##############

CMD []
