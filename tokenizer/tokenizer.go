package tokenizer

import (
	"strings"
	"unicode"
)

func Tokenize(line string) ([]string) {
    var args [] string
    var current strings.Builder
    inSingle := false
    inDouble := false

    for i := 0; i < len(line); i++ {
        c := line[i]

		if c == '\\' && !inSingle {
			if i + 1 < len(line) {
				next := line[i + 1]
				current.WriteByte(next)
				i++
				continue
			}
			current.WriteByte('\\')
			continue
		}

        switch c {
        case '\'' :
            if !inDouble {
                inSingle = !inSingle
                continue
            }
            current.WriteByte(c)

        case '"' :
            if !inSingle {
                inDouble = !inDouble
                continue
            }
            current.WriteByte(c)

        default :
            if unicode.IsSpace(rune(c)) && !inSingle && !inDouble {
                if current.Len() > 0 {
                args = append(args, current.String())
                    current.Reset()
                }
            } else {
                current.WriteByte(c)
            }
        }
    }

    if current.Len() > 0 {
        args = append(args, current.String())
    }

    return args
}
