FROM golang:1.22-alpine as build

WORKDIR /app

COPY . ./

RUN apk add --no-cache git

ENV GOPRIVATE="gitlab.com/sportgroup_hq"

ENV TOKEN=$GITHUB_TOKEN

RUN --mount=type=secret,id=GIT_AUTH_TOKEN,target=/root/gitlab-token.txt \
    git config --global url."https://$(cat /root/gitlab-token.txt):x-oauth-basic@github.com/".insteadOf "https://github.com/"

RUN go mod download
RUN go build ./cmd/server

WORKDIR /

FROM alpine

COPY --from=build /app/server /server

ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO /zoneinfo.zip

ENTRYPOINT ["/server"]
