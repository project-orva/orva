package orva

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type skillPayload struct {
	UserID            string
	Message           string
	UserAccessLevel   int32
	DeviceAccessLevel int32
}

// SkillProxy makes skill request with the ctx
func SkillProxy(endoint string, ctx *SessionContext) (*Response, error) {
	payload, err := skill_buildPayload(ctx)

	req, err := http.NewRequest("POST", endoint, bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ioout, ioErr := ioutil.ReadAll(resp.Body)

	if ioErr != nil {
		return nil, ioErr
	}

	respDo := &Response{}
	marErr := json.Unmarshal(ioout, respDo)

	return respDo, marErr
}

// skill_buildPayload builds the skill payload
func skill_buildPayload(ctx *SessionContext) (string, error) {
	payload := &skillPayload{
		UserID:            ctx.InitialInput.UserID,
		Message:           ctx.InitialInput.Message,
		UserAccessLevel:   int32(ctx.UserAccessLvl),
		DeviceAccessLevel: int32(ctx.DeviceAccessLvl),
	}
	b, err := json.Marshal(payload)

	return string(b), err
}
