FROM golang:1.18-alpine AS BUILDER

WORKDIR /rss3-pregod
COPY . .

RUN apk add \
    gcc \
    g++ \
    git

ARG GH_USER
ARG GH_TOKEN

ENV GOPRIVATE="github.com/naturalselectionlabs/kurora"
ENV GH_USER=$GH_USER
ENV GH_TOKEN=$GH_TOKEN
RUN git config --global url."https://${GH_USER}:${GH_TOKEN}@github.com".insteadOf "https://github.com"

RUN go build -o dist/hub ./service/hub/cmd
RUN go build -o dist/indexer ./service/indexer/cmd
RUN go build -o dist/crawler ./service/crawler/cmd
RUN go build -o dist/beat ./service/beat/cmd

FROM alpine:latest AS RUNNER

COPY --from=builder /rss3-pregod/dist/* /usr/local/bin/
COPY --from=builder /rss3-pregod/deploy/config/* /etc/pregod/
