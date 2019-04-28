#!/bin/bash

mkdir -p vendor/golang.org/x 

git clone https://github.com/kr/pty vendor/github.com/kr/pty
git clone https://github.com/gin-gonic/gin vendor/github.com/gin-gonic/gin
git clone https://github.com/golang/protobuf vendor/github.com/golang/protobuf
git clone https://github.com/mattn/go-isatty vendor/github.com/mattn/go-isatty
git clone https://github.com/ugorji/go vendor/github.com/ugorji/go
git clone https://gopkg.in/go-playground/validator.v8 vendor/gopkg.in/go-playground/validator.v8
git clone https://gopkg.in/yaml.v2 vendor/gopkg.in/yaml.v2
git clone https://github.com/gin-contrib/sse vendor/github.com/gin-contrib/sse

cd vendor/golang.org/x
git clone https://github.com/golang/crypto.git
git clone https://github.com/golang/net.git 
git clone https://github.com/golang/sys.git 
git clone https://github.com/golang/tools.git

