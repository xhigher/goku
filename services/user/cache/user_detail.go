package cache

import (
	"go.uber.org/zap"
	"goku.net/framework/commons"
	proto_cache "goku.net/protos/cache"
	proto_model "goku.net/protos/model"
)

func GetUserDetail(mid int64) (data proto_model.UserDetailModel, err error) {
	result := Cache().Get(proto_cache.UserKeys.Detail(mid))
	if !result.OK() {
		err = result.Error()
		return
	}
	if result.Exist() {
		err = result.String2Struct(&data)
	}
	commons.Logger().Info("cache.GetUserDetail", zap.Any("data", data))
	return
}
