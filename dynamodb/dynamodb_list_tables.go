package main

import (
    "log"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
svc := dynamodb.New(session.New(&aws.Config{Region: aws.String("us-west-1")}))
result, err := svc.ListTables(&dynamodb.ListTablesInput{})
if err != nil {
    log.Println(err)
    return
}

log.Println("Tables:")
for _, table := range result.TableNames {
    log.Println(*table)
}
}
