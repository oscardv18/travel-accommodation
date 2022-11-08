FROM node:18.12.0-alpine3.16

RUN mkdir /home/travel-accommodation

WORKDIR /home/travel-accommodation

EXPOSE 4000

CMD npx vite --port=4000