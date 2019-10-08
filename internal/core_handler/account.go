package core_handler

import (
	grpcAccount "github.com/GuyARoss/project-orva/pkg/grpc/account"

	"github.com/GuyARoss/project-orva/pkg/orva"
)

// AccountRoutineHandler handles the profiles routine
func (req *RoutineRequest) AccountRoutineHandler(ctx *orva.SessionContext) {
	user, userErr := req.verifyUserAccount(ctx.InitialInput.UserID)
	device, deviceErr := req.verifyDeviceAccount(ctx.InitialInput.DeviceID)

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

// account_checkUser checks user profiaccountle under the account routine handler.
func (req *RoutineRequest) verifyUserAccount(userID string) (*grpcAccount.Account, error) {
	profileReq := &grpcAccount.AccountRequest{
		ID:   userID,
		Type: int32(orva.UserAccount),
	}

	return req.AccountClient.RetrieveFromId(nil, profileReq)
}

// account_checkDevice checks device account under the account routine handler.
func (req *RoutineRequest) verifyDeviceAccount(deviceID string) (*grpcAccount.Account, error) {
	profileReq := &grpcAccount.AccountRequest{
		ID:   deviceID,
		Type: int32(orva.DeviceAccount),
	}

	return req.AccountClient.RetrieveFromId(nil, profileReq)
}
