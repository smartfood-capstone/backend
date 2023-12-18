package storage

type Storage struct {
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Store()  {}
func (s *Storage) Get()    {}
func (s *Storage) Update() {}
