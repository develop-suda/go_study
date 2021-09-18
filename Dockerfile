# ベースとなるDockerイメージ指定
FROM golang:1.17 as build

ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV GOBIN=$GOPATH/bin
ENV PATH $PATH:$GOROOT:$GOPATH:$GOBIN
ENV GO111MODULE=on

# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/youtub-board-do
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/youtub-board-do
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/youtub-board-do