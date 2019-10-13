package httperror

type Error struct {
	HttpCode int    `json:"HttpCode"`
	Message  string `json:"Message" bson:"Message"`
}

func New(httpCode int, message string) *Error {
	return &Error{httpCode, message}
}
