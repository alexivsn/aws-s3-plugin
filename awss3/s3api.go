// Package awss3 will be the face to the AWS S3 layer.
package awss3

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// AwsS3 struct
type AwsS3 struct {
	*s3.S3
}

// GetS3 is the Getter for fetching the S3 session.
func GetS3() *AwsS3 {
	sess := session.Must(session.NewSessionWithOptions(session.Options{SharedConfigState: session.SharedConfigEnable}))
	return &AwsS3{s3.New(sess)}
}

// ListAllBuckets will print all the Listing of Buckets.
func (svc *AwsS3) ListAllBuckets() {
	log.Print("List all the buckets in S3. Inside s3api.")
	res, _ := ListBuckets(svc.S3)
	// Later replace with a customized struct.
	for _, b := range res.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}

}
