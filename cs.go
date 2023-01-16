package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func getConfigPath() string {
    base := os.Getenv("HOME")
    return path.Join(base, "/.config/cs")
}


func prepareConfigEnv() {
    _path := getConfigPath()
    os.MkdirAll(_path, os.ModePerm)
}

func handleEdit(app string) {
    filePath := path.Join(getConfigPath(), app)
    editor := os.Getenv("EDITOR")

    cmd := exec.Command(editor, filePath)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func handleRead(app string) {
    filePath := path.Join(getConfigPath(), app)

    content, err := ioutil.ReadFile(filePath)

    if err != nil {
        fmt.Printf("There's no content for %s\n", app)
        return
    }

    fmt.Print(string(content))
}

func handleList() {
    files, err := ioutil.ReadDir(getConfigPath())

    if err != nil {
        return
    }

    for _, file := range files {
        fmt.Println(file.Name())
    }
}

func handleHelp() {
    fmt.Print(`Examples:

List cheatsheets:
$ cs list

Edit or add a cheatsheet:
$ cs edit git
$ cs edit my-command

Read some cheatsheet (example git):
$ cs git
$ cs my-command
`)
}

func handleExport(exportToPath string) {
    _path, err := filepath.Abs(exportToPath)

    if err != nil {
        fmt.Printf("The folder %s does not exist\n", _path)
        return
    }

    files, err := ioutil.ReadDir(getConfigPath())

    if err != nil {
        return
    }

    for _, file := range files {
        source := path.Join(getConfigPath(), file.Name())
        dist := path.Join(_path, file.Name())

        data, _ := ioutil.ReadFile(source)

        ioutil.WriteFile(dist, data, 0644)
    }
}
