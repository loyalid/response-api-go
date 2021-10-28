package response

type API struct {
	StatusCode string         `json:"status_code"`
	StatusText string         `json:"status_text"`
	Message    string         `json:"message"`
	Data       interface{}    `json:"data"`
	Paginator  map[string]int `json:"paginator"`
}

func (response *API) Simple(statusCode string, statusText string, message string) {
	response.StatusCode = statusCode
	response.StatusText = statusText
	response.Message = message
}

func (response *API) WithData(statusCode string, statusText string, message string, data interface{}) {
	response.StatusCode = statusCode
	response.StatusText = statusText
	response.Message = message
	response.Data = data
}

func (response *API) WithPagination(statusCode string, statusText string, message string, data interface{}, paginator map[string]int) {
	response.StatusCode = statusCode
	response.StatusText = statusText
	response.Message = message
	response.Data = data
	response.Paginator = map[string]int{
		"per_page":      paginator["PageSize"],
		"current_page":  paginator["Page"],
		"previous_page": paginator["Page"] - 1,
		"next_page":     paginator["Page"] + 1,
	}
}
