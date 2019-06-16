package muses

import (
	"fmt"
	"github.com/yitume/muses/pkg/app"
	"github.com/yitume/muses/pkg/common"
	"github.com/yitume/muses/pkg/logger"
	"github.com/yitume/muses/pkg/prom"
)

func Container(cfg interface{}, callerFuncs ...common.CallerFunc) (err error) {
	var cfgByte []byte
	switch cfg.(type) {
	case string:
		cfgByte, err = parseFile(cfg.(string))
		if err != nil {
			return
		}
	case []byte:
		cfgByte = cfg.([]byte)
	default:
		return fmt.Errorf("type is error %s", cfg)
	}

	allCallers := []common.CallerFunc{app.Register, logger.Register, prom.Register}
	allCallers = append(allCallers, callerFuncs...)

	callers, err := sortCallers(allCallers)
	if err != nil {
		return
	}

	for _, caller := range callers {
		name := getCallerName(caller)
		fmt.Println("module", name, "start")
		if err = caller.InitCfg(cfgByte); err != nil {
			fmt.Println("module", name, "init config error")
			return
		}
		fmt.Println("module", name, "init config ok")
		if err = caller.InitCaller(); err != nil {
			fmt.Println("module", name, "init caller error")
			return
		}
		fmt.Println("module", name, "init caller ok")
		fmt.Println("module", name, "end")
	}
	return nil
}
