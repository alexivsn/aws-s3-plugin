package main

import (
	"fmt"

	"log"

	"github.com/anandiyergit/aws-s3-plugin.git/awss3"
	"github.com/anandiyergit/aws-s3-plugin.git/cmd"
)

// The main fucntion.
func main() {

	commands := printCommands()
	s3 := awss3.GetS3()
	var operation, attribute, bucket string
	if opr, ok := commands["Operation"]; ok {
		log.Println("The operation to be performed is:", opr)
		operation = opr
	} else {
		log.Fatal("Operation name is must!")
	}
	if attr, ok := commands["Attributes"]; ok {
		log.Println("The attribute to be used is:", attr)
		attribute = attr
	} else {
		log.Println("The attribute was not used!")
	}
	if buck, ok := commands["Bucket"]; ok {
		log.Println("The Bucket to be used is:", buck)
		bucket = buck
	} else {
		log.Println("The attribute was not used!")
	}
	if operation == "lb" {
		s3.ListAllBuckets()
	}
	if operation == "cb" {
		if attribute != "" {
			s3.CreateBucket(attribute)
		}
	}
	if operation == "db" {
		if attribute != "" {
			s3.DeleteBucket(attribute)
		}
	}
	if operation == "ao" {
		if attribute != "" && bucket != "" {
			s3.AddObject(bucket, attribute)
		}
	}
	if operation == "do" {
		if attribute != "" && bucket != "" {
			s3.DeleteObject(bucket, attribute)
		}
	}
	if operation == "uo" {
		if attribute != "" {
			s3.UploadObject(attribute)
		}
	}
	if operation == "ro" {
		if attribute != "" {
			s3.RemoveObject(attribute)
		}
	}
	log.Println("Program Ended!")
}

// This should validate and print commands.
func printCommands() map[string]string {
	commands := make(map[string]string)
	fmt.Println("Operation is:", *cmd.Operation)
	commands["Operation"] = *cmd.Operation
	fmt.Println("Attribute is:", *cmd.Attribute)
	commands["Attributes"] = *cmd.Attribute
	commands["Bucket"] = *cmd.Bucket
	return commands
}
