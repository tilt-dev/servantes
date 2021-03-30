FROM node:10-alpine

ADD package.json /app/package.json
ADD yarn.lock /app/yarn.lock
RUN cd /app && yarn install

ADD src /app

ENTRYPOINT [ "node", "/app/index.js" ]
