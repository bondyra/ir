package command

import (
	"strings"

	"github.com/bondyra/wtf/internal/config"
	"github.com/bondyra/wtf/internal/handler"
)

func Execute(args []string) {
	switch args[0] {
	case "config":
		panic("dupa")
	default:
		handler.QueryHandler{Query: strings.Join(args, " ")}.Execute(config.Config{})
	}
}
