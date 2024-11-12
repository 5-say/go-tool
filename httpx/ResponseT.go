package httpx

// ResponseT ..
type ResponseT struct {
	Error      error
	Body       []byte
	StatusCode int
}
