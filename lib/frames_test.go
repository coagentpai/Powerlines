package main
import "testing"

func TestCommandAliases(t *testing.T) {
	for _, frameId := range AllFrameIds {
		var inRequests  = false
		var inResponses = false
		if _, ok := FrameRequestAliases[frameId]; ok {
			inRequests = true
		}
		if _, ok := FrameResponseAliases[frameId]; ok {
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
