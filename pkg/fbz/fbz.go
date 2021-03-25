package fbz

type Driver interface {
	Post(string, []byte) Response
	Token() string
}

type Response interface {
	Okay() bool
	Data() []byte
	Error() error
	Header(string) string
}
