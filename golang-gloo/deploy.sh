#!/bin/bash

ACCT_ID="${AWS_ACCOUNT_ID:-931713665590}"
REGION="${AWS_REGION:-us-east-2}"
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
NAME="$(basename $SCRIPT_DIR)-latest"

GOOS=linux GOARCH=amd64 go build -o main

zip lambda.zip main

aws s3 cp lambda.zip s3://solo-io-terraform-${REGION}-${ACCT_ID}/lambda/${NAME}.zip

aws lambda create-function \
    --function-name ${NAME} \
    --runtime go1.x \
    --code=S3Bucket=solo-io-terraform-${REGION}-${ACCT_ID},S3Key=lambda/${NAME}.zip \
    --handler main \
    --role arn:aws:iam::${ACCT_ID}:role/lambda-basic \
    --region ${AWS_REGION} \
    || true

aws lambda update-function-code \
    --function-name ${NAME} \
    --s3-bucket solo-io-terraform-${REGION}-${ACCT_ID} \
    --region ${AWS_REGION} \
    --s3-key lambda/${NAME}.zip

# kubectl apply -f - <<EOF
# apiVersion: lambda.aws.crossplane.io/v1beta1
# kind: Function
# metadata:
#   name: ${NAME}
#   namespace: crossplane-system
# spec:
#   forProvider:
#     region: us-east-2
#     description: ${NAME}
#     runtime: go1.x
#     code:
#       s3Bucket: solo-io-terraform-${ACCT_ID}
#       s3Key: lambda/${NAME}.zip
# EOF