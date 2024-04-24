package transport

// Transport is object that handling the communication in a network
type Transport interface {
	ListenAndAccept() error
}
