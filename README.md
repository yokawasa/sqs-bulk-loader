# sqs-bulk-loader

[![Upload Release Asset](https://github.com/yokawasa/sqs-bulk-loader/actions/workflows/release.yml/badge.svg)](https://github.com/yokawasa/sqs-bulk-loader/actions/workflows/release.yml)

A Golang tool that sends bulk messages in parallel to Amazon SQS

## Usage

```
sqs-bulk-loader [options...] <sqs-url>

Options:
-m string            (Required) Amazon SQS message payload to send"
-c connections       Number of parallel simultaneous SQS session
                     By default 1; Must be more than 0
-n num-calls         Run for exactly this number of calls by each SQS session
                     By default 1; Must be more than 0
-g group-id          SQS message group ID
                     By default "1"
-r retry-num         Number fo Retry in each message send
                     By default 1; Must be more than 0
-version             Prints out build version information
-verbose             Verbose option
-h                   help message
```

## Build and Run

To build, simply run `make` like below
```
make

golint $HOME/dev/github/sqs-bulk-loader
GOOS=linux GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=v0.0.1" -o $HOME/dev/github/sqs-bulk-loader/dist/sqs-bulk-loader_linux $HOME/dev/github/sqs-bulk-loader/src
GOOS=darwin GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=v0.0.1" -o $HOME/dev/github/sqs-bulk-loader/dist/sqs-bulk-loader_darwin $HOME/dev/github/sqs-bulk-loader/src
```

Suppose you are using macOS, run the `sqs-bulk-loader_darwin` (while `sqs-bulk-loader_linux` if you are using Linux) like below

```bash
connections=10
numcalls=10
retry=1

./dist/sqs-bulk-loader_darwin \
  -c ${connections} \
  -n ${numcalls} \
  -r ${retry} \
  -m "{\"name\": \"Yoichi Kawasaki\"}" \
  https://sqs.ap-northeast-1.amazonaws.com/1234567890/my-sqs-test.fifo

-----------------------
SQS Bench Summary
-----------------------
Sent messages: 100
Errors: 0
Duration (sec): 0.674113624
Average (ms): 6
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/yokawasa/sqs-bulk-loader.
