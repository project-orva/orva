package core_handler

import (
	"github.com/GuyARoss/project-orva/pkg/orva"
	"github.com/GuyARoss/project-orva/pkg/pg-schemas/skill"
)

func determineSkill(initialStatement string) (string, error) { 
	return "http://localhost:3000/test", nil
}

// SkillRoutineHandler handles the skill routine
func (req *RoutineRequest) SkillRoutineHandler(ctx *orva.SessionContext) {
	skillEndpoint, err := determineSkill(ctx.InitialInput.Message)
	if err != nil {
		return
	}

	skillResp, skillErr := orva.SkillWrapper(skillEndpoint, ctx)

	// if an error occurs or if a skill response is not found 
	// then we want the speech routine to kick in and handle the speech. 
	if skillErr != nil {
		return
	}

	ctx.Append(skillResp)
}
