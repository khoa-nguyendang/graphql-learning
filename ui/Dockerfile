FROM node:20-alpine AS builder
WORKDIR /app
COPY . .
RUN npm install -g pnpm
RUN apk add --no-cache libc6-compat
RUN apk update
ENV PNPM_HOME=/app/.pnpm
ENV PATH=$PNPM_HOME:$PATH
ARG REACT_APP_HOST_IP_ADDRESS
ENV REACT_APP_HOST_IP_ADDRESS $REACT_APP_HOST_IP_ADDRESS
RUN pnpm install
RUN pnpm run build
RUN pnpm prune --production

FROM nginx:1.27-alpine
WORKDIR /app
RUN apk add --update nodejs npm
COPY --from=builder /app/dist ./build
RUN rm -rf /usr/share/nginx/html/*

ENV NODE_ENV=production
ENTRYPOINT npx react-inject-env set && cp -R /app/build/* /usr/share/nginx/html/ && 'nginx' '-g' 'daemon off;'