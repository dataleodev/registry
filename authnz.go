package registry

type AuthNZ interface {
	Issue()
	Identify()
	Retrieve()
	Authorize()
}
