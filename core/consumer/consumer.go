/*
   Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.

   This file is licensed under the Apache License, Version 2.0 (the "License").
   You may not use this file except in compliance with the License. A copy of
   the License is located at

    http://aws.amazon.com/apache2.0/

   This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
   CONDITIONS OF ANY KIND, either express or implied. See the License for the
   specific language governing permissions and limitations under the License.
*/
// snippet-start:[sqs.go.receive_messages]
package consumer

// snippet-start:[sqs.go.receive_messages.imports]
import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Sedfik/QR_lambgo/core/sqs"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context) (Response, error) {

	lc, _ := lambdacontext.FromContext(ctx)
	log.Print(lc.Identity.CognitoIdentityPoolID)

	var buf bytes.Buffer
	body, err := json.Marshal(map[string]interface{}{
		"message": "Go Serverless v1.0! Your function executed successfully!",
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}

func test() {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-3"),
		Credentials: credentials.NewSharedCredentials("", "perso"),
	})

	queue := string("qr-request")
	// Get URL of queue
	urlResult, err := sqs.GetQueueURL(sess, &queue)
	if err != nil {
		fmt.Println("Got an error getting the queue URL:")
		fmt.Println(err)
		return
	}

	// snippet-start:[sqs.go.receive_message.url]
	queueURL := urlResult.QueueUrl
	// snippet-end:[sqs.go.receive_message.url]

	timeout := int64(100)
	msgResult, err := sqs.GetMessages(sess, queueURL, &timeout)
	if err != nil {
		fmt.Println("Got an error receiving messages:")
		fmt.Println(err)
		return
	}

	fmt.Println("Message ID:     " + *msgResult.Messages[0].MessageId)

	fmt.Println("Message Handle: " + *msgResult.Messages[0].ReceiptHandle)
	fmt.Println("Message Body: " + *msgResult.Messages[0].Body)

}
