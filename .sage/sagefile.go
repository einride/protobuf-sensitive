package main

import (
	"context"

	"go.einride.tech/sage/sg"
	"go.einride.tech/sage/tools/sgbuf"
	"go.einride.tech/sage/tools/sgclangformat"
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
	sg.Deps(ctx, ClangFormatProto, BufLint)
	return nil
}

func ClangFormatProto(ctx context.Context) error {
	sg.Logger(ctx).Println("formatting proto files...")
	return sgclangformat.FormatProto(ctx)
}

func BufLint(ctx context.Context) error {
	sg.Logger(ctx).Println("linting proto files..")
	return sgbuf.Command(ctx, "lint").Run()
}

func BufPush(ctx context.Context) error {
	sg.Logger(ctx).Println("pushing Buf module..")
	return sgbuf.Command(ctx, "push").Run()
}
