FROM node:16.14-slim

WORKDIR /app

RUN rm -rf node_modules/ 

COPY ./ultimate_timer/ .

RUN yarn global add next@12.1.0  \
&&  yarn upgrade \
&&  yarn install

# CMD [ "yarn", "dev" ]
# CMD [ "yarn", "build", "&&", "yarn", "start"]
