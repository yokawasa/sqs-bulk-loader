# sqs-bulk-loader

[![Upload Release Asset](https://github.com/yokawasa/sqs-bulk-loader/actions/workflows/release.yml/badge.svg)](https://github.com/yokawasa/sqs-bulk-loader/actions/workflows/release.yml)

A Golang tool that sends bulk messages in parallel to Amazon SQS

## Usage

```
sqs-bench [options...] <sqs-url>

Options:
-m string            (Required) Amazon SQS message payload to send "
-c connections       Number of parallel simultaneous SQS session
                     By default 1; Must be more than 0
-n num-calls         Run for exactly this number of calls by each SQS session
                     By default 1; Must be more than 0
-g group-id          SQS message group ID
                     By default "1"
-r retry-num         Number fo Retry in each message send
                     By default 1; Must be more than 0
-endpoint-url string The URL to send the API request to
                     By default "", which mean the AWS SDK automatically determines the URL
-version             Prints out build version information
-verbose             Verbose option
-h                   help message
```

## Download

You can download the compiled command with [downloader](https://github.com/yokawasa/sqs-bulk-loader/blob/main/downloader) like this:

```
# Download latest command
./downloader

# Download the command with a specified version
./downloader v0.0.2
```
Or you can download it on the fly with the following commmand:

```
curl -sS https://raw.githubusercontent.com/yokawasa/sqs-bulk-loader/main/downloader | bash --
```


Output would be like this:
```
Archive:  sqs-bulk-loader.zip
  inflating: sqs-bulk-loader_darwin_amd64
sqs-bulk-loader
Downloaded into sqs-bulk-loader
Please add sqs-bulk-loader to your path; e.g copy paste in your shell and/or ~/.profile
```

## Execute the command

```bash
connections=10
numcalls=10
retry=1

sqs-bulk-loader \
  -c ${connections} \
  -n ${numcalls} \
  -r ${retry} \
  -m "{\"name\": \"Yoichi Kawasaki\"}" \
  https://sqs.ap-northeast-1.amazonaws.com/1234567890/my-queue.fifo
```

Sample output would be like this:

```
-----------------------
SQS Bench Summary
-----------------------
Sent messages: 100
Errors: 0
Duration (sec): 0.674113624
Average (ms): 6
```

## Build and Run (For Developer)

To build, simply run `make` like below
```
make

golint $HOME/dev/github/sqs-bulk-loader
GOOS=linux GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=v0.0.1" -o $HOME/dev/github/sqs-bulk-loader/dist/sqs-bulk-loader_linux $HOME/dev/github/sqs-bulk-loader/src
GOOS=darwin GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=v0.0.1" -o $HOME/dev/github/sqs-bulk-loader/dist/sqs-bulk-loader_darwin $HOME/dev/github/sqs-bulk-loader/src
GOOS=windows GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=v0.0.1" -o $HOME/dev/github/sqs-bulk-loader/dist/sqs-bulk-loader_windows $HOME/dev/github/sqs-bulk-loader/src
```

Suppose you are using macOS, run the `sqs-bulk-loader_darwin_amd64` (while `sqs-bulk-loader_linux_amd64` if you are using Linux, `sqs-bulk-loader_windows_amd64` if using Windows) like below

```bash
./dist/sqs-bulk-loader_darwin -m "test" -c 10 -n 10 https://sqs.ap-northeast-1.amazonaws.com/1234567890/my-queue.fifo
```

Finally clean built commands

```
make clean
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/yokawasa/sqs-bulk-loader.
