package main

import (
	"tmdbGotutorial/tmdbgS3"
	"tmdbGotutorial/tmdbgapi"
	"tmdbGotutorial/tmdbgbody"
	"tmdbGotutorial/tmdbgoresources"
	"tmdbGotutorial/tmdbgutil"

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

		strBaseUrl, strToken, resError := resources.GetDataAccessTMDB()

		if resError != nil {

			apiResponse = events.APIGatewayProxyResponse{
				Body:       errorPage.GetHtmlPageError(dataObjStreamBackgroundException, resError.Error()),
				StatusCode: 200,
				Headers:    mapHeader,
			}

		} else {

			httpApi = tmdbgapi.NewTmdbgApiHttp(strToken, strBaseUrl, pagePopularFilms.RenderPage)

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
	mBucket, errCreate := tmdbgS3.NewTmdbgBucket(tmdbgutil.ATTR_BUCKET_S3_NAME, nameObj)

	if errCreate != nil {
		// Tratamento de proteção, caso ocorra erro de carregamento do bucket. Entrar aqui
		// significa que você deve revisar o seu bucket S3
		msgErrorLoadBucket = tmdbgutil.MSG_ERROR_LOAD_BUCKET_S3
		return true, msgErrorLoadBucket, ""
	}

	strResult, s3Err := mBucket.GetErrorMessage()

	if s3Err != nil {
		// Tratamento de proteção, caso ocorra erro de carregamento do bucket. Entrar aqui
		// significa que você deve revisar o seu bucket S3
		msgErrorLoadBucket = tmdbgutil.MSG_ERROR_PROCESS_BUCKET_S3
		return true, msgErrorLoadBucket, ""
	}

	strResultData = strResult

	return false, "", strResultData
}

// Carrega todos os objetos existentes no Bucket
func loadDataBucketS3() (existError bool, messageError string) {

	errorLoadBucketBE, msgErrorLoadBE, resultBE := getObjBucket(tmdbgutil.ATTR_NAME_IMG_TUTOR_BACK_EXCEPTION_DAT)

	if errorLoadBucketBE {
		return true, msgErrorLoadBE
	}

	errorLoadBH, msgErrorLoadBH, resultBH := getObjBucket(tmdbgutil.ATTR_NAME_IMG_TUTOR_HOME_DAT)

	if errorLoadBH {
		return true, msgErrorLoadBH
	}

	errorLoadBC, msgErrorLoadBC, resultBC := getObjBucket(tmdbgutil.ATTR_NAME_IMG_TUTOR_BACK_CARD)

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
