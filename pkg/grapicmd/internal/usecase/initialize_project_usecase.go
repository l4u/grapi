package usecase

import (
	"context"

	"github.com/izumin5210/gex"
	"github.com/pkg/errors"

	"github.com/izumin5210/grapi/pkg/cli"
	"github.com/izumin5210/grapi/pkg/grapicmd/internal/module"
)

// InitializeProjectUsecase is an interface to create a new grapi project.
type InitializeProjectUsecase interface {
	Perform(rootDir string, pkgName string) error
	GenerateProject(rootDir, pkgName string) error
	InstallDeps(rootDir string) error
}

// NewInitializeProjectUsecase creates a new InitializeProjectUsecase instance.
func NewInitializeProjectUsecase(ui cli.UI, generator module.ProjectGenerator, gexCfg *gex.Config) InitializeProjectUsecase {
	return &initializeProjectUsecase{
		ui:        ui,
		generator: generator,
		gexCfg:    gexCfg,
	}
}

type initializeProjectUsecase struct {
	ui        cli.UI
	generator module.ProjectGenerator
	gexCfg    *gex.Config
}

func (u *initializeProjectUsecase) Perform(rootDir, pkgName string) error {
	u.ui.Section("Initialize project")

	var err error
	err = u.GenerateProject(rootDir, pkgName)
	if err != nil {
		return errors.Wrap(err, "failed to initialize project")
	}

	u.ui.Subsection("Install dependencies")
	err = u.InstallDeps(rootDir)
	if err != nil {
		return errors.Wrap(err, "failed to execute `dep ensure`")
	}

	return nil
}

func (u *initializeProjectUsecase) GenerateProject(rootDir, pkgName string) error {
	return errors.WithStack(u.generator.GenerateProject(rootDir, pkgName))
}

func (u *initializeProjectUsecase) InstallDeps(rootDir string) error {
	u.gexCfg.WorkingDir = rootDir
	repo, err := u.gexCfg.Create()
	if err == nil {
		err = repo.Add(
			context.TODO(),
			"github.com/izumin5210/grapi/cmd/grapi",
			// TODO: make configurable
			"github.com/golang/protobuf/protoc-gen-go",
			"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway",
			"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger",
		)
	}
	return errors.WithStack(err)
}
