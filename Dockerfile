FROM alpine

WORKDIR /app

RUN mkdir /app/bin
RUN mkdir /app/core
RUN mkdir /app/ext
RUN mkdir /app/views
RUN mkdir /app/public

COPY ./bin/waterway /app/bin
COPY ./bin/cli_amd /app/bin
COPY ./core /app/core
COPY ./ext /app/ext
COPY ./views /app/views
COPY ./public /app/public

ENV WATERWAY_ENV=production
ENV WATERWAY_PORT=2000

ENV APP_PWD=/app


ENV AW_ASSET_VERSION=1
ENV TZ="Asia/Shanghai"

EXPOSE 2000

CMD ["/app/bin/waterway"]
