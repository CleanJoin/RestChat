FROM node

RUN mkdir -p /usr/src/restchat-client

WORKDIR /usr/src/restchat-client

COPY package.json package.json
COPY yarn.lock yarn.lock
COPY src/ src/
COPY public/ public/

RUN yarn install

CMD yarn start

