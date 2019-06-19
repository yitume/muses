package beego

import (
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yitume/muses/pkg/common"
)

var defaultCaller = &callerStore{
	Name: common.ModTmplBeegoName,
	cfg: Cfg{
		Muses: CallerMuses{
			Tmpl: CallerTmpl{
				Beego: CallerCfg{
					Debug:         false,
					TplExt:        "tpl",
					ViewPath:      "views",
					TemplateLeft:  "{{",
					TemplateRight: "}}",
				},
			},
		},
	},
}

type callerStore struct {
	Name string
	cfg  Cfg
}

func Register() common.Caller {
	return defaultCaller
}

func Caller(tplPath string) (obj *Tmpl, err error) {
	obj, err = provider(defaultCaller.cfg.Muses.Tmpl.Beego, tplPath)
	if err != nil {
		return
	}
	return
}

func Config() CallerCfg {
	return defaultCaller.cfg.Muses.Tmpl.Beego
}

func (c *callerStore) InitCfg(cfg []byte) error {
	if err := toml.Unmarshal(cfg, &c.cfg); err != nil {
		return err
	}
	return nil
}

func (c *callerStore) InitCaller() error {
	if err := AddViewPath(c.cfg.Muses.Tmpl.Beego.ViewPath); err != nil {
		return err
	}
	return nil
}

func provider(cfg CallerCfg, tplPath string) (resp *Tmpl, err error) {
	obj := &Tmpl{}
	obj.Init(tplPath, cfg.TplExt, cfg.ViewPath)

	return obj, err
}
