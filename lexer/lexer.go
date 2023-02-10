package lexer

type Lexer struct {
	input        string
	position     int  // 指向当前字符
	readPosition int  // 指向当前字符之后的一个字符
	ch           byte // 当前正在查看的字符
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

// readChar 读取 input 中的下一个字符
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
