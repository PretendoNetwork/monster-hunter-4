package globals

import (
	"context"

	"github.com/PretendoNetwork/nex-go/v2/types"

	"time"

	pb "github.com/PretendoNetwork/grpc-go/account"
	"github.com/PretendoNetwork/nex-go/v2"
	"google.golang.org/grpc/metadata"
)

type UserDataCacheEntry struct {
	userData     *pb.GetUserDataResponse
	creationTime time.Time
}

var UserDataCache map[types.PID]UserDataCacheEntry

func UserDataFromPID(pid types.PID) (*pb.GetUserDataResponse, uint32) {

	if UserDataCache == nil {
		UserDataCache = make(map[types.PID]UserDataCacheEntry)
	}

	data, exists := UserDataCache[pid]
	if !exists || data.creationTime.Add(time.Hour*24).Before(time.Now().UTC()) {
		ctx := metadata.NewOutgoingContext(context.Background(), GRPCAccountCommonMetadata)

		response, err := GRPCAccountClient.GetUserData(ctx, &pb.GetUserDataRequest{Pid: uint32(pid)})
		if err != nil {
			Logger.Error(err.Error())
			UserDataCache[pid] = UserDataCacheEntry{userData: nil, creationTime: time.Now().UTC()}
			return &pb.GetUserDataResponse{}, nex.ResultCodes.RendezVous.InvalidUsername
		}

		UserDataCache[pid] = UserDataCacheEntry{userData: response, creationTime: time.Now().UTC()}
	}

	if data.userData != nil {
		return data.userData, 0
	}

	return nil, nex.ResultCodes.Core.Unknown
}
