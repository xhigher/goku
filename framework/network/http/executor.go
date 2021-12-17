package http

import (
	"io/ioutil"
	"log"
	"net/http"

	"go.uber.org/zap"
	"goku.net/framework/commons"
	"goku.net/utils"
)

type Executor interface {
	SupportMethods() []string
	CheckQueryValues(values map[string][]string) bool
	CheckBodyData(data string) bool
	GetBaseData() *BaseBodyData
	RequireSession() bool
	Execute() ResponseData
}

type BaseExecutor struct {
	QueryValues map[string][]string
	BodyData    interface{}
	BaseData    *BaseBodyData
	Version     int
}

func (executor *BaseExecutor) SupportMethods() []string {
	return []string{"GET", "POST"}
}

func (executor *BaseExecutor) CheckQueryValues(values map[string][]string) bool {
	if executor.QueryValues == nil {
		executor.QueryValues = values
	}
	return true
}

func (executor *BaseExecutor) CheckBodyData(data string) bool {
	if executor.BaseData != nil {
		executor.BaseData = &BaseBodyData{}
		err := utils.ParseStruct(data, executor.BaseData)
		if err != nil {
			commons.Logger().Error("CheckBodyData", zap.String("data", data), zap.Any("err", err))
			return false
		}
	}
	if executor.BodyData != nil {
		err := utils.ParseStruct(data, executor.BodyData)
		if err != nil {
			commons.Logger().Error("CheckBodyData", zap.String("data", data), zap.Any("err", err))
			return false
		}
	}
	return true
}

func (executor *BaseExecutor) QueryValue(key string) (value interface{}, ok bool) {
	if executor.QueryValues == nil {
		value, ok = executor.QueryValues[key]
	}
	return
}

func (executor *BaseExecutor) GetBaseData() *BaseBodyData {
	return executor.BaseData
}

func (executor *BaseExecutor) RequireSession() bool {
	return false
}

func (executor *BaseExecutor) ResultOK() ResponseData {
	return executor.OutputResult(OK, "", nil)
}

func (executor *BaseExecutor) ResultOKData(data interface{}) ResponseData {
	return executor.OutputResult(OK, "", data)
}

func (executor *BaseExecutor) ResultError() ResponseData {
	return executor.OutputResult(NOK, "", nil)
}

func (executor *BaseExecutor) ResultErrorInternal() ResponseData {
	return executor.OutputResult(ERR_INTERNAL, "ERR_INTERNAL", nil)
}

func (executor *BaseExecutor) ResultErrorRequest() ResponseData {
	return executor.OutputResult(ERR_REQUEST, "ERR_REQUEST", nil)
}

func (executor *BaseExecutor) ResultErrorParameter(msg string) ResponseData {
	return executor.OutputResult(ERR_PARAMETER, msg, nil)
}

func (executor *BaseExecutor) ResultErrorSession() ResponseData {
	return executor.OutputResult(ERR_SESSION, "ERR_SESSION", nil)
}

func (executor *BaseExecutor) ResultErrorCaptcha() ResponseData {
	return executor.OutputResult(ERR_CAPTCHA, "ERR_CAPTCHA", nil)
}

func (executor *BaseExecutor) ResultErrorFrequently() ResponseData {
	return executor.OutputResult(ERR_FREQUENTLY, "ERR_FREQUENTLY", nil)
}

func (executor *BaseExecutor) OutputResult(code ErrorCode, msg string, data interface{}) ResponseData {
	return ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

type ModuleExecutorFactory interface {
	ModuleName() string
	Create(version int, action string) Executor
}

type LogicExecutor struct {
	executor Executor
}

func (logic *LogicExecutor) CheckMethod(method string) bool {
	supportMethods := logic.executor.SupportMethods()
	for _, mtd := range supportMethods {
		if mtd == method {
			return true
		}
	}
	return false
}

func (logic *LogicExecutor) CheckParams(request *http.Request) bool {
	if !logic.executor.CheckQueryValues(request.URL.Query()) {
		return false
	}
	contentType := request.Header.Get(ContentType)
	if IsJSONData(contentType) {
		var bodyData, err = ioutil.ReadAll(request.Body)
		if err != nil {
			log.Print("Exception for parsing json parameters", err.Error())
			return false
		}
		if !logic.executor.CheckBodyData(utils.BytesToString(bodyData)) {
			return false
		}
	}
	return true
}

func (logic *LogicExecutor) Execute(write http.ResponseWriter) {

	result := logic.executor.Execute()

	resultString, err := utils.ToJSONString(result)
	if err != nil {
		write.WriteHeader(502)
		return
	}

	write.Header().Set(ContentType, "application/json;charset=UTF-8")
	write.Write(utils.StringToBytes(resultString))
}
