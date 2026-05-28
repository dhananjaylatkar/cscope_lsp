package lsp

type ShutdownRequest struct {
	Request
}

type ShutdownResponse struct {
	Response

	// Must be null.
	Result *int `json:"result"`
}

func NewShutdownResponse(id int) ShutdownResponse {
	return ShutdownResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
	}
}
