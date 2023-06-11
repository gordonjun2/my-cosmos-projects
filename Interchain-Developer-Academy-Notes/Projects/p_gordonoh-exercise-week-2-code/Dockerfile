FROM --platform=linux golang:1.19.5-bullseye

# Change your versions here
ENV IGNITE_VERSION=0.22.1

ENV PACKAGES curl gcc jq make unzip
RUN apt-get update
RUN apt-get install -y $PACKAGES

# Install Ignite
RUN curl -L https://get.ignite.com/cli@v${IGNITE_VERSION}! | bash

WORKDIR /original
ADD . /original
RUN go mod download

WORKDIR /exercise