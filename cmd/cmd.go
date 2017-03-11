// Package cmd would be responsible for parsing all the commands.
package cmd

import (
	"flag"
)

// Operation : Relates to S3 operations.
var Operation = flag.String("operation", "", "S3 operation")

// Attribute : attribute name related to S3.
var Attribute = flag.String("attribute", "", "S3 attribute")

// This just initializes the commands.
func init() {

	flag.Parse()

}
