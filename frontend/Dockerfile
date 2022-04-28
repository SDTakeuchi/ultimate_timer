FROM node:16.14-slim

WORKDIR /app

RUN apt-get update \
&&  rm -rf node_modules/ 

COPY ./ultimate_timer/ .

RUN yarn install

CMD [ "yarn", "dev" ]
# CMD [ "yarn", "build", "&&", "yarn", "start"]
