package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"os"
)

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func (m model) Setup(f Framework) {
	m.bibi.Text = "Cloning the bibimbap repo"

	Clone()
	MoveFiles(f)
}

func Clone() {
	r, err := git.PlainClone("./tmp", false, &git.CloneOptions{
		URL:               "https://github.com/theokbokki/bibimbap",
		ReferenceName:     "refs/heads/39-separate-shared-and-framework-specific-files",
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	CheckIfError(err)

	ref, err := r.Head()
	CheckIfError(err)
	_, err = r.CommitObject(ref.Hash())
	CheckIfError(err)
}

func MoveFiles(f Framework) {
	err := os.Rename("./tmp/shared/.gitignore", "./my-project/.gitignore")
	err = os.Rename("./tmp/shared/.npmrc", "./my-project/.npmrc")
	err = os.Rename("./tmp/shared/apps", "./my-project/apps")
	err = os.Rename("./tmp/shared/package.json", "./my-project/package.json")
	err = os.Rename("./tmp/shared/packages", "./my-project/packages")
	err = os.Rename("./tmp/shared/pnpm-workspace.yaml", "./my-project/pnpm-workspace.yaml")

	CheckIfError(err)
}
