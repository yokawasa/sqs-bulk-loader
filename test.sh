#!/bin/bash

awsaccountid=1234567890
connections=10
numcalls=10
retry=1

./dist/sqs-bench_darwin \
  -c ${connections} \
  -n ${numcalls} \
  -r ${retry} \
  -m "{\"name\": \"Yoichi Kawasaki\"}" \
  -verbose \
  https://sqs.ap-northeast-1.amazonaws.com/${awsaccountid}/my-sqs-test.fifo 
