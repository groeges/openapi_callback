/*
 * test
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// DefaultApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() DefaultApiServicer {
	return &DefaultApiService{}
}

// Callback - Subscribe to a webhook
func (s *DefaultApiService) Callback(ctx context.Context, inlineObject InlineObject) (ImplResponse, error) {
	// TODO - update CallbackPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	done := make(chan bool)

	go func() {
		sendEvent(done, inlineObject.EventId, inlineObject.CallbackUrl)
	}()

	//TODO: Uncomment the next line to return response Response(201, {}) or use other options such as http.Ok ...
	return Response(201, inlineObject), nil

	//return Response(http.StatusNotImplemented, nil), errors.New("CallbackPost method not implemented")
}

func sendEvent(done chan<- bool, eventId string, callbackURL string) (okay bool) {
	defer func() {
		done <- okay
	}()

	fmt.Println(fmt.Sprintf("%s: Sending an event after 20 seconds", eventId))

	time.Sleep(20 * time.Second)

	values := map[string]string{"message": fmt.Sprintf("Remote process has completed for eventId: %s", eventId)}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", callbackURL, bytes.NewBuffer(json_data))

	req.Header.Add("ce-specversion", "1.0")
	req.Header.Add("ce-source", "pm")
	req.Header.Add("ce-type", "move")
	req.Header.Add("ce-id", "random")
	req.Header.Add("ce-kogitoprocrefid", eventId)
	req.Header.Add("ce-workflowdata", "{}")
	resp, err := http.DefaultClient.Do(req)

	fmt.Println(fmt.Sprintf("Response from request is: %d", resp.StatusCode))
	if err != nil {
		fmt.Println(fmt.Sprintf("%s: Successfully sent event", eventId))
		return false
	} else {
		fmt.Println(fmt.Sprintf("%s: Failed to send event: %v", eventId, err))
		return true
	}
}
