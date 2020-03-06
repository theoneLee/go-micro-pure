package handler

import (
	"context"
	"test.lee/hello_world/model"
	proto "test.lee/hello_world/proto"
)

var dao model.Dao

type Service struct{}

func (Service) Echo(ctx context.Context, req *proto.EchoRequest, resp *proto.EchoResponse) error {
	user, err := dao.GetUser(req.UserName, req.Content)
	if err != nil {
		resp.Success = false
		resp.Error = &proto.Error{
			Code:   500,
			Detail: err.Error(),
		}
		return nil
	}

	resp.User = user
	resp.Success = true
	return nil
}

func (Service) Print(ctx context.Context, req *proto.PrintRequest, resp *proto.PrintResponse) error {
	content, err := dao.GetContent()
	if err != nil {
		resp.Success = false
		resp.Error = &proto.Error{
			Code:   500,
			Detail: err.Error(),
		}
		return nil
	}

	resp.PrintContent = content
	resp.Success = true
	return nil
}
