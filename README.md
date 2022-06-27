# TODO

Programa de Interfaz de línea de comandos para gestionar una lista de tareas por
hacer (to do list).

## Contruccion y ejecucion

### Clonar

```bash
git clone https://github.com/joshdev15/todo.git
```

### Correr (para desarrolladores)

```bash
git run cmd/todo/todo.go <command> <subcommand> <subcommand-value>
```

### Construir / Compilar

```bash
git build cmd/todo/todo.go
```

Ejecurar desde el compilado en tu directorio raíz

```bash
./todo <command> <subcommand> <subcommand-value>
```

### Instalación

Si ya tienes el repositorio en tu ordenador puedes usar el comando go install
para instalar el ejecurable binario en tu GOPATH

```bash
git install cmd/todo/todo.go
```

luego es mas facil correr el programa directamente con el comando "todo"

```bash
todo <command> <subcommand> <subcommand-value>
```
