package tmdbgoresources

import (
	"errors"
	"os"
	"tmdbGotutorial/tmdbgutil"
)

// Responsável por prover os recursos para a solução
type TmdbgResources struct {
	baseUrlTMDB string
	tokenTMDB   string
}

// NewTmdbgResources: Cria uma referência da estrutura em memória. Podemos dizer também, um objeto
func NewTmdbgResources() TmdbgResources {
	var strBaseUrlTMDB string = os.Getenv("BASE_URL_TMDB")
	var strTokenTMDB string = os.Getenv("TOKEN_TMDB")

	return TmdbgResources{
		baseUrlTMDB: strBaseUrlTMDB,
		tokenTMDB:   strTokenTMDB,
	}
}

// GetDataAccessTMDB : Retorna os dados necessários ao acesso da API, incluídos como variáveis de ambiente
func (ob TmdbgResources) GetDataAccessTMDB() (baseUrl string, token string, erro error) {

	refBaseUrl := ob.baseUrlTMDB
	refToken := ob.tokenTMDB

	if len(refBaseUrl) == 0 {
		return "", "", errors.New(tmdbgutil.MSG_ERROR_BASE_URL_API_NOT_FOUND)
	}

	if len(refToken) == 0 {
		return "", "", errors.New(tmdbgutil.MSG_ERROR_API_NOT_FOUND)
	}

	return ob.baseUrlTMDB, ob.tokenTMDB, nil
}
