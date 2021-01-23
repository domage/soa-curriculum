FROM node:12.17.0-slim
EXPOSE 3000

WORKDIR /usr/src/app
COPY package*.json package-lock*.json /usr/src/app/
RUN npm install
RUN npm install -g nodemon

CMD ["nodemon", "index.js"]

