package orva

// AppliedMessage base message that can be after routines cycle
type AppliedMessage struct {
	Statement    string
	AssignedFrom string
}

// Input user input for the session
type Input struct {
	Message  string
	DeviceID string
	UserID   string
}

// SessionContext ~Core Route Session Context
type SessionContext struct {
	AppliedMessages *[]AppliedMessage
	InitialInput    *Input
	DeviceAccessLvl AccessType
	UserAccessLvl   AccessType
}

// Append session dialog to context
func (ctx *SessionContext) Append(message string, assignedFrom string) {
	appLen := len(*ctx.AppliedMessages) + 1
	newMessages := make([]AppliedMessage, appLen)

	derefMsgs := *ctx.AppliedMessages
	for i := 0; i < appLen; i++ {
		newMessages[i] = derefMsgs[i]
	}

	newMessages[appLen] = AppliedMessage{
		Statement:    message,
		AssignedFrom: assignedFrom,
	}

	ctx.AppliedMessages = &newMessages
}

func (ctx *SessionContext) JsonMarshal() string {

}
