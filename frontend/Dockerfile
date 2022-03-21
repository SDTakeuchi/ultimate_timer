FROM node:16.14-alpine

WORKDIR /app

COPY ["ultimate_timer/package.json", "ultimate_timer/yarn.lock", "./"]

RUN yarn install

COPY ultimate_timer/ .

CMD [ "yarn", "dev" ]
# CMD [ "yarn", "build", "&&", "yarn", "start"]