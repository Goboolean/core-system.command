package msg

import "github.com/Goboolean/command-server/internal/domain/entity"

func ServerError() *entity.TaskAllocResponse {
	return &entity.TaskAllocResponse{
		Status:      500,
		AccessToken: "",
	}
}

func NotFound() *entity.TaskAllocResponse {
	return &entity.TaskAllocResponse{
		Status:      404,
		AccessToken: "",
	}
}

func Success(uuid string) *entity.TaskAllocResponse {
	return &entity.TaskAllocResponse{
		Status:      200,
		AccessToken: uuid,
	}
}

func NewMsg(stat int32, accTok string) *entity.TaskAllocResponse {
	return &entity.TaskAllocResponse{
		Status:      stat,
		AccessToken: accTok,
	}
}

func String(res *entity.TaskAllocResponse) string {
	return "status: " + string(res.Status) + "\n" + "access token: " + res.AccessToken
}
