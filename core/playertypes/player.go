package playertypes

type Player interface {
	Symbol() string
	Move() interface{}
}
