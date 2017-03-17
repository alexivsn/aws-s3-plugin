# aws-s3-plugin
Plugin for S3. This allows creation of buckets, objects and also life cycle them. It also facilitates to upload objects into S3 with really less or no knowledge of S3. 

## s3-cli
> Command line tool for using S3 functionalities.

**s3-plugin** is an library written in Go. It was designed from the ground up to be robust enough for cloud applications and at the same time sufficiently lean to use for any application using s3 as storage.


## Installation
### Prerequisites
* The Go programming language 1.7 or later should be [installed](https://golang.org/doc/install).
* Set GOPATH environment variable on your system
* In order to simplify development and building in Go, we are using the **gb** build tool.  It can be downloaded from [here](https://getgb.io).  

### Install aws-s3-plugin
    go get github.com/anandiyergit/aws-s3-plugin/...

### Update aws-s3-plugin
    go get -u github.com/anandiyergit/aws-s3-plugin/...
    
## Getting Started
This simple example demonstrates how to use s3 using s3-cli
The below section showcases some sample usages

#### Creating Bucket
> ./aws-s3-plugin.git -operation cb -attribute tibcobucket

#### Removing Bucket
> ./aws-s3-plugin.git -operation db -attribute tibcobucket

#### Adding Object
> ./aws-s3-plugin.git -operation ao -attribute /home/aniyer/tibcoobject.txt -bucket tibcobucket

#### Deleting Object
> ./aws-s3-plugin.git -operation do -attribute tibcoobject.txt -bucket tibcobucket

#### Uploading Object without Bucket Information
> ./aws-s3-plugin.git -operation uo -attribute /home/aniyer/tibcoobject.txt

#### Removing Object without Bucket Information
> ./aws-s3-plugin.git -operation ro -attribute tibcoobject.txt




