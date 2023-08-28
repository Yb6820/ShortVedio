package main

import (
	_ "gfVersion/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gfVersion/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
