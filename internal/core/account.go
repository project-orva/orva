package core

import (
	"github.com/GuyARoss/project-orva/pkg/orva"
	pgUser"github.com/GuyARoss/project-orva/pkg/pg-schemas/user"
	pgDevice"github.com/GuyARoss/project-orva/pkg/pg-schemas/device"
)

// AccountRoutineHandler handles the profiles routine
func (req *RoutineRequest) AccountRoutineHandler(ctx *orva.SessionContext) {
	user, userErr := req.verifyUser(ctx.InitialInput.UserID)
	device, deviceErr := req.verifyDevice(ctx.InitialInput.DeviceID)

	if userErr != nil {
		ctx.UserAccessLvl = orva.AnonAccess
		resp := &orva.Response{
			Statement: "Could not validate your user access",
		}
		ctx.Append(resp)
	} else {
		ctx.UserAccessLvl = orva.AccessType(user.AccessLevel)
	}

	if deviceErr != nil {
		ctx.DeviceAccessLvl = orva.AnonAccess

		resp := &orva.Response{
			Statement: "Could not validate your device access",
		}

		ctx.Append(resp)
	} else {
		ctx.DeviceAccessLvl = orva.AccessType(device.AccessLevel)
	}
}

// account_checkUser checks user profile account under the account routine handler.
func (req *RoutineRequest) verifyUser(UID string) (*pgUser.User, error) {
	pgReq := &user.Request{req.PgCreds}
	
	return pgReq.FindById(UID)
}

// account_checkDevice checks device account under the account routine handler.
func (req *RoutineRequest) verifyDevice(DID string) (*grpcAccount.Account, error) {
	pgReq := &device.Request{reg.PgCreds}

	return pgReq.FindById(DID)
}
