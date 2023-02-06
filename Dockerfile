FROM alpine:3.16.3

WORKDIR /

COPY .env ./
COPY ./bin/app-386-linux ./

EXPOSE 9998

CMD [ "/app-386-linux" ]