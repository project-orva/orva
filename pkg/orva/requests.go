package orva

import (
	"encoding/json"
	"net/http"
)

type skillPayload struct {
	UserID            string
	Message           string
	UserAccessLevel   int32
	DeviceAccessLevel int32
}

// SkillWrapper makes skill request with the ctx
func SkillWrapper(endoint string, ctx *SessionContext) (string, error) {
	payload, err := skill_buildPayload(ctx)
	
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	
    if err != nil {
        return "", error
    }
    defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
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
