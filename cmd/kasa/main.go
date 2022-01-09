package main

import (
	"github.com/alecthomas/kong"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa"
	"github.com/winebarrel/kasa/subcmd"
)

var version string

var cli struct {
	Version kong.VersionFlag
	Team    string           `required:"" env:"ESA_TEAM" help:"esa team"`
	Token   string           `required:"" env:"ESA_TOKEN" help:"esa access token"`
	Cat     subcmd.CatCmd    `cmd:"" help:"Print post."`
	Info    subcmd.InfoCmd   `cmd:"" help:"Show post info."`
	Ls      subcmd.LsCmd     `cmd:"" help:"List posts."`
	Mv      subcmd.MvCmd     `cmd:"" help:"Move posts."`
	Mvcat   subcmd.MvcatCmd  `cmd:"" help:"Move category."`
	Post    subcmd.PostCmd   `cmd:"" help:"New post."`
	Rm      subcmd.RmCmd     `cmd:"" help:"Delete post."`
	Rmx     subcmd.RmxCmd    `cmd:"" help:"Delete multiple posts."`
	Search  subcmd.SearchCmd `cmd:"" help:"Search posts."`
}

func main() {
	ctx := kong.Parse(&cli, kong.Vars{"version": version})

	err := ctx.Run(&kasa.Context{
		Driver: esa.NewDriver(cli.Team, cli.Token),
	})

	ctx.FatalIfErrorf(err)
}
