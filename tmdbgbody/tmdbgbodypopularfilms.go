package tmdbgbody

import (
	"encoding/json"
	"fmt"
	"strconv"
	"tmdbGotutorial/tmdbgomodel"
	"tmdbGotutorial/tmdbgutil"
)

var ERROR_HTTP_RESPONSE string = `<!DOCTYPE html>
	<html lang="pt-br">
		<head>
			<meta charset="utf-8">
			<title>Teste</title>
			<style>
				.tmdb-home-card {
					display: block;
					width: 100vw;
					height: 100vh;
					padding: 0;
					text-align: center;
					color: white;
					background-image: url(data:image/jpeg;base64,%s);
					background-position: 30% 45%;
					background-color: black;
					-webkit-background-size: co;
					-moz-background-size: cover;
					background-size: cover;
					-o-background-size: cover;
				 }
				 
				.row-tmdb-warning{
					display: flex;
					justify-content: center;
					align-items: center;
					flex-wrap: wrap;
				}
				
				.box-warning {
					border-radius:18px;
					background-color:#000000;
					width: 500px;
					height: 300px;
				}
				
			</style>
		</head>
		<body class="tmdb-home-card">
			<div class="box-warning">
			  <h1>😔 Não foi possível encontrar o resultado para sua solicitação</h1>
			  <h2>👉 Detalhes : %s</h2>
			</div>
		</body>
	</html>`

var SUCESS_HTTP_RESPONSE = `<!DOCTYPE html>
	<html lang="pt-br">
		<head>
			<meta charset="utf-8">
			<title>Tmdb Tutorial - Filmes Populares - Versão de Teste</title>
			<style>
				.tmdb-home-card {
						display: block;
						width: 100vw;
						height: 100vh;
						padding: 0;
						text-align: center;
						color: white;
						background-image: url(data:image/jpeg;base64,%s);
						background-position: %s;
						background-color: black;
						-webkit-background-size: co;
						-moz-background-size: cover;
						background-size: cover;
						-o-background-size: cover;
				 }
	
				.row-tmdb-flip-card {
					display: flex;
					justify-content: center;
					align-items: center;
					flex-wrap: wrap;
				}
				
				.card-flip-container {
					border-radius: 8px;
					margin-right: 5px;
					margin-top:5px;
					width: 225px;
					height: 400px;
					perspective: 1000px;
				}
	
				.flipper {
					position: relative;
					width: %s;
					height: %s;
					transition: transform 0.8s;
					transform-style: preserve-3d;
				}
	
				.card-flip-container:hover .flipper{
					transform: rotateY(180deg);
					cursor: pointer;
				}
	
				.front, .back{
					position: absolute;
					border-radius: 8px;
					width: %s;
					height: %s;
					backface-visibility: hidden;
				}
	
				.back{
					background-color: black;
					background-image: url(data:image/jpeg;base64,%s);
					transform: rotateY(180deg);
				}		
				
				.transparencia {
					filter: alpha(opacity=70);
					opacity: 0.7;
					-moz-opacity: 0.7;
					-webkit-opacity: 0.7;
				}
				
				.overwiew-status {
					border-radius:8px;
					margin:3px;
					height:140px; 
					width:220px;
					background-color:#000000;
				}
				
				.overwiew-data {
					overflow-y:auto;
					border-radius:8px;
					margin:3px;
					height:230px;
					width:220px;
					background-color:#000000
				}
				
				img {
					border-radius: 8px;
					width:220px;
					height:400px;
				}
			</style>
		</head>
		<body class="tmdb-home-card">
			<div class="row-tmdb-flip-card">%s</div>
		</body>
	</html>`

