# API REST 

API Restful com controle de cliente: 
* Implementa um CRUD (create, read,update,delete)
* Os registros são salvos em um banco de dados mongoDB. 
* Cada operação executada na API gera uma mensagem de LOG informativa,
  que é publicada em uma FILA RabbitMQ.
* Cada registro de um cliente tem as informações de NOME, CPF, NASCIMENTO, ENDEREÇO
    * Nome (Name):
        * Tipo: string
    * CPF (CPF):
        * Tipo: string
    * NASCIMENTO (Nascimento):
        * Tipo: string
    * ENDEREÇO (Endereco):
        * Tipo: string
        * Endereco é uma outra struct que contem: RUA, BAIRRO, CIDADE, ESTADO

Um segundo serviço, consome os eventos gerados em fila, e atualiza


## REQUISITO

Antes de começar, você precisará ter as seguintes ferramentas instaladas em sua máquina:
* [Git](https://git-scm.com)

Para rodar via docker
* [Docker](https://docs.docker.com/)

Para rodar Local
* [Go](https://golang.org/)
* [MongoDB](https://docs.mongodb.com/)
* [RabbitMQ](https://www.rabbitmq.com/)


## Iniciando Projeto 

iniciar config com docker-compose up

### Local
 

```bash
# Clone este repositório
$ git clone https://github.com/brunoOchoa.com/API-REST

# Instale as dependências e rode o projeto
$ go run main.go

# Server is running
```

### Docker

```bash
# Clone este repositório
$ git clone https://github.com/brunoOchoa.com/API-REST

# subir o docker compose 
$ docker-compose up 

```
#### Configuração RabbitMQ

```
* Acessar via navegador http://localhost:15672, 

```
 
## Rotas

```
http://localhost:8080/api/cliente/

```

| Rotas  |  HTTP Method  | Params  |  Descrição  | 
| :---: | :---: | :---: | :---: |
|  /cliente |  POST |  Body: ``` Name ```, ``` CPF ```, ``` Nascimento ```, ``` Endereco ``` |  Cadastra um novo cliente |
|  /cliente |  GET |  -  | Recupera todos os clientes |
|  /cliente/:id |  GET |  Params: ``` id ``` |  Consulta apenas um cliente pela id |
|  /cliente/:id |  PUT |   Params: ``` id ``` Body: ``` Name ```, ``` cpf ```, ``` Nascimento ```, ``` Endereco ``` |  Edite o cadastro de um cliente |
|  /cliente/:id |  DELETE |  Params: ``` id ``` |  Exclui o cadastro de um cliente|



### Requisições
* ``` POST /cliente ```

Corpo da requisição:
  
```
{
    "name": "cliente",
    "cpf": "111111111",
    "nascimento": "11/11/1111",
    "endereco": {
        "rua": "rua cliente",
        "bairro": "bairro",
        "cidade": "cidade",
        "estado": "estado"
    }
}

```

* ``` PUT /cliente ```

Corpo da requisição:
  
```
{
    "name": "cliente",
    "cpf": "111111111",
    "nascimento": "11/11/1111",
    "endereco": {
        "rua": "rua cliente",
        "bairro": "bairro",
        "cidade": "cidade",
        "estado": "estado"
    }
}

```
Parâmetro da requisição: 

```  http://localhost:8080/api/cliente/{id} ```



## Testes
Para executar os testes :

```bash

  # Rode os testes
  $ Rodar o main.go na raiz, depois rodar o worker pelo main.go na pasta worker
  
  # Postman
  $ fazer as requisiçoes pelo postman
  
```