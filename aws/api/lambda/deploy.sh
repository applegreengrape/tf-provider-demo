#!/bin/sh

zip tag.zip tag.py

aws lambda update-function-code \
    --function-name  tag \
    --zip-file fileb://tag.zip