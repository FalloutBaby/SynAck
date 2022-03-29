package delivery

type Delivery interface {
	GetNetwork() string
}

type Http struct {
}

const tcp = "tcp"

func (http Http) GetNetwork() string {
	return tcp
}
