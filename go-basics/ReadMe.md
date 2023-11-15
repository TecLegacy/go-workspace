# how to build a go program

1. go build <filename>
2. run executable <filename>.exe
3. go run <filename.go> -> runs file without creating a byte-file

# Go workspaces (Go 1.18+)

Starting with Go 1.18, the go command has native support for multi-module workspaces, via go.work files. These files are recognized by gopls starting with gopls@v0.8.0.

The easiest way to work on multiple modules in Go 1.18 and later is therefore to create a go.work file containing the modules you wish to work on, and set your workspace root to the directory containing the go.work file.

For example, suppose this repo is checked out into the $WORK/tools directory. We can work on both golang.org/x/tools and golang.org/x/tools/gopls simultaneously by creating a go.work file using go work init, followed by go work use MODULE_DIRECTORIES... to add directories containing go.mod files to the workspace:

1. cd $WORK
2. go work init
3. go work use ./tools/ ./tools/gopls/

# default formatter

shift + alt + f
