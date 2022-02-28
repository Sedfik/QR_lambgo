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
	"fmt"

	"github.com/Sedfik/QR_lambgo/core/sqs"
	"github.com/aws/aws-sdk-go/aws/session"
)

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
