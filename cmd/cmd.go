package cmd

import (
	"flag"
	"fmt"
)

var command = flag.String("operation", "", "S3 operation")
var attribute = flag.String("attribute", "", "S3 attribute")

func init() {

	flag.Parse()
	fmt.Println("operation:", *command)
	fmt.Println("attribute:", *attribute)
	//s3_service.ListBuckets()

}
