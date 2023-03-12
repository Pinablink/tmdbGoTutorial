# tmdbGoTutorial

![Alt text](img/tmdbGoTutorialResult.gif)

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Pinablink/tmdbGoTutorial?style=plastic)

**Para mais detalhes leia também https://medium.com/@weberasantos/golang-consumindo-api-tmdb-em-uma-infra-aws-4a8c540a93d6**

Este aplicativo foi desenvolvido para ser um **material básico de consulta**. Que possa proporcionar a você visitante, uma referência ao uso da **linguagem Golang** na **Infra AWS**. Com esse recurso, você será capaz de obter um conhecimento básico do consumo de uma Aplicação escrita em Golang, aos dados da **API TMDB** www.themoviedb.org.
</br>

### 👀 Você deve se perguntar. O que esse conteúdo vai me proporcionar?

👉 **Os seguintes conhecimentos estão sendo abordados aqui nesse projeto:** 

* Desenvolvimento de uma **Função Lambda AWS** 😉
* Get de um objeto existente em um **Bucket S3** 😉
* Configuração e chamada de uma **API Externa** 😉

</br>

⚠ **Importante** ⚠

Antes de fazer o download desse projeto: 

* Verifique o arquivo **go.mod** para conhecer suas dependências. 
* Efetue o cadastro de usuário no www.themoviedb.org. Para que desse modo, você tenha acesso a API e outros recursos interessantes. 
* Obtenha uma conta na **AWS**. 
* Leia algum conteúdo sobre **Função Lambda AWS** em **Golang**. 
* Leia algum material sobre o uso de uma aplicação **Golang** no Gerenciamento de um **Bucket S3**
</br>

😏 **Curiosidade**
</br>
Existe uma aplicação em Android que desenvolvi no passado. Que também utiliza a API TMDB. Se ficou curioso, acesse:
https://github.com/Pinablink/FilmesFamosos

## Sobre a estrutura desse Projeto

O projeto desse aplicativo está organizado nos seguintes pacotes. 

* 📁 assets
</br> **Imagens codificadas em Base64 para composição da página de resultado** 
</br>

* 📁 tmdbgapi
</br> **Atende as configurações e as chamadas da API**
</br>

* 📁 tmdbgbody
</br> **Organiza os dados para renderização e guarda a estrutura HTML da tela**
</br>

* 📁 tmdbgomodel
</br> **Representa do Modelo de Dados que a API irá disponibilizar, referente a Lista de Filmes Populares**
</br>

* 📁 tmdbgoresources
</br> **Obtêm todos os recursos necessário para a execução da aplicação**
</br>

* 📁 tmdbgS3
</br> **Possui as funções necessárias ao acesso do Bucket S3 da aplicação**
</br>

* 📁 tmdbgutil
</br> **Pacote para disponibilizar funcionalidades comuns a aplicação. Contem uma estrutura com a definição de constantes que atende toda a aplicação**

</br>

👍 **Obrigado pela sua visita!** 👍