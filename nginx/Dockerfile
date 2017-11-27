FROM node:alpine

RUN mkdir -p /tmp/app

RUN apk upgrade -U \
    && apk add nginx \
    && rm -rf /var/cache/* \
    && mkdir -p /run/nginx \
    && ln -sf /dev/stdout /var/log/nginx/access.log \
    && ln -sf /dev/stderr /var/log/nginx/error.log

ADD ./frontend/package.json /tmp/app
RUN cd /tmp/app && npm install --dev
ADD ./frontend/ /tmp/app

WORKDIR /tmp/app

RUN npm run build

RUN mkdir -p /src/app && mv /tmp/app/dist /src/app
RUN rm -rf /tmp/app

COPY ./nginx/nginx.conf /etc/nginx/nginx.conf

STOPSIGNAL SIGTERM

CMD ["nginx"]