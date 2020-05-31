package orva

import "fmt"

// Response base message that can be after routines cycle
type Response struct {
	Statement    string
	AssignedFrom string
	GraphicURL   string
	GraphicType  int32
	Error        string
	Duration     float32
}

// Input user input for the session
type Input struct {
	Message  string
	DeviceID string
	UserID   string
}

// SessionContext ~Core Route Session Context
type SessionContext struct {
	AppliedMessages []Response
	InitialInput    Input
	DeviceAccessLvl AccessType
	UserAccessLvl   AccessType
}

// CreateContext creates session context from input
func CreateContext(input *Input) *SessionContext {
	emptyApplied := make([]Response, 0)

	ctx := &SessionContext{
		AppliedMessages: emptyApplied,
	}
	ctx.InitialInput = *input

	return ctx
}

// Append session dialog to context
func (ctx *SessionContext) Append(resp *Response) {
	fmt.Println(resp)
	fmt.Println(ctx)
	
	ctx.AppliedMessages = append(ctx.AppliedMessages, *resp)
}

// Responses extracts and returns the response from the session context.
func (ctx *SessionContext) Responses() []Response {
	return ctx.AppliedMessages
}
