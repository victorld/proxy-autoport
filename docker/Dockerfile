FROM golang:1.20.11-alpine3.18

COPY ./ /app/proxy-autoport/

WORKDIR /app/proxy-autoport/

RUN chmod +x ./docker/entrypoint.sh

#ENTRYPOINT ["go", "run" ,"main/webserver.go"]
ENTRYPOINT ["./docker/entrypoint.sh"]