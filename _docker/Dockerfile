FROM golang:1.20.2-bullseye

# realize - 自動ビルド用のディレクトリ監視ツールインストール
WORKDIR /root

WORKDIR /usr/src
# GitHub Private Access
RUN git config --global url."https://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:x-oauth-basic@github.com/".insteadOf "https://github.com/"

# supervisor
RUN apt-get update -y && apt-get install supervisor -y
COPY ./supervisor/* /etc/supervisor/conf.d/
CMD ["supervisord", "-n", "-c", "/etc/supervisor/supervisord.conf"]