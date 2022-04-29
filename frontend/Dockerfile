FROM node:18-alpine3.15

WORKDIR /usr/app/

RUN apk update \
&&  rm -rf node_modules/ \
&&  rm -rf /var/lib/apt/lists/*

COPY ./ultimate_timer/ .

RUN yarn install

CMD [ "yarn", "dev" ]
# CMD [ "yarn", "build", "&&", "yarn", "start"]
