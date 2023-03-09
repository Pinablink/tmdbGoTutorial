package main

import (
	"tmdbGotutorial/tmdbgS3"
	"tmdbGotutorial/tmdbgapi"
	"tmdbGotutorial/tmdbgbody"
	"tmdbGotutorial/tmdbgoresources"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Recursos Api e Http
var pagePopularFilms tmdbgbody.TmdbgBodyPopularFilms
var httpApi tmdbgapi.TmdbgApiHttp

// Objetos de imagem no Bucket em Base64
var dataObjStreamBackgroundHome string
var dataObjStreamBackgroundException string
var dataObjStreamBackgroundCard string

// Interface com sua Implementação para o AWS entender oque precisa ser feito.
// O AWS é um cara legal e vai seguir exatamente oque você pedir aqui!
func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Carregamento das funcionalidades básicas
	var apiResponse events.APIGatewayProxyResponse
	var errorPage tmdbgbody.TmdbgBodyError = tmdbgbody.NewTmdbgBodyError()

	// Para configurar o Header do Response
	var mapHeader map[string]string
	mapHeader = make(map[string]string)
	mapHeader["Content-Type"] = "text/html"

	// Carrega os objetos disponíveis no Bucket
	existErrorLoadS3, msgErrorLoadS3Str := loadDataBucketS3()

	if existErrorLoadS3 {

		apiResponse = events.APIGatewayProxyResponse{
			Body:       errorPage.GetHtmlPageError("", msgErrorLoadS3Str),
			StatusCode: 200,
			Headers:    mapHeader,
		}

	} else {

		// Passando os Streams das Imagens existentes no bucket
		pagePopularFilms.SetStreamB64Resources(dataObjStreamBackgroundHome, dataObjStreamBackgroundException, dataObjStreamBackgroundCard)

		// Obtêm os recursos necessários para a execução da Função
		var resources tmdbgoresources.TmdbgResources = tmdbgoresources.NewTmdbgResources()

		strBaseUrl, strKeyToken, strToken, resError := resources.GetDataAccessTMDB()

		if resError != nil {

			apiResponse = events.APIGatewayProxyResponse{
				Body:       errorPage.GetHtmlPageError(dataObjStreamBackgroundException, resError.Error()),
				StatusCode: 200,
				Headers:    mapHeader,
			}

		} else {

			httpApi = tmdbgapi.NewTmdbgApiHttp(strKeyToken, strToken, strBaseUrl, pagePopularFilms.RenderPage)

			htmlResponse, errResponse := httpApi.GetPopularFilms()

			if errResponse != nil {
				htmlResponse = errorPage.GetHtmlPageError(dataObjStreamBackgroundException, errResponse.Error())
			}

			apiResponse = events.APIGatewayProxyResponse{
				Body:       htmlResponse,
				StatusCode: 200,
				Headers:    mapHeader,
			}
		}

	}

	return apiResponse, nil
}

// Carrega para a Função todos os objetos existentes no Bucket
func getObjBucket(nameObj string) (errorLoadBucket bool, messageErrorLoadBucket string, resultBucket string) {

	var msgErrorLoadBucket, strResultData string
	//mBucket, errCreate := tmdbgS3.NewTmdbgBucket("tmdb-tutorial", "my_img_tutor_tmdb.dat")
	mBucket, errCreate := tmdbgS3.NewTmdbgBucket("tmdb-tutorial", nameObj)

	if errCreate != nil {
		// Tratamento de proteção, caso ocorra erro de carregamento do bucket. Entrar aqui
		// significa que você deve revisar o seu bucket S3
		msgErrorLoadBucket = "Erro no Carregamento do Bucket"
		return true, msgErrorLoadBucket, ""
	}

	strResult, s3Err := mBucket.GetErrorMessage()

	if s3Err != nil {
		// Tratamento de proteção, caso ocorra erro de carregamento do bucket. Entrar aqui
		// significa que você deve revisar o seu bucket S3
		msgErrorLoadBucket = "Erro no processamento do S3"
		return true, msgErrorLoadBucket, ""
	}

	strResultData = strResult

	return false, "", strResultData
}

// Carrega todos os objetos existentes no Bucket
func loadDataBucketS3() (existError bool, messageError string) {

	errorLoadBucketBE, msgErrorLoadBE, resultBE := getObjBucket("my_img_tutor_tmdb_exception.dat")

	if errorLoadBucketBE {
		return true, msgErrorLoadBE
	}

	errorLoadBH, msgErrorLoadBH, resultBH := getObjBucket("my_img_tutor_tmdb_home.dat")

	if errorLoadBH {
		return true, msgErrorLoadBH
	}

	errorLoadBC, msgErrorLoadBC, resultBC := getObjBucket("my_img_tutor_tmdb_back_card.dat")

	if errorLoadBC {
		return true, msgErrorLoadBC
	}

	dataObjStreamBackgroundHome = resultBH
	dataObjStreamBackgroundException = resultBE
	dataObjStreamBackgroundCard = resultBC

	return false, ""
}

// Função principal da Aplicação
func main() {
	pagePopularFilms = tmdbgbody.NewTmdbgBodyPopularFilms()
	lambda.Start(HandleRequest)
}
