package impl

import (
	"log"

	"github.com/junaozun/monitoring-service/services/postservice/pb"
	"github.com/junaozun/monitoring-service/services/userservice/export"
	context "golang.org/x/net/context"
)

type Service struct {
}

func (s Service) GetPost(context context.Context, req *pb.GetPostReq) (*pb.GetPostResp, error) {
	userResp, err := export.GetUser(context, 13)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.GetPostResp{
		Uid:  userResp.Uid,
		Text: "get post success the user is " + userResp.Name,
	}, nil
}
