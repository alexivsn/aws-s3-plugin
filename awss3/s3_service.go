// Package awss3 will be the face to the AWS S3 layer.
package awss3

import (
	"log"

	"os"

	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

/*
ListBuckets : Lists all the buckets in S3.
*/
func ListBuckets(svc *s3.S3) (*s3.ListBucketsOutput, error) {
	res, err := svc.ListBuckets(nil)
	if err != nil {
		log.Fatal("Error while listing Buckets.", err)
		return nil, err
	}
	return res, nil
}

/*
CreateBucket : Creates a Bucket in the S3 environment. Currently it does not support deep down permission support.
*/
func CreateBucket(svc *s3.S3, bucketName string) (*s3.CreateBucketOutput, error) {
	if bucketName == "" {
		log.Fatal("Bucket name cannot be empty. It is a required field for creating Bucket Input.")
		return nil, nil
	}
	bucketInput := getBucketInput(bucketName)
	resp, err := svc.CreateBucket(&bucketInput)
	if err != nil {
		log.Fatal("Could not create bucket ", bucketName)
		return nil, err
	}
	// Wait until bucket is created before finishing
	log.Println("Waiting for bucket " + bucketName + " to be created.")
	err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Fatal("Error occurred while waiting for bucket " + bucketName + " to be created.")
	}
	return resp, nil
}

/*
DeleteBucket : Deletes a Bucket in the S3 environment. Currently it does not support deep down permission sexport PYTHONPATH=$PYTHONPATH:/usr/lib/python2.7upport.
*/
func DeleteBucket(svc *s3.S3, bucketName string) error {
	if bucketName == "" {
		log.Fatal("Bucket name cannot be empty. It is a required field for creating Bucket Input.")
		return nil
	}
	bucketInput := getBucketInputForDelete(bucketName)
	_, err := svc.DeleteBucket(&bucketInput)
	if err != nil {
		log.Fatal("Could not delete bucket "+bucketName+" Error is:", err)
		return err
	}
	return nil
}

/*
PutObject : Adds Objects into the bucket.
*/
func PutObject(svc *s3.S3, bucketName string, fileName string, file *os.File) error {
	if bucketName == "" {
		log.Fatal("Bucket name cannot be empty. It is a required field for adding Objects into Bucket.")
		return nil
	}
	objectInput := getObjectInput(bucketName, fileName, file)
	_, err := svc.PutObject(&objectInput)
	if err != nil {
		log.Fatal("Could not add object " + fileName + " into bucket " + bucketName)
		return err
	}
	return nil
}

/*
DeleteObject : Deletes Objects from the bucket.
*/
func DeleteObject(svc *s3.S3, bucketName string, fileName string) error {
	if bucketName == "" {
		log.Fatal("Bucket name cannot be empty. It is a required field for creating Bucket Input.")
		return nil
	}
	objectInput := getObjectInputForDelete(bucketName, fileName)
	_, err := svc.DeleteObject(&objectInput)
	if err != nil {
		log.Fatal("Could not delete object " + fileName + " from bucket " + bucketName)
		return err
	}
	return nil
}

/*
GetObject : Fetches Object from the bucket.
*/
func GetObject(svc *s3.S3, bucketName string, fileName string) (*s3.GetObjectOutput, error) {
	if bucketName == "" {
		log.Fatal("Bucket name cannot be empty. It is a required field for fetching Object.")
		return nil, nil
	}
	objectInput := getObjectInputForGet(bucketName, fileName)
	resp, err := svc.GetObject(&objectInput)
	if err != nil {
		if strings.Contains(err.Error(), "NoSuchKey") {
			return nil, nil
		}
		log.Fatal("Could not delete object "+fileName+" from bucket "+bucketName+"Error:", err)
		return resp, err
	}
	return resp, nil
}

//######################## Private functions ###################################
/*
Will prepare BucketInput for creating a bucket.
*/
func getBucketInput(bucketName string) s3.CreateBucketInput {
	if bucketName == "" {
		log.Fatal("Bucket name cannot be empty. It is a required field for creating Bucket Input.")
		return s3.CreateBucketInput{}
	}
	createBucketInput := s3.CreateBucketInput{}
	createBucketInput.Bucket = aws.String(bucketName)
	// <TODO> Add more permission related properties here.
	return createBucketInput
}

/*
Will prepare BucketInput for deleting a bucket.
*/
func getBucketInputForDelete(bucketName string) s3.DeleteBucketInput {
	if bucketName == "" {
		log.Fatal("Bucket name cannot be empty. It is a required field for creating Bucket Input.")
		return s3.DeleteBucketInput{}
	}
	deleteBucketInput := s3.DeleteBucketInput{}
	deleteBucketInput.Bucket = aws.String(bucketName)
	// <TODO> Add more permission related properties here.
	return deleteBucketInput
}

/*
Will prepare ObjectInput for creating/uploading an Object.
*/
func getObjectInput(bucketName string, fileName string, file *os.File) s3.PutObjectInput {
	if bucketName == "" {
		log.Fatal("Bucket name cannot be empty. It is a required field for creating Bucket Input.")
		return s3.PutObjectInput{}
	}
	putObjectInput := s3.PutObjectInput{}
	putObjectInput.Bucket = aws.String(bucketName)
	putObjectInput.Key = aws.String(fileName)
	if file != nil {
		putObjectInput.Body = file
	}
	// <TODO> Add more permission related properties here.
	return putObjectInput
}

/*
Will prepare ObjectInput for deleting an Object.
*/
func getObjectInputForDelete(bucketName string, fileName string) s3.DeleteObjectInput {
	if bucketName == "" {
		log.Fatal("Bucket name cannot be empty. It is a required field for deleting Object Input.")
		return s3.DeleteObjectInput{}
	}
	deleteObjectInput := s3.DeleteObjectInput{}
	deleteObjectInput.Bucket = aws.String(bucketName)
	deleteObjectInput.Key = aws.String(fileName)
	// <TODO> Add more permission related properties here.
	return deleteObjectInput
}

/*
Will prepare ObjectInput for getting an Object.
*/
func getObjectInputForGet(bucketName string, fileName string) s3.GetObjectInput {
	if bucketName == "" {
		log.Fatal("Bucket name cannot be empty. It is a required field for fetching Object Input.")
		return s3.GetObjectInput{}
	}
	getObjectInput := s3.GetObjectInput{}
	getObjectInput.Bucket = aws.String(bucketName)
	getObjectInput.Key = aws.String(fileName)
	// <TODO> Add more permission related properties here.
	return getObjectInput
}
