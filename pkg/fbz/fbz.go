package fbz

type Driver interface {
	Post(string, Params, []byte) Response
	Token() string
}

type Params map[string][]string

func (params Params) Set(key string, value string) {
	params[key] = []string{value}
}

type Response struct {
	Data  []byte
	Error error
}

func (response Response) Okay() bool {
	if response.Error == nil {
		return true
	}
	return false
}
