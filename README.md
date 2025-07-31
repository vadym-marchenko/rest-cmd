# rest-cmd

A Go CLI application built with [Cobra](https://github.com/spf13/cobra).

## Getting Started

1. **Install Go**: https://golang.org/dl/
2. **Build the app:**
   ```sh
   go build -o rest-cmd
   ```
3. **Run the app:**
   ```sh
   ./rest-cmd
   ```

## Usage

- Run `rest-cmd help` to see available commands.
- Example subcommand: `rest-cmd example`

## Project Structure
- `cmd/` - Cobra commands live here.
- `main.go` - Entry point.

## Customization
Add new commands with:
```sh
cobra-cli add <command-name>
```

## License
MIT
