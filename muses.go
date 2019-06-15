package muses

import (
	"fmt"
	"github.com/yitume/muses/pkg/common"
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

	callers, err := sortCallers(callerFuncs)
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
		if isCallerBackground(caller) {
			go func() {
				if err = caller.InitCaller(); err != nil {
					fmt.Println("module", name, "init caller error")
					return
				}
			}()
		} else {
			if err = caller.InitCaller(); err != nil {
				fmt.Println("module", name, "init caller error")
				return
			}
		}
		fmt.Println("module", name, "init caller ok")
		fmt.Println("module", name, "end")
	}
	return nil
}