var HOME_CARD_FILM_POPULAR = `<div class="card-flip-container">
	<div class="flipper">
		<div class="front">
			<img src="%s" />
		</div>
		<div class="back">
			<div class="overwiew-status" style="filter:alpha(opacity=70);opacity:0.7;-moz-opacity:0.7;-webkit-opacity:0.7;">
				<div>%s</div>
				<div>
				   <span>Média de Votos: </span><span>%s</span>
				</div>
				<div>
				   <span>Votos Contabilizados: </span><span>%s</span>
				</div>
				<div>
				   <span>Data de Lançamento: </span><span>%s</span>
				</div>
			</div>
			<div class="overwiew-data" style="filter:alpha(opacity=70);opacity:0.7;-moz-opacity:0.7;-webkit-opacity:0.7;">
			   <p>%s</p>
			</div>
		</div>
	</div>
</div>`

var backgroundImageHomeBase64 string
var backgroundImageExceptionBase64 string
var backgroundImageCardBase64 string

//
type TmdbgBodyPopularFilms struct {
	RenderPage func(interface{}, int) (resultHtml string)
}

//
func NewTmdbgBodyPopularFilms() TmdbgBodyPopularFilms {
	return TmdbgBodyPopularFilms{
		RenderPage: renderPopularFilmsPage,
	}
}

//
func (ref TmdbgBodyPopularFilms) SetStreamB64Resources(strB64ImgHome, strB64ImgExcep, strB64ImgCard string) {
	backgroundImageHomeBase64 = strB64ImgHome
	backgroundImageExceptionBase64 = strB64ImgExcep
	backgroundImageCardBase64 = strB64ImgCard
}

//
func prepareCards(refStructFilmsPopular *tmdbgomodel.TmdbListPopularFilmModel) (panelCards string) {

	var list []tmdbgomodel.TmdbPopularFilmModel = refStructFilmsPopular.Results
	var viewCards string = ""

	for _, popularFilm := range list {
		imgSrc := popularFilm.Poster_Path
		title := popularFilm.Original_Title
		votesAverage := strconv.FormatFloat(popularFilm.Vote_Average, 'f', 0, 32)
		voteCount := strconv.Itoa(popularFilm.Vote_Count)
		dateRelease := popularFilm.Release_Date
		overwiew := popularFilm.Overview
		imgSrc = fmt.Sprintf("%s%s", tmdbgutil.VAL_PATH_URL_POSTER_IMAGE_TMDB, imgSrc)
		viewCards = fmt.Sprintf("%s%s", viewCards, fmt.Sprintf(HOME_CARD_FILM_POPULAR, imgSrc, title, votesAverage, voteCount, dateRelease, overwiew))
	}

	return viewCards
}

//
func renderPopularFilmsPage(data interface{}, statusCode int) (resultHtml string) {
	var dataResponse []byte = data.([]byte)
	var errorPage TmdbgBodyError = NewTmdbgBodyError()
	var formatErrorMsg string = "%s \n %s"
	var modelException *tmdbgomodel.TmdbgModelResponseError = &tmdbgomodel.TmdbgModelResponseError{}

	if statusCode == 200 {

		var responseStructFilmsPopular *tmdbgomodel.TmdbListPopularFilmModel = &tmdbgomodel.TmdbListPopularFilmModel{}
		iError := json.Unmarshal(dataResponse, responseStructFilmsPopular)

		if iError != nil {
			message := fmt.Sprintf(formatErrorMsg, tmdbgutil.MSG_ERROR_JSON_PROCESS, iError.Error())
			return errorPage.GetHtmlPageError(backgroundImageExceptionBase64, message)
		} else {
			messageRet := prepareCards(responseStructFilmsPopular)
			messageRet = fmt.Sprintf(SUCESS_HTTP_RESPONSE, backgroundImageHomeBase64, "30% 45%", "100%", "100%", "100%", "100%", backgroundImageCardBase64, messageRet)

			return messageRet
		}

	} else {

		iError := json.Unmarshal(dataResponse, modelException)

		if iError != nil {
			message := fmt.Sprintf(formatErrorMsg, tmdbgutil.MSG_ERROR_JSON_PROCESS, iError.Error())
			return errorPage.GetHtmlPageError(backgroundImageExceptionBase64, message)
		} else {
			messageRet := fmt.Sprintf(ERROR_HTTP_RESPONSE, backgroundImageHomeBase64, modelException.Status_Message)
			return messageRet
		}

	}

}
