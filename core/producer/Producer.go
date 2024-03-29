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
// snippet-start:[sqs.go.list_queues]
package producer

// snippet-start:[sqs.go.list_queues.imports]
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
		"ctx": lc,
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

func Test() {
	// Create a session that gets credential values from ~/.aws/credentials
	// and the default region from ~/.aws/config
	// snippet-start:[sqs.go.list_queues.sess]
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	// snippet-end:[sqs.go.list_queues.sess]
	queueName := string("qr-request")
	queueUrl, _ := sqs.GetQueueURL(sess, &queueName)
	// snippet-start:[sqs.go.list_queues.display]
	fmt.Println(queueUrl)
	// sqs.SendMsg(sess, queueUrl.QueueUrl, "hello, world")
	// snippet-end:[sqs.go.list_queues.display]
	err := sqs.SendMsg(sess, queueUrl.QueueUrl, "hello, world")
	if err != nil {
		fmt.Println("Got an error sending the message:")
		fmt.Println(err)
		return
	}

	fmt.Println("Sent message to queue ")
}

// snippet-end:[sqs.go.list_queues]
