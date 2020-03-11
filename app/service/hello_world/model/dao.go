package model

import (
	proto "test.lee/common/proto/hello_world"
)

type Dao struct {
}

func (Dao) GetUser(UserName, Content string) (*proto.User, error) {
	return &proto.User{SayContent: Content, UserName: UserName, Id: 111}, nil
}

func (Dao) GetContent() (string, error) {
	return "ssss", nil
}
