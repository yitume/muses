package muses

import (
	"errors"
	"io/ioutil"
	"reflect"

	"github.com/yitume/muses/pkg/common"
)

// 通过反射取包里面的值
var orderCallerList = []callerAttr{
	{true, common.ModAppName},
	{true, common.ModLoggerName},
	{false, common.ModMysqlName},
	{false, common.ModRedisName},
	{false, common.ModMongoName},
	{false, common.ModGinSessionName},
	{false, common.ModEchoSessionName},
	{false, common.ModStatName},
	{false, common.ModGinName},
}

type callerAttr struct {
	IsNecessary bool
	Name        string
}

// Container from file.
func parseFile(path string) ([]byte, error) {
	// read file to []byte
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return b, err
	}
	return b, nil
}

func sortCallers(callers []common.CallerFunc) (callerSort []common.Caller, err error) {
	callerMap := make(map[string]common.Caller)
	callerSort = make([]common.Caller, 0)
	for _, caller := range callers {
		obj := caller()
		name := getCallerName(obj)
		callerMap[name] = obj
	}

	for _, value := range orderCallerList {
		caller, ok := callerMap[value.Name]
		if !ok {
			// 如果是必须加载的组件
			if value.IsNecessary {
				err = errors.New(value.Name + " is not exist")
				return
			}
		} else {
			// 如果存在于map，加入到排序里的caller sort
			callerSort = append(callerSort, caller)
		}
	}
	return
}

func getCallerName(caller common.Caller) string {
	return reflect.ValueOf(caller).Elem().FieldByName("Name").String()
}

func isCallerBackground(caller common.Caller) bool {
	return reflect.ValueOf(caller).Elem().FieldByName("IsBackground").Bool()
}
