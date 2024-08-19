ARG IMAGE_NAME

FROM ${IMAGE_NAME} as base

ARG PORT
ARG APP_NAME
ARG SERVICE

ENV APP_NAME=${APP_NAME}
ENV SERVICE=${SERVICE}
ENV PORT=${PORT}

WORKDIR /${APP_NAME}
COPY . .

FROM base as local

WORKDIR /
RUN apk --update add --no-cache bash gcc musl-dev git mercurial openssh curl && \
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

WORKDIR /${APP_NAME}

CMD go mod init grib
CMD air -c ./config/$SERVICE/.air.toml

EXPOSE ${PORT}
