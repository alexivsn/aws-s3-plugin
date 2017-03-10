package test

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func test() {

	// Initialize a session
	sess := session.Must(session.NewSessionWithOptions(session.Options{SharedConfigState: session.SharedConfigEnable}))

	// create s3 client
	svc := s3.New(sess)

	result, err := svc.ListBuckets(nil)
	if err != nil {
		log.Fatal("Unable to list buckets", err)
	}
	fmt.Println("Buckets:")
	for _, b := range result.Buckets {
		fmt.Printf("* %s created in %s \n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}

	//Create Bucket
	// bucket := "andybucket008"
	// params := &s3.CreateBucketInput{
	// 	Bucket: aws.String(bucket), // Required
	// }
	// resp, err := svc.CreateBucket(params)

	// if err != nil {
	// 	// Print the error, cast err to awserr.Error to get the Code and
	// 	// Message from an error.
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// // Pretty-print the response data.
	// fmt.Println(resp)

	// delete Bucket
	// params1 := &s3.DeleteBucketInput{
	// 	Bucket: aws.String("andybucket007"), // Required
	// }
	// resp1, err := svc.DeleteBucket(params1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp1)

	// Add objects into bucket
	// fileName := "/home/aniyer/notes/aws_dev_setup"
	// file, err := os.Open(fileName)
	// if err != nil {
	// 	log.Fatal("Unable to open file!", err)
	// }
	// defer file.Close()

	// uploader := s3manager.NewUploader(sess)

	// _, err = uploader.Upload(&s3manager.UploadInput{Bucket: aws.String(bucket), Key: aws.String(fileName), Body: file})
	// if err != nil {
	// 	log.Fatal("Unable to upload file!", err)
	// }

	// Delete Objects from bucket
	// fileName := "/home/aniyer/notes/aws_dev_setup"
	// params := &s3.DeleteObjectInput{Bucket: aws.String("andybucket008"), Key: aws.String(fileName)}
	// resp, err := svc.DeleteObject(params)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp)

}
