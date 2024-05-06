package lsp

type ReferencesRequest struct {
	Request
	Params ReferencesParams `json:"params"`
}

type ReferencesParams struct {
	TextDocumentPositionParams
}

type ReferencesResponse struct {
	Response
	Result []Location `json:"result"`
}
