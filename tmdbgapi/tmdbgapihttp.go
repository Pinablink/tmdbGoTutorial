package tmdbgapi

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"tmdbGotutorial/tmdbgutil"
)

//
type TmdbgApiHttp struct {
	apiTokentmdb string
	urlapitmdb   string
	rendererPage func(interface{}, int) string
}

// Criar uma estrutura de comunicação com a API TMDB.
func NewTmdbgApiHttp(strTokenTmdb, strUrlApiTmdb string, renderer func(interface{}, int) string) TmdbgApiHttp {
	return TmdbgApiHttp{
		apiTokentmdb: strTokenTmdb,
		urlapitmdb:   strUrlApiTmdb,
		rendererPage: renderer,
	}
}

// Cria a Autorização para adicionar ao Cabeçalho da Requisição da API
func (ref TmdbgApiHttp) createAuthorization() (myAuthorization string) {
	var STR_FORMAT_AUTHORIZATION string = "%s %s"
	var STR_AUTHORIZATION string = fmt.Sprintf(STR_FORMAT_AUTHORIZATION, tmdbgutil.VAL_AUTHORIZATION_TMDB, ref.apiTokentmdb)

	return STR_AUTHORIZATION
}

// Obter os Filmes Populares da API
func (ref TmdbgApiHttp) GetPopularFilms() (strHtml string, getPopularFilmsError error) {

	var streamBody []byte
	var streamReadError error
	var FORMAT_ERROR_MESSAGE string = "%s \n %s"
	var FORMAT_URL_OPT_POPULAR string = "%s%s"

	ref.urlapitmdb = fmt.Sprintf(FORMAT_URL_OPT_POPULAR, ref.urlapitmdb, tmdbgutil.VAL_POPULAR_FILMS_CATEG_TMDB)
	fmt.Println(ref.urlapitmdb)
	request, iError := http.NewRequest("GET", ref.urlapitmdb, nil)

	if iError != nil {
		return "", errors.New(tmdbgutil.MSG_ERROR_CREATE_REQUEST_TMDB_API)
	}

	// Preparando a Requisição para a API
	var strAuthorization string = ref.createAuthorization()
	request.Header.Add(tmdbgutil.ATTR_AUTHORIZATION_TMDB, strAuthorization)

	client := &http.Client{}
	response, response_error := client.Do(request)

	if response_error != nil {
		var msgError string = fmt.Sprintf(FORMAT_ERROR_MESSAGE, tmdbgutil.MSG_ERROR_RESPONSE_TMDB_API, response_error.Error())
		return "", errors.New(msgError)
	}

	defer response.Body.Close()
	streamBody, streamReadError = ioutil.ReadAll(response.Body)

	if streamReadError != nil {
		var msgError string = fmt.Sprintf(FORMAT_ERROR_MESSAGE, tmdbgutil.MSG_ERROR_STREAM_READ, streamReadError.Error())
		return "", errors.New(msgError)
	}

	strResponseHtml := ref.rendererPage(streamBody, response.StatusCode)

	return strResponseHtml, nil
}
