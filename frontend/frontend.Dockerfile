FROM node:25.9-alpine3.23

WORKDIR /app

COPY package*.json .
RUN npm i
COPY . .


EXPOSE 4200

CMD ["npm","run","start"]
