package models

import (
	"fmt"
	"strings"

	"time"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	cfg, err := config.NewConfig(`json`, `conf/database.json`)
	if err != nil {
		panic(err)
	}

	connectionName := cfg.String("default")
	conn := make(map[string]string)
	for _, k := range []string{"driver", "host", "port", "database", "charset", "collation", "username", "password"} {
		conn[k] = cfg.String(strings.Join([]string{"connections", connectionName, k}, "::"))
	}

	orm.DefaultTimeLoc = time.UTC

	orm.RegisterDriver("myql", orm.DRMySQL)

	orm.RegisterDataBase("default", conn[`driver`], fmt.Sprintf(
		"%s:%s@/%s?charset=%s",
		conn[`username`],
		conn[`password`],
		conn[`database`],
		conn[`charset`],
	))

	orm.RegisterModel(&Document{})

	orm.RunSyncdb("default", false, true)
}
