package tmdbgomodel

//
type TmdbPopularFilmModel struct {
	Adult             bool     `json:"adult"`
	Backdrop_Path     string   `json:"backdrop_path"`
	Genre_Ids         []uint16 `json:"genre_ids"`
	Id                uint32   `json:"id"`
	Original_Language string   `json:"original_language"`
	Original_Title    string   `json:"original_title"`
	Overview          string   `json:"overview"`
	Popularity        float64  `json:"popularity"`
	Poster_Path       string   `json:"poster_path"`
	Release_Date      string   `json:"release_date"`
	Title             string   `json:"title"`
	Video             bool     `json:"video"`
	Vote_Average      float64  `json:"vote_average"`
	Vote_Count        int      `json:"vote_count"`
}

//
type TmdbListPopularFilmModel struct {
	Page          int                    `json:"page"`
	Results       []TmdbPopularFilmModel `json:"results"`
	Total_Pages   int                    `json:"total_pages"`
	Total_Results int                    `json:"total_results"`
}
