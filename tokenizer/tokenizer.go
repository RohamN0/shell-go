package tokenizer

import (
	"strings"
)

func Tokenize(line string) ([]string) {
    var args [] string
    var current strings.Builder
    inSingle := false
    inDouble := false
    escaped := false

    for i := 0; i < len(line); i++ {
        c := line[i]

        if escaped {
            current.WriteByte(c)
            escaped = false
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

        case ' ', '\t', '\n' :
            if !inSingle && !inDouble {
                if current.Len() > 0 {
                    args = append(args, current.String())
                    current.Reset()
                }
                continue
            }
            current.WriteByte(c)

        default :
            current.WriteByte(c)
        }
    }

    if current.Len() > 0 {
        args = append(args, current.String())
    }

    if inSingle || inDouble {
        return nil
    }

    return args
}

func EchoTokenize(text string) string {
	if strings.Count(text, "'") / 2 != 0 {
		text = strings.ReplaceAll(text, "'", "")
	} else {
		text_splited := strings.Fields(text)
		text = strings.Join(text_splited, " ")
	}

	return text
}
