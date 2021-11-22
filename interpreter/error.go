package interpreter

import (
   "dexianta/glox/scanner"
   "fmt"
)

type RuntimeError struct {
   Token scanner.Token
   Msg string
}

func (r RuntimeError) Error() string {
   return fmt.Sprintf("%v: %s\n", r.Token, r.Msg)
}
