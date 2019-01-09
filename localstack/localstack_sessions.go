package localstack

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var testCredentials = credentials.NewStaticCredentials("AKID", "SECRET", "SESSION")
var testRegion = aws.String("eu-west-1")
var disableSSL = aws.Bool(true)

// S3Session provides an AWS session suitable for use with S3 from LocalStack
func S3Session() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Credentials: testCredentials,
		Region:      testRegion,
		Endpoint:    aws.String("http://localhost:4572"),
		DisableSSL:  disableSSL,
	})
	checkError(err)
	return sess
}

// SNSSession provides an AWS session suitable for use with SNS from LocalStack
func SNSSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Credentials: testCredentials,
		Region:      testRegion,
		Endpoint:    aws.String("http://localhost:4575"),
		DisableSSL:  disableSSL,
	})
	checkError(err)
	return sess
}

// SQSSession provides an AWS session suitable for use with SQS from LocalStack
func SQSSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Credentials: testCredentials,
		Region:      testRegion,
		Endpoint:    aws.String("http://localhost:4576"),
		DisableSSL:  disableSSL,
	})
	checkError(err)
	return sess
}

func checkError(err error) {
	if err != nil {
		panic(fmt.Sprintf("Unable to create AWS session, %v", err))
	}
}