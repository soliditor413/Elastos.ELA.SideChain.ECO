# Build Geth in a stock Go builder container
FROM golang:1.13-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

ADD . /Elastos.ELA.SideChain.ECO
RUN cd /Elastos.ELA.SideChain.ECO && make geth bootnode

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /Elastos.ELA.SideChain.ECO/build/bin/* /usr/local/bin/

EXPOSE 20656 20655 8547 20658 20658/udp
#ENTRYPOINT ["geth"]
