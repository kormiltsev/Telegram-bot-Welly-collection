FROM alpine:3.16.3

WORKDIR /

COPY .env ./
COPY pic.png ./
COPY ./bin/welly-linux ./

EXPOSE 9998

CMD [ "/welly-linux" ]