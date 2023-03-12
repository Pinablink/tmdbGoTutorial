package tmdbgutil

const (
	MSG_ERROR_CREATE_REQUEST_TMDB_API      string = "Ocorreu um erro na criação da requisição para a API TMDB"
	MSG_ERROR_RESPONSE_TMDB_API                   = "Ocorreu um erro na resposta da API TMDB"
	MSG_ERROR_STREAM_READ                         = "Ocorreu um erro na obtenção dos dados no resultado da API TMDB"
	MSG_ERROR_JSON_PROCESS                        = "Ocorreu um erro no processamento do resultado da API"
	MSG_ERROR_BASE_URL_API_NOT_FOUND              = "Não foi encontrado Base URL da API"
	MSG_ERROR_API_NOT_FOUND                       = "Não foi encontrado o token da API"
	MSG_ERROR_LOAD_RESOURCE_BUCKET_S3             = "Erro ao Carregar recurso do Bucket S3"
	MSG_ERROR_READ_STRUCT_OBJ_BUCKET_S3           = "Erro ao ler estrutura do objeto S3"
	MSG_ERROR_LOAD_BUCKET_S3                      = "Erro no Carregamento do Bucket"
	MSG_ERROR_PROCESS_BUCKET_S3                   = "Erro no processamento do S3"
	ATTR_AUTHORIZATION_TMDB                       = "Authorization"
	ATTR_BUCKET_S3_NAME                           = "tmdb-tutorial"
	ATTR_API_KEY_TMDB                             = "api_key"
	ATTR_NAME_IMG_TUTOR_BACK_EXCEPTION_DAT        = "my_img_tutor_tmdb_exception.dat"
	ATTR_NAME_IMG_TUTOR_HOME_DAT                  = "my_img_tutor_tmdb_home.dat"
	ATTR_NAME_IMG_TUTOR_BACK_CARD                 = "my_img_tutor_tmdb_back_card.dat"
	VAL_AUTHORIZATION_TMDB                        = "Bearer "
	VAL_POPULAR_FILMS_CATEG_TMDB                  = "popular"
	VAL_PATH_URL_POSTER_IMAGE_TMDB                = "https://image.tmdb.org/t/p/original"
)
