package main

import (
	"context"

	"go.einride.tech/sage/sg"
	"go.einride.tech/sage/tools/sgbuf"
	"go.einride.tech/sage/tools/sggolangcilint"
)

func main() {
	sg.GenerateMakefiles(
		sg.Makefile{
			Path:          sg.FromGitRoot("Makefile"),
			DefaultTarget: All,
		},
	)
}

func All(ctx context.Context) error {
	sg.Deps(ctx, BufFormat, BufLint, GoLint)
	return nil
}

func BufFormat(ctx context.Context) error {
	sg.Logger(ctx).Println("formatting proto files...")
	return sgbuf.Command(ctx, "format", "-w").Run()
}

func BufLint(ctx context.Context) error {
	sg.Logger(ctx).Println("linting proto files..")
	return sgbuf.Command(ctx, "lint").Run()
}

func BufPush(ctx context.Context) error {
	sg.Logger(ctx).Println("pushing Buf module..")
	return sgbuf.Command(ctx, "push").Run()
}

func GoLint(ctx context.Context) error {
	sg.Logger(ctx).Println("linting Go modules..")
	return sggolangcilint.Run(ctx)
}
