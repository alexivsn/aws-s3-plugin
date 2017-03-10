package awss3

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/service/s3"
)

/*
ListBuckets : Lists all the buckets in S3.
*/
func ListBuckets(svc *s3.S3) (interface{}, error) {
	res, err := svc.ListBuckets(nil)
	if err != nil {
		log.Fatal("Error while listing Buckets.")
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}
