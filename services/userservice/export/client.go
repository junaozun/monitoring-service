package export

import (
	"log"

	"github.com/junaozun/monitoring-service/services/userservice/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var client *grpc.ClientConn

func init() {
	var err error
	client, err = grpc.Dial("userservice:8083", grpc.WithInsecure())
	if err != nil {
		log.Fatal("init post client fail ", err)
	}
}

func GetUser(ctx context.Context, uid int) (*pb.GetUserResp, error) {
	return pb.NewUserServiceClient(client).GetUser(ctx, &pb.GetUserReq{
		Id: int64(uid),
	})
}
