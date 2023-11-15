# how to build a go program

1. go build <filename>
2. run executable <filename>.exe
3. go run <filename.go> -> runs file without creating a byte-file

4. go mod init <path of module>
5. go work init <directory-path> <- this will only work if go.mod is already present in the working directory that you want to add>

# Go workspaces (Go 1.18+)

Starting with Go 1.18, the go command has native support for multi-module workspaces, via go.work files. These files are recognized by gopls starting with gopls@v0.8.0.

The easiest way to work on multiple modules in Go 1.18 and later is therefore to create a go.work file containing the modules you wish to work on, and set your workspace root to the directory containing the go.work file.

For example, suppose this repo is checked out into the $WORK/tools directory. We can work on both golang.org/x/tools and golang.org/x/tools/gopls simultaneously by creating a go.work file using go work init, followed by go work use MODULE_DIRECTORIES... to add directories containing go.mod files to the workspace:

1. cd $WORK
2. go work init
3. go work use ./tools/ ./tools/gopls/

# default formatter

shift + alt + f

# go mod multi-module workspaces

1. use case - to work with multiple directories

- mkdir test , cd test
- go mod init example/user/test => creates go.mod
- touch test.go -> basic go program to print hello world
- go install example/user/test -> this will create a binary exe & install that binary to home/user/go/bin/test
- now you can run that binary on terminal directly - test // hello world

### 2. go work init ./test

- to work with more than 1 directory
  at the root of dir

  - go work init ./test => this will create a go.work (which will have all the working directory that uses go mentioned)

  "go 2.41
  use ./hello"

The go.work file has similar syntax to go.mod.

The go directive tells Go which version of Go the file should be interpreted with. It’s similar to the go directive in the go.mod file.

The use directive tells Go that the module in the hello directory "should be main modules when doing a build."

So in any subdirectory of workspace the module will be active
