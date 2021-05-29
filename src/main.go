package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var buildVersion string

func usage() {
	fmt.Println(usageText)
	os.Exit(0)
}

var usageText = `sqs-bench [options...] <sqs-url>

Options:
-m string            (Required) SQS message payload"
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
`

type SQSSender struct {
	QueueURL       string
	Message        string
	Connections    int
	NumCalls       int
	MessageGroupId string
	RetryNum       int
	Verbose        bool
}

func randomStr(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

// https://stackoverflow.com/questions/47606761/repeat-code-if-an-error-occured
func retry(attempts int, sleep time.Duration, f func() error) (err error) {
	for i := 0; ; i++ {
		err = f()
		if err == nil {
			return
		}

		if i >= (attempts - 1) {
			break
		}

		time.Sleep(sleep)
		fmt.Printf("retrying after error:%s\n", err)
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

func (c *SQSSender) Run() {
	successCount := uint32(0)
	errorCount := uint32(0)
	startTime := time.Now()

	var wg sync.WaitGroup
	for i := 1; i <= c.Connections; i++ {
		wg.Add(1)
		go c.startWorker(i, &wg, &successCount, &errorCount)
	}
	wg.Wait()

	duration := time.Since(startTime).Seconds()
	duration_ms := time.Since(startTime).Milliseconds()
	average_ms := duration_ms / (int64(successCount) + int64(errorCount))

	fmt.Println("-----------------------")
	fmt.Println("SQS Bench Summary")
	fmt.Println("-----------------------")
	fmt.Printf("Sent messages: %v\n", successCount)
	fmt.Printf("Errors: %v\n", errorCount)
	fmt.Printf("Duration (sec): %v\n", duration)
	fmt.Printf("Average (ms): %v\n", average_ms)
}

func (c *SQSSender) startWorker(id int, wg *sync.WaitGroup, successCount *uint32, errorCount *uint32) {
	defer wg.Done()

	queue := getSqsSession()
	randomString := randomStr(10)
	for i := 1; i <= c.NumCalls; i++ {
		deduplicationId := fmt.Sprintf("%s%d", randomString, i)
		if c.Verbose {
			fmt.Printf("[Verbose] Mssage GroupID %s Deduplication ID %s Body %s\n", c.MessageGroupId, deduplicationId, c.Message)
		}
		err := retry(c.RetryNum, 2*time.Second, func() (err error) {
			_, qserr := queue.SendMessage(
				&sqs.SendMessageInput{
					DelaySeconds:           aws.Int64(0),
					MessageBody:            &c.Message,
					QueueUrl:               &c.QueueURL,
					MessageGroupId:         &c.MessageGroupId,
					MessageDeduplicationId: &deduplicationId,
				})
			return qserr
		})

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			atomic.AddUint32(errorCount, 1)
			continue
		}

		atomic.AddUint32(successCount, 1)
	}
}

func getSqsSession() *sqs.SQS {
	sess := session.Must(
		session.NewSessionWithOptions(
			session.Options{
				SharedConfigState: session.SharedConfigEnable,
			}))

	return sqs.New(sess)
}

func main() {

	var (
		queueURL       string
		message        string
		connections    int
		numCalls       int
		messageGroupId string
		retryNum       int
		version        bool
		verbose        bool
	)

	flag.StringVar(&message, "m", "", "(Required) SQS message payload")
	flag.IntVar(&connections, "c", 1, "Number of parallel simultaneous SQS session")
	flag.IntVar(&numCalls, "n", 1, "Run for exactly this number of calls by each SQS session")
	flag.StringVar(&messageGroupId, "g", "1", "SQS message group ID")
	flag.IntVar(&retryNum, "r", 1, "Number fo Retry in each message send")
	flag.BoolVar(&version, "version", false, "Build version")
	flag.BoolVar(&verbose, "verbose", false, "Verbose option")
	flag.Usage = usage
	flag.Parse()

	if version {
		fmt.Printf("version: %s\n", buildVersion)
		os.Exit(0)
	}

	args := flag.Args()
	//fmt.Printf("Dump args: %v\n", args)
	if len(args) != 1 {
		usage()
	}
	queueURL = args[0]

	s := SQSSender{
		QueueURL:       queueURL,
		Message:        message,
		Connections:    connections,
		NumCalls:       numCalls,
		MessageGroupId: messageGroupId,
		RetryNum:       retryNum,
		Verbose:        verbose,
	}

	s.Run()
}
