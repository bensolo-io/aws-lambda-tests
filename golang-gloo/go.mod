module github.com/bdlilley/aws-lambda-tests/golang-gloo

go 1.18

require (
	github.com/aws/aws-lambda-go v1.36.0
	github.com/davecgh/go-spew v1.1.1
	github.com/google/uuid v1.3.0
)

replace github.com/bdlilley/aws-lambda-tests/golang-util => ../golang-util
