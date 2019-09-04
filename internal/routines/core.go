package routines

import (
	"github.com/GuyARoss/project-orva/pkg/orva"
)

// CoreHandler unifies orva routines
func (req *RoutineRequest) CoreHandler(ctx *orva.SessionContext) {
	req.ProfileRoutineHandler(ctx)

	req.SkillRouineHandler(ctx)

	req.SpeechRoutineHandler(ctx)
}
