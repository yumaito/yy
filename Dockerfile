FROM node:13

WORKDIR /app
ENV YARN_VERSION 1.22.4
ENV HOST 0.0.0.0
ENV PORT 3000

RUN apt-get update && \
    apt-get install -y curl && \
    npm install -g npm && \
    curl -o- -L https://yarnpkg.com/install.sh | bash -s -- --version $YARN_VERSION

EXPOSE 3000
