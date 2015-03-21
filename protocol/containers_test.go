package protocol

import "testing"

func TestCommandAliases(t *testing.T) {
	for _, frameId := range AllContainerIds {
		var inRequests  = false
		var inResponses = false
		if _, ok := ContainerRequestAliases[frameId]; ok {
			inRequests = true
		}
		if _, ok := ContainerResponseAliases[frameId]; ok {
			inResponses = true
		}
		if !inResponses && !inRequests {
			t.Errorf("Frame not defined as a request or response: %d", frameId)
		}
		if inRequests && inResponses {
			t.Errorf("Frame defined as both a request and response: %d", frameId)
		}
	}
}
