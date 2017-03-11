// Package awss3 will be the face to the AWS S3 layer.
package awss3

import (
	"log"

	"github.com/aws/aws-sdk-go/service/s3"
)

/*
ListBuckets : Lists all the buckets in S3.
*/
func ListBuckets(svc *s3.S3) (*s3.ListBucketsOutput, error) {
	log.Println("About to hit the S3 layer. In s3_service")
	res, err := svc.ListBuckets(nil)
	if err != nil {
		log.Fatal("Error while listing Buckets.", err)
		return nil, err
	}
	return res, nil
}
