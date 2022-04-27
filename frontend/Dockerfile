FROM node:16.14-slim

WORKDIR /app

COPY ["ultimate_timer/package.json", "ultimate_timer/yarn.lock", "./"]

RUN yarn global add next \
&&  yarn install

COPY ultimate_timer/ .

CMD [ "yarn", "dev" ]
# CMD [ "yarn", "build", "&&", "yarn", "start"]
