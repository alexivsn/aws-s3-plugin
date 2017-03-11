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
	if val, ok := commands["Operation"]; ok {
		if val == "lb" {
			log.Println("Listing all buckets. Inside main go-routine.")
			s3.ListAllBuckets()
		} else {
			log.Fatal("Incorrect operation name")
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
	return commands
}
