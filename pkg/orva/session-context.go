package orva

// Response base message that can be after routines cycle
type Response struct {
	Statement    string
	AssignedFrom string
	GraphicURL   string
	GraphicType  int32
	Error        string
}

// Input user input for the session
type Input struct {
	Message  string
	DeviceID string
	UserID   string
}

// SessionContext ~Core Route Session Context
type SessionContext struct {
	AppliedMessages *[]Response
	InitialInput    *Input
	DeviceAccessLvl AccessType
	UserAccessLvl   AccessType
}

// CreateContext creates session context from input
func CreateContext(input *Input) *SessionContext {
	ctx := &SessionContext{}
	ctx.InitialInput = input

	return ctx
}

// Append session dialog to context
func (ctx *SessionContext) Append(resp *Response) {
	appLen := len(*ctx.AppliedMessages) + 1
	newMessages := make([]Response, appLen)

	derefMsgs := *ctx.AppliedMessages
	for i := 0; i < appLen; i++ {
		newMessages[i] = derefMsgs[i]
	}

	newMessages[appLen] = *resp

	ctx.AppliedMessages = &newMessages
}

// Responses extracts and returns the response from the session context.
func (ctx *SessionContext) Responses() *[]Response {
	return ctx.AppliedMessages
}