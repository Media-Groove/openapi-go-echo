FROM golang:1.15-stretch

# realize - 自動ビルド用のディレクトリ監視ツールインストール
WORKDIR /root
RUN go get github.com/oxequa/realize

WORKDIR /usr/src
# GitHub Private Access
RUN git config --global url."https://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:x-oauth-basic@github.com/".insteadOf "https://github.com/"
# 起動スクリプト実行
COPY ./startup.sh /startup.sh
RUN chmod 744 /startup.sh
CMD /startup.sh
