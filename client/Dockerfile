FROM node

CMD mkdir -p /usr/src/restchat-client

WORKDIR /usr/src/restchat-client

COPY restchat-client/package.json package.json
COPY restchat-client/yarn.lock yarn.lock
COPY restchat-client/src src/
COPY restchat-client/public public/

RUN yarn install

CMD yarn start

