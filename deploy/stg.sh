#!/usr/bin/env bash

APP_NAME=helmet_shop

docker load -i ${APP_NAME}.tar
docker rm -f ${APP_NAME}

docker run -d --name ${APP_NAME} \
  --network my-net \
  -e VIRTUAL_HOST="my-domain-name" \
  -e LETSENCRYPT_HOST="my-domain-name" \
  -e LETSENCRYPT_EMAIL="demo@gmail.com" \
  -e DBConectionStr="helmet_shop:1234@tcp(mysql:3306)/helmet_shop?charset=utf8mb4&parseTime=True&loc=Local" \
  -e S3Secretkey="asbBuwc34W0n5mXhbM9x2WU019r8/LUbJ5Q/o5Qj" \
  -e S3ApiKey="AKIA42E3U7VD3UISMCM4" \
  -e S3Region="ap-southeast-1" \
  -e S3BucketName="bucket"\
  -e S3Domain="https://d3pfouzi5at9lt.cloudfront.net" \
  -e SYSTEM_SECRET="demo" \
  -p 8080:8080 \
  ${APP_NAME}