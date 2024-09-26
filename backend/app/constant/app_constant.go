package constant

type ResponseStatus int
type Headers int
type General int
type Duration string
type Channel int
type MessageStatus string
type SponsoredLevel string
type KdpScoreAttr string

const (
	Success ResponseStatus = iota + 1
	DataNotFound
	UnknownError
	BadRequest
	Unauthorized
	Forbidden
	InternalServer
)

func (r ResponseStatus) GetResponseStatus() string {
	return [...]string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "INVALID_REQUEST", "UNAUTHORIZED", "FORBIDDEN", "INTERNAL_SERVER"}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{"Success", "Data Not Found", "Unknown Error", "Invalid Request", "Unauthorized", "Forbidden", "Internal Server"}[r-1]
}
