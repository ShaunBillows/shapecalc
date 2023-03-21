package app

type StringReader interface {
	ReadString(delim byte) (string, error)
}
