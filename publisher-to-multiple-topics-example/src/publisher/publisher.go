package publisher

//Publisher interface
type Publisher interface {
	Publish(payload []byte, headers map[string]string)
}
