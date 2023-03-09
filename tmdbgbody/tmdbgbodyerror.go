package tmdbgbody

import "fmt"

//
type TmdbgBodyError struct {
	bodyTextMessage string
}

//
func NewTmdbgBodyError() TmdbgBodyError {
	return TmdbgBodyError{}
}

//
func (refStruct TmdbgBodyError) GetHtmlPageError(dataImageStr, msgError string) string {
	refStruct.bodyTextMessage = iHtml(dataImageStr, msgError)
	return refStruct.bodyTextMessage
}

//
func iHtml(dataImageStrB64, message string) string {
	var iHtmlStr string = `<!DOCTYPE html>
	<html lang="pt-br">
		<head>
			<meta charset="utf-8">
			<title>Erro Tutorial - Golang consumindo uma API do TMDB em uma Infra AWS</title>
			<style>
				.intro {
						display: block;
						width: 100vw;
						height: 100vh;
						padding: 0;
						text-align: center;
						color: white;
						background: url(data:image/png;base64,%s) no-repeat bottom center scroll;
						background-position: 30% 45%;
						background-color: black;
						-webkit-background-size: co;
						-moz-background-size: cover;
						background-size: cover;
						-o-background-size: cover;
				}
	
				p {
					font-size: large;
					font: bold;
				}
			</style>    
		</head>    
		<body class="intro">
			<div>
				<h1>Aviso</h1>
				<p>
				 Ocorreu um erro no Acesso ao Sistema!!
				</p>
				<p>
					Mensagem: %s
				</p>    
			</div>    
		</body>    
	</html>    `

	iHtmlStr = fmt.Sprintf(iHtmlStr, dataImageStrB64, message)
	return iHtmlStr
}
