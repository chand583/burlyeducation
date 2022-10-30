FROM golang:1.18

RUN mkdir -p /usr/local/go/burlyed
RUN export GOROOT=/usr/local/go && \
    export GOPATH=/usr/local/go && \
    export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
RUN echo "export GOROOT=/usr/local/go"  >> /root/.bashrc && \
    echo "export GOPATH=/usr/local/go" >> /root/.bashrc && \
    echo "export PATH=$GOROOT/bin:$GOPATH/bin:$PATH" >> /root/.bashrc
ENV APPMODE=dev

WORKDIR /usr/local/go/burlyed
COPY . .
RUN ln -s app-local.conf  conf/app.conf
RUN go mod tidy
RUN go get -u github.com/beego/bee
CMD ["bee", "run"]
COPY ./docker-entrypoint.sh /
RUN chmod +x /docker-entrypoint.sh
ENTRYPOINT [ "/docker-entrypoint.sh" ]
