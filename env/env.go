package env

import (
	"flag"
	"fmt"
	"strings"
)

var (
	active Environment
	dev    Environment = &environment{value: "dev"}
	prod   Environment = &environment{value: "prod"}
)

var _ Environment = (*environment)(nil)

type Environment interface {
	Value() string
	IsDev() bool
	IsProd() bool
	t()
}

type environment struct {
	value string
}

func (env *environment) Value() string {
	return env.value
}

func (env *environment) IsDev() bool {
	return env.value == "dev"
}

func (env *environment) IsProd() bool {
	return env.value == "prod"
}

func (env *environment) t() {}

func init() {
	env := flag.String("env", "", "请输入运行环境:\n dev:开发环境\n prod:生产环境\n")
	flag.Parse()

	switch strings.ToLower(strings.TrimSpace(*env)) {
	case "dev":
		active = dev
	case "prod":
		active = prod
	default:
		active = dev
		fmt.Println("默认运行环境为dev")
	}
}

func Active() Environment {
	return active
}
