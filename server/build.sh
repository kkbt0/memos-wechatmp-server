#!/bin/bash
docker buildx build -t kkbt/memos-wechatmp-server:latest --platform=linux/arm,linux/arm64,linux/amd64 . --push