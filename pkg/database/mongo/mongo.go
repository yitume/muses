package mongo

import (
	"fmt"
	"os"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/globalsign/mgo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yitume/muses/pkg/common"
)

var defaultCaller = &callerStore{
	Name: common.ModMongoName,
}

type callerStore struct {
	Name         string
	IsBackground bool
	caller       sync.Map
	cfg          Cfg
}

type Client struct {
	*mgo.Database
}

func Register() common.Caller {

	return defaultCaller
}

func Caller(name string) *Client {
	obj, ok := defaultCaller.caller.Load(name)
	if !ok {
		return nil
	}
	return obj.(*Client)
}

func (c *callerStore) InitCfg(cfg []byte) error {
	if err := toml.Unmarshal(cfg, &c.cfg); err != nil {
		return err
	}
	return nil
}

func (c *callerStore) InitCaller() error {
	for name, cfg := range c.cfg.Muses.Mongo {
		db, err := provider(cfg)
		if err != nil {
			return err
		}
		defaultCaller.caller.Store(name, db)
	}
	return nil
}

func provider(cfg CallerCfg) (resp *Client, err error) {
	session, err := mgo.Dial(cfg.URL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	fmt.Println("cfg.debug", cfg.Debug)
	mgo.SetDebug(cfg.Debug)
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return &Client{session.DB(cfg.Database)}, err
}