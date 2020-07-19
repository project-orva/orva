package auth

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"net/http"
)

type DispatchAuthPayload struct {
	ID string `json:"resource_id"`
	Key string `json:"resource_key"`	
}

type DispatchAuthResponse struct {
	IdentityToken string `json:"identity_token"`
	IAT uint64 `json:"iat"`
}

func FetchIdentityToken(
	authURI string,
	payload *DispatchAuthPayload,
) (*DispatchAuthResponse, error) {
	payloadBytes, marshalErr := json.Marshal(payload)
	if marshalErr != nil {
		return nil, marshalErr
	}

	authURL := fmt.Sprintf("%s/dispatch", authURI)
	authRequest, err := http.NewRequest("POST", authURL, bytes.NewBuffer(payloadBytes))
	authRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(authRequest)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ioout, ioErr := ioutil.ReadAll(resp.Body)

	if ioErr != nil {
		return nil, ioErr
	}

	dispatchResp := &DispatchAuthResponse{}
	marErr := json.Unmarshal(ioout, dispatchResp)
	if marErr != nil {
		return nil, marErr
	}

	return dispatchResp, nil
}
