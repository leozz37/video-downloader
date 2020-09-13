# Base image
FROM node:12.2.0-alpine

# Getting frontend port from env
ARG PORT_WEB
ENV PORT=$PORT_WEB

# Set working directory
COPY . /app
WORKDIR /app

# Add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# Install and cache app dependencies
RUN npm install
RUN npm install @vue/cli@3.7.0 -g

# Exposing frontend port
EXPOSE $PORT

# Start app
CMD ["npm", "run", "serve"]