package routines

import (
	grpcProfile "github.com/GuyARoss/project-orva/pkg/grpc/profile"

	"github.com/GuyARoss/project-orva/pkg/orva"
)

// ProfilesHandler handles the profiles routine
func (req *RoutineRequest) ProfilesHandler(ctx *orva.SessionContext, statement string) {
	// TODO: makes these request goroutines
	user, userErr := req.checkUser(ctx.InitialInput.UserID)
	device, deviceErr := req.checkDevice(ctx.InitialInput.DeviceID)

	if userErr != nil {
		ctx.UserAccessLvl = orva.AnonAccess
		ctx.Append("Could not validate your user access", "profile handler")
	} else {
		ctx.UserAccessLvl = orva.AccessType(user.AccessLevel)
	}

	if deviceErr != nil {
		ctx.DeviceAccessLvl = orva.AnonAccess
		ctx.Append("Could not validate your device access", "profile handler")
	} else {
		ctx.DeviceAccessLvl = orva.AccessType(device.AccessLevel)
	}
}

func (req *RoutineRequest) checkUser(userID string) (*grpcProfile.Profile, error) {
	profileReq := &grpcProfile.ProfileRequest{
		ID: userID,
		// todo: add account type here
	}

	return req.ProfileClient.RetrieveFromId(nil, profileReq)
}

func (req *RoutineRequest) checkDevice(deviceID string) (*grpcProfile.Profile, error) {
	profileReq := &grpcProfile.ProfileRequest{
		ID: deviceID,
		// todo: add account type here
	}

	return req.ProfileClient.RetrieveFromId(nil, profileReq)
}
