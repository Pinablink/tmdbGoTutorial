package tmdbgoresources

import (
	"errors"
	"os"
)

// Responsável por prover os recursos para a solução
type TmdbgResources struct {
	baseUrlTMDB  string
	keyTokenTMDB string
	tokenTMDB    string
}

// NewTmdbgResources: Cria uma referência da estrutura em memória. Podemos dizer também, um objeto
func NewTmdbgResources() TmdbgResources {
	var strBaseUrlTMDB string = os.Getenv("BASE_URL_TMDB")
	var strKeyTokenTMDB string = os.Getenv("KEY_TOKEN_TMDB")
	var strTokenTMDB string = os.Getenv("TOKEN_TMDB")

	return TmdbgResources{
		baseUrlTMDB:  strBaseUrlTMDB,
		keyTokenTMDB: strKeyTokenTMDB,
		tokenTMDB:    strTokenTMDB,
	}
}

// GetDataAccessTMDB : Retorna os dados necessários ao acesso da API, incluídos como variáveis de ambiente
func (ob TmdbgResources) GetDataAccessTMDB() (baseUrl string, keyToken string, token string, erro error) {

	refBaseUrl := ob.baseUrlTMDB
	refKeyToken := ob.keyTokenTMDB
	refToken := ob.tokenTMDB

	if len(refBaseUrl) == 0 {
		return "", "", "", errors.New("Não foi encontrado Base URL da API")
	}

	if len(refKeyToken) == 0 {
		return "", "", "", errors.New("Não foi encontrado a chave token da API")
	}

	if len(refToken) == 0 {
		return "", "", "", errors.New("Não foi encontrado o token da API")
	}

	return ob.baseUrlTMDB, ob.keyTokenTMDB, ob.tokenTMDB, nil
}
