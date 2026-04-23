package repl

import (
	"fmt"
	"io"

	"github.com/WedelFelix/Monkey-Lang-Interpreter-Go/lexer"
	"github.com/WedelFelix/Monkey-Lang-Interpreter-Go/token"
	"github.com/chzyer/readline"
)

func Start() {
	// Configure readline with a prompt and a history file
	rl, err := readline.NewEx(&readline.Config{
		Prompt:      ">> ",
		HistoryFile: "/tmp/readline.tmp", // Enables persistent history
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			if err == readline.ErrInterrupt || err == io.EOF {
				break // Exit on Ctrl+C or Ctrl+D
			}
			break
		}

		l := lexer.New(line)
		fmt.Printf("Input: %s\n", line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
