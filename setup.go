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
	m.MoveFiles(f)
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

func (m model) MoveFiles(f Framework) {
	m.bibi.Text = "Moving files in the right place"

	err := os.Mkdir("./my-project", 0755)
	err = os.Rename("./tmp/shared/.gitignore", "./my-project/.gitignore")
	err = os.Rename("./tmp/shared/.npmrc", "./my-project/.npmrc")
	err = os.Rename("./tmp/shared/apps", "./my-project/apps")
	err = os.Rename("./tmp/shared/package.json", "./my-project/package.json")
	err = os.Rename("./tmp/shared/packages", "./my-project/packages")
	err = os.Rename("./tmp/shared/pnpm-workspace.yaml", "./my-project/pnpm-workspace.yaml")

	if f.Title == "Nuxt" {
		err = os.Rename("./tmp/nuxt/pnpm-lock.yaml", "./my-project/pnpm-lock.yaml")
		err = os.Rename("./tmp/nuxt/desktop", "./my-project/apps/desktop")
		err = os.Rename("./tmp/shared/src-tauri", "./my-project/apps/desktop/src-tauri")
		err = os.Rename("./tmp/nuxt/tauri.conf.json", "./my-project/apps/desktop/src-tauri/tauri.conf.json")
		err = os.Rename("./tmp/nuxt/ui/components", "./my-project/packages/ui/lib/components")
		err = os.Rename("./tmp/nuxt/ui/index.ts", "./my-project/packages/ui/index.ts")
		err = os.Rename("./tmp/nuxt/ui/package.json", "./my-project/packages/ui/package.json")
		err = os.Rename("./tmp/nuxt/web/package.json", "./my-project/apps/web/package.json")
		err = os.Rename("./tmp/nuxt/web/js", "./my-project/apps/web/resources/js")
		err = os.Rename("./tmp/nuxt/web/vite.config.js", "./my-project/apps/web/vite.config.js")
	}

	CheckIfError(err)
}
