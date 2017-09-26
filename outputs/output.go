package outputs

type Output interface {
	Send(logs []*HttpLoggerClient) error
}
