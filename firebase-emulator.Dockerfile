FROM node:alpine
RUN apk add openjdk11
RUN npm install -g firebase-tools
WORKDIR /
CMD [ "firebase", "emulators:start", "--only", "storage"]
