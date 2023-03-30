package prompter

type StringReader interface {
	ReadString(delim byte) (string, error)
}
