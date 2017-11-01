FROM partlab/ubuntu

MAINTAINER Régis Gaidot <regis@partlab.co>

ENV DEBIAN_FRONTEND noninteractive
ENV INITRD No
ENV LANG en_US.UTF-8
ENV GOVERSION 1.9
ENV GOROOT /opt/go
ENV GOPATH /root/.go

RUN cd /opt && wget https://storage.googleapis.com/golang/go${GOVERSION}.linux-amd64.tar.gz && \
    tar zxf go${GOVERSION}.linux-amd64.tar.gz && rm go${GOVERSION}.linux-amd64.tar.gz && \
    ln -s /opt/go/bin/go /usr/bin/ && \
    mkdir $GOPATH

RUN go get github.com/gorilla/websocket
RUN cd /root/.go/src/github.com/ && mkdir carrot-ar
RUN cd /root/.go/src/github.com/carrot-ar/ && git clone -b fix_data_race https://github.com/carrot-ar/carrot
RUN cd /root/.go/src/github.com/ && mkdir clandry94
RUN cd /root/.go/src/github.com/clandry94/ && git clone https://github.com/clandry94/echo

RUN cp /root/.go/src/github.com/clandry94/echo/home.html /home.html

CMD ["/usr/bin/go"]
ENTRYPOINT ["/usr/bin/go", "run", "/root/.go/src/github.com/clandry94/echo/echo.go"]
