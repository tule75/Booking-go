#!/bin/bash

# docker build
make docker_build

echo "127.0.0.1:8081" >> /etc/hosts

