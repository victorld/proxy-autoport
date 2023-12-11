FROM golang:1.20.11-alpine3.18

#COPY ./docker/ /app/proxy-autoport/

COPY api /app/proxy-autoport/api
COPY client /app/proxy-autoport/client
COPY cons /app/proxy-autoport/cons
COPY docker /
COPY main /app/proxy-autoport/main
COPY plugin /app/proxy-autoport/plugin
COPY resource /app/proxy-autoport/resource
COPY route /app/proxy-autoport/route
COPY server /app/proxy-autoport/server
COPY tools /app/proxy-autoport/tools
COPY vendor /app/proxy-autoport/vendor
COPY config.yaml /app/proxy-autoport/
COPY go.mod /app/proxy-autoport/

WORKDIR /app/proxy-autoport/

RUN chmod +x /entrypoint.sh

#ENTRYPOINT ["go", "run" ,"main/webserver.go"]
ENTRYPOINT ["/entrypoint.sh"]