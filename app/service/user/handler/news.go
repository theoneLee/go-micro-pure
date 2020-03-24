package handler

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/pborman/uuid"
	"test.lee/common/proto/user"
	"test.lee/user/model"
)

var dao model.UserSQLDao = model.UserSQLDao{}

type Service struct {
}

func (s Service) SaveUser(ctx context.Context, req *user.UserReq, resp *user.UserResp) error {
	//从req包装一个user
	userW := model.User{
		ID:       uuid.NewUUID().String(),
		UserName: req.User.UserName,
		//Password:req.User.P
		Password: req.User.UserName,
	}
	log.Info(userW)
	dao.SaveUser(userW)
	var err error
	if err != nil {
		resp.Success = false
		resp.Error = &user.Error{
			Code:   500,
			Detail: err.Error(),
		}
		return nil
	}

	// todo 从user打包成protobuf定义的结构,但是这样做也太复杂了吧？怎么处理这个？https://blog.csdn.net/jxust_tj/article/details/84754345
	// todo 可以尝试编辑protobuf生成的model的tag？然后dao操作的model是proto生成的？https://github.com/gogo/protobuf/blob/master/extensions.md#more-serialization-formats 或者 https://github.com/favadi/protoc-go-inject-tag 推荐用前者。甚至可以尝试装了之后直接用，然后后接--micro_out=. 看下情况
	// todo 但是实际上采取这个方案有些问题。因为协议上的字段定义有可能和自己实体字段不一致的。为了少些一遍实体代码就考虑共用也不见的好。
	// todo 顺便一提，这块rpcx采用的messagepack比proto好很多，就写代码的流畅性来说。
	//resp.User=&userW //编译错误
	resp.User = packUser(userW) //从user打包成protobuf定义的结构
	resp.Success = true
	return nil
}

func packUser(userW model.User) *user.User {
	return &user.User{UserName: userW.UserName, Id: userW.ID}
}

func (s Service) GetUser(context.Context, *user.UserReq, *user.UserResp) error {
	panic("implement me")
}

func (s Service) SaveNotice(context.Context, *user.NoticeReq, *user.NoticeResp) error {
	panic("implement me")
}

//
