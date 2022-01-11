package internal

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (s Server) Run() error {
	return nil
}
