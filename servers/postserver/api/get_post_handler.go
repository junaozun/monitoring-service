package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/junaozun/monitoring-service/services/postservice/export"
)

func GetPostHandler(resp http.ResponseWriter, req *http.Request) {
	postResp, err := export.GetPost(req.Context(), "12331")
	if err != nil {
		log.Println(err)
		return
	}
	resMap := make(map[string]interface{})
	resMap["code"] = 200
	resMap["uid"] = postResp.Uid
	resMap["text"] = postResp.Text

	data, err := json.Marshal(resMap)
	if err != nil {
		log.Println(err)
		return
	}
	resp.Write(data)
}
