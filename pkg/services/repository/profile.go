package profile

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"	
)

type Profile struct {
	FirstName string `json:"firstName"`
	lastName string `json:"lastName"`
	AccessLevel uint32 `json:"accessLevel"`
	ID string `json:"id"`
}

func FetchUserProfile(
	repositoryURI string,
	userID string,
	identityToken string,
) (*Profile, error) {
	uri := fmt.Sprintf("%s/profile?id=%s", repositoryURI, userID)	
	httpReq, err := http.NewRequest("GET", uri, nil)
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-identity", identityToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ioout, ioErr := ioutil.ReadAll(resp.Body)

	if ioErr != nil {
		return nil, ioErr
	}

	dispatchResp := &Profile{}
	marErr := json.Unmarshal(ioout, dispatchResp)
	if marErr != nil {
		return nil, marErr
	}

	return dispatchResp, nil
}