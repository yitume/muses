package beego

type Cfg struct {
	Muses CallerMuses `toml:"muses"`
}

type CallerMuses struct {
	Tmpl CallerTmpl `toml:"tmpl"`
}

type CallerTmpl struct {
	Beego CallerCfg `toml:"beego"`
}

type CallerCfg struct {
	Debug         bool
	TplExt        string
	ViewPath      string
	TemplateLeft  string
	TemplateRight string
}
