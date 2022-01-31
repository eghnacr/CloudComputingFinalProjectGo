package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	_ "github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"net/http"
	"regexp"
	"time"
)

const (
	connectionString = "Endpoint=sb://final-service-bus.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=oZ7zntJjjG3LjJxyLBBWngKEnqZyjh31gq05FrN5nwo="
	processQueue     = "process-queue"
	resultQueue      = "result-queue"
)

func main() {
	fmt.Println("Starting service bus receiver")
	client, _ := azservicebus.NewClientFromConnectionString(connectionString, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	receiver, _ := client.NewReceiverForQueue(processQueue, nil)
	for {
		messages, _ := receiver.ReceiveMessages(ctx, 1, nil)

		for _, message := range messages {
			err := receiver.CompleteMessage(context.Background(), message)
			if err != nil {
				fmt.Println(err)
			}
			msg, _ := message.Body()
			operationLocation := string(msg)
			//wait for the operation to complete
			time.Sleep(2 * time.Second)
			result := getResultFromOperationLocation(operationLocation)
			tckn := extractTckn(result)
			writeResultToQueue(tckn)
			fmt.Println("Received message: ", string(msg))
		}
	}
}

func writeResultToQueue(result string) {
	client, _ := azservicebus.NewClientFromConnectionString(connectionString, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	sender, _ := client.NewSender(resultQueue, nil)
	sender.SendMessage(ctx, &azservicebus.Message{
		Body: []byte(result),
	})
}

func getResultFromOperationLocation(url string) Result {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Ocp-Apim-Subscription-Key", "6485104fc36d4f5a81905d7b95a7548a")
	req.Header.Set("Content-Type", "application/json")

	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	var result Result
	json.NewDecoder(resp.Body).Decode(&result)
	return result
}

func extractTckn(result Result) string {

	page := result.AnalyzeResult.ReadResults[0]
	r, _ := regexp.Compile("[0-9]{11}")
	for _, sentence := range page.Lines {
		if r.MatchString(sentence.Text) {
			return sentence.Text
		}
	}
	return "can not find tckn"
}
