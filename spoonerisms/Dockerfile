FROM node:10

ADD src /app
ADD package.json /app/package.json
ADD yarn.lock /app/yarn.lock

RUN cd /app && yarn install

ENTRYPOINT [ "node", "/app/index.js" ]
