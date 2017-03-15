// Package awss3 will be the face to the AWS S3 layer.
package awss3

import (
	"fmt"
	"log"

	"os"

	"strings"

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

// CreateBucket will create bucket in S3
func (svc *AwsS3) CreateBucket(bucketName string) {
	log.Println("Creating bucket ", bucketName)
	if bucketName != "" {
		resp, err := CreateBucket(svc.S3, bucketName)
		if err != nil {
			log.Fatal("Error while creating bucket.", err)
		}
		fmt.Println("Bucket created with location:", aws.StringValue(resp.Location))
	} else {
		log.Fatal("Cannot create bucket without bucket name.")
	}

}

// DeleteBucket will create bucket in S3
func (svc *AwsS3) DeleteBucket(bucketName string) {
	log.Println("Deleting bucket ", bucketName)
	if bucketName != "" {
		err := DeleteBucket(svc.S3, bucketName)
		if err != nil {
			log.Fatal("Error while deleting bucket.", err)
		}
		fmt.Println("Successfully deleted bucket:", bucketName)
	} else {
		log.Fatal("Cannot delete bucket without bucket name.")
	}

}

// AddObject will add objects into the specified Bucket.
func (svc *AwsS3) AddObject(bucketName string, fileName string) {
	log.Println("Adding Object " + fileName + " into bucket " + bucketName)
	svc.uploadObject(bucketName, fileName)
}

// DeleteObject will delete object from the specified Bucket.
func (svc *AwsS3) DeleteObject(bucketName string, fileName string) {
	log.Println("Deleting Object " + fileName + " from bucket " + bucketName)
	if bucketName != "" && fileName != "" {
		err := DeleteObject(svc.S3, bucketName, fileName)
		if err != nil {
			log.Fatal("Error while adding objects into the bucket.")
		}
		fmt.Println("Successfully deleted object from the bucket.")
	}
}

// UploadObject will upload object. It will automatically add it in the bucket.
func (svc *AwsS3) UploadObject(fileName string) {

	res, _ := ListBuckets(svc.S3)
	var bucketName string
	// Later replace with a customized struct.
	for _, b := range res.Buckets {
		bucketName = aws.StringValue(b.Name)
	}
	log.Println("Uploading Object " + fileName + " into bucket " + bucketName)
	svc.uploadObject(bucketName, fileName)
}

// RemoveObject will remove object from the bucket. Here the Object is searched in every bucket.
// It is advisable to use this function only when the bucket information is not available.
func (svc *AwsS3) RemoveObject(fileName string) {

	res, _ := ListBuckets(svc.S3)
	var bucket string
	var bucketName string
	// Later replace with a customized struct.
	for _, b := range res.Buckets {
		bucketName = aws.StringValue(b.Name)
		resp, err := GetObject(svc.S3, bucketName, fileName)
		if err != nil {
			log.Fatal("Unable to fetch the object " + fileName + " from bucket " + bucketName)
		}
		if resp != nil {
			bucket = bucketName
			log.Println("Bucket to be used:", bucketName)
		}

	}
	if bucket == "" {
		log.Fatal("Unable to fetch the object " + fileName)
	}
	log.Println("Deleting Object " + fileName + " from bucket " + bucketName)
	svc.DeleteObject(bucketName, fileName)
}

//######################## private functions ##################################################3

func (svc *AwsS3) uploadObject(bucketName string, fileName string) {
	if bucketName != "" && fileName != "" {
		file, fileErr := os.Open(fileName)
		s := strings.Split(file.Name(), "/")
		if len(s)-1 > 0 {
			fileName = s[len(s)-1]
		} else {
			fileName = s[0]
		}
		if fileErr != nil {
			log.Fatal("Unable to open file!")
		}
		err := PutObject(svc.S3, bucketName, fileName, file)
		if err != nil {
			log.Fatal("Error while adding objects into the bucket.")
		}
		fmt.Println("Successfully added object into the bucket.")
	}
}
