package main

import (
	"os"
)

func main() {
    prepareConfigEnv()

    args := os.Args[1:]

    if len(args) == 0 {
        handleHelp()
        return
    }

    command := args[0]

    switch(command) {
    case "edit":
        app := args[1]
        handleEdit(app)
        break

    case "list":
        handleList()
        break

    case "export":
        handleExport(args[1])
        break

    case "import":
        handleImport(args[1])
        break

    default:
        app := args[0]
        handleRead(app)
        break
    }
}
