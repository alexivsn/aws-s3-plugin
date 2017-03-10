package awss3

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AwsS3 struct {
	*s3.S3
}

func NewSession() *AwsS3 {
	// Initialize a session
	sess := session.Must(session.NewSessionWithOptions(session.Options{SharedConfigState: session.SharedConfigEnable}))
	return &AwsS3{s3.New(sess)}
}

func (svc *AwsS3) ListAllBuckets() {
	log.Print("List all the buckets in S3.")
	_, _ = ListBuckets(nil)
	//fmt.Println(res)
}
