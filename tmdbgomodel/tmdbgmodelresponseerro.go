package tmdbgomodel

//
type TmdbgModelResponseError struct {
	Success        bool   `json:"success"`
	Status_Code    int    `json:"status_code"`
	Status_Message string `json:"status_message"`
}
