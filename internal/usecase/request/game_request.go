package request

type GameRequest struct {
	ID      int
	Players []*PlayerForm
}
