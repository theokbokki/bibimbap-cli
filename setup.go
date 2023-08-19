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
		URL: "https://github.com/theokbokki/bibimbap",
		// ReferenceName:     "refs/heads/39-separate-shared-and-framework-specific-files",
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	CheckIfError(err)

	ref, err := r.Head()
	CheckIfError(err)
	_, err = r.CommitObject(ref.Hash())
	CheckIfError(err)
}

func (m model) MoveFiles(f Framework) {
	err := os.Mkdir("./"+m.textinput.Value(), 0755)

	err = os.Rename("./tmp/shared/.gitignore", "./"+m.textinput.Value()+"/.gitignore")
	err = os.Rename("./tmp/shared/.npmrc", "./"+m.textinput.Value()+"/.npmrc")
	err = os.Rename("./tmp/shared/apps", "./"+m.textinput.Value()+"/apps")
	err = os.Rename("./tmp/shared/package.json", "./"+m.textinput.Value()+"/package.json")
	err = os.Rename("./tmp/shared/packages", "./"+m.textinput.Value()+"/packages")
	err = os.Rename("./tmp/shared/pnpm-workspace.yaml", "./"+m.textinput.Value()+"/pnpm-workspace.yaml")

	if f.Title == "Nuxt" {
		err = os.Rename("./tmp/nuxt/pnpm-lock.yaml", "./"+m.textinput.Value()+"/pnpm-lock.yaml")
		err = os.Rename("./tmp/nuxt/desktop", "./"+m.textinput.Value()+"/apps/desktop")
		err = os.Rename("./tmp/nuxt/tauri.conf.json", "./"+m.textinput.Value()+"/apps/desktop/src-tauri/tauri.conf.json")
		err = os.Rename("./tmp/nuxt/ui/components", "./"+m.textinput.Value()+"/packages/ui/lib/components")
		err = os.Rename("./tmp/nuxt/ui/index.ts", "./"+m.textinput.Value()+"/packages/ui/index.ts")
		err = os.Rename("./tmp/nuxt/ui/package.json", "./"+m.textinput.Value()+"/packages/ui/package.json")
		err = os.Rename("./tmp/nuxt/web/package.json", "./"+m.textinput.Value()+"/apps/web/package.json")
		err = os.Rename("./tmp/nuxt/web/js", "./"+m.textinput.Value()+"/apps/web/resources/js")
		err = os.Rename("./tmp/nuxt/web/vite.config.js", "./"+m.textinput.Value()+"/apps/web/vite.config.js")
	}

	_, err = os.Stat("./" + m.textinput.Value() + "/apps/desktop")
	if os.IsNotExist(err) {
		err = os.Mkdir("./"+m.textinput.Value()+"/apps/desktop", 0755)
	}
	err = os.Rename("./tmp/shared/src-tauri", "./"+m.textinput.Value()+"/apps/desktop/src-tauri")

	os.Remove("./tmp")

	m.bibi.Text = "Setup complete! Have fun :))"

	CheckIfError(err)
}
