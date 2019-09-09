package routines

import (
	grpcSkill "github.com/GuyARoss/project-orva/pkg/grpc/skill"
	"github.com/GuyARoss/project-orva/pkg/orva"
)

// SkillRouineHandler handles the skill routine
func (req *RoutineRequest) SkillRouineHandler(ctx *orva.SessionContext) {
	skillQuery := &grpcSkill.QueryRequest{
		Statement: ctx.InitialInput.Message,
	}

	resp, err := req.SkillClient.QuerySkill(nil, skillQuery)

	if err != nil {
		return
	}

	skillResp, skillErr := orva.SkillWrapper(resp.Endpoint, ctx)

	if skillErr != nil {
		return
	}

	ctx.Append(skillResp)
}
