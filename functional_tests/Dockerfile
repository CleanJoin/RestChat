FROM cypress/browsers:latest

RUN mkdir -p /usr/src/restchat_functional_tests

WORKDIR /usr/src/restchat_functional_tests

COPY package.json package.json
COPY yarn.lock yarn.lock
COPY cypress.json cypress.json
COPY cypress/ cypress/

RUN yarn install

CMD yarn cypress run
