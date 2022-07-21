<div>
  <img src="./docs/todo.svg" width="100%">
</div>

Command line interface program to manage a to do list.

This small program has no real practical use, it is only a learning resource.

Writen in GO.

## Language

[README in Spanish](https://github.com/joshdev15/todo/blob/main/docs/README.ES.md)

## Build and run

### Clone

```bash
git clone https://github.com/joshdev15/todo.git
```

### Run (for developers)

```bash
git run cmd/todo/todo.go <command> <subcommand> <subcommand-value>
```

### Build / Compile

```bash
git build cmd/todo/todo.go
```

Run from the compiled in your root directory

```bash
./todo <command> <subcommand> <subcommand-value>
```

### Installation

If you already have the repository on your computer you can use the command go
install command to install the executable binary to your GOPATH

```bash
git install cmd/todo/todo.go
```

then it is easier to run the program directly with the command `todo`.

```bash
todo <command> <subcommand> <subcommand-value>
```

## Basic commands

[Commands](https://github.com/joshdev15/todo/blob/main/docs/commands.md)
