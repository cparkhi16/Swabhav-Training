FROM node:13-alpine

RUN mkdir -p /home/app
COPY ./nodejs-mysql-starter-kit2 /home/app
WORKDIR /home/app

RUN npm install
CMD ["npm","start"]