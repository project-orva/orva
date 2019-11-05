package core_handler

import (
	grpcSkill "github.com/GuyARoss/project-orva/pkg/grpc/skill"
	"github.com/GuyARoss/project-orva/pkg/orva"
)

// SkillRoutineHandler handles the skill routine
func (req *RoutineRequest) SkillRoutineHandler(ctx *orva.SessionContext) {
	skillQuery := &grpcSkill.QueryRequest{
		Statement: ctx.InitialInput.Message,
	}

	resp, err := req.SkillClient.QuerySkill(nil, skillQuery)

	if err != nil {
		return
	}

	skillResp, skillErr := orva.SkillWrapper(resp.Endpoint, ctx)

	// if an error occurs or if a skill response is not found 
	// then we want the speech routine to kick in and handle the speech. 
	if skillErr != nil {
		return
	}

	ctx.Append(skillResp)
}
