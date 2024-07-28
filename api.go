package api

import (
	"net/rpc"

	"github.com/natefinch/pie"
	"github.com/urfave/cli/v2"
)

type Lyra struct{
	client *rpc.Client
}

func (l Lyra) Start(id string) error {
	l.client = pie.NewConsumer()
	return l.client.Call("Lyra.Init", id, nil)
}

func (l Lyra) AddCommands(cmds []cli.Command) error {
	return l.client.Call("Lyra.AddCommands", cmds, nil)
}
