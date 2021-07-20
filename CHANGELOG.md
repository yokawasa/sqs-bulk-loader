# Change Log

All notable changes to the "sqs-bulk-loader" will be documented in this file.

## v0.0.3

- fix go module path to github.com/yokawasa/sqs-bulk-loader

## v0.0.2

- Add `-endpoint-url` to specify the URL to send the API request to. For most cases, the AWS SDK automatically determines the URL based on the selected service and the specified AWS Region

**NOTE**
From v0.0.2, you can specify the url to send API requests to. For example, you can give the localstack endpoint (http://localhost:4566) when you develop SQS application using localstack like this:

```
sqs-bulk-loader -m "test" -c 10 -n 10 -endpoint-url http://localhost:4566  http://localhost:4566/000000000000/my-queue.fifo
``` 

## v0.0.1

- Initial release
