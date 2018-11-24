FROM docker:18.06-dind

############
## GOLANG ##
############

RUN apk add --no-cache git make musl-dev go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

WORKDIR $GOPATH

##############
## FINALIZE ##
##############

CMD []
