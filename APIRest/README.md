# API Rest in Golang

Uma simples implementação de API Rest usando o `Muxer Pat`.

## Estrutura do projeto

* `Models`: Esta camada armazena todos os modelos (`struct`) e usado por todas as outras camadas.
* `Repository`: Respository esta responsavel pelo trabalho de relacionar com a database como querying, inserting/sorting ou deleting.
* 

__TODO__:
* `DB Connection SQLite and MySQL` 
* `HANDLER ERRORS`
* ``
 
### Resource PING

Endpoint teste.
__GET__:`/api/ping`

Um MiddleWare para interceptar as requisições recebidas para fins de DEBUG é LOG.

Uma simples resposta de pong com codigo de status da requisição.

Requisição:
```bash
$ curl -X GET -H "Content-type: application/json" localhost:8000/api/ping
```
Response:
```python
{
  "Status": 200,
  "Message": "pong"
}
```

### Resource Registry

Endpoint para cadastro de usuàrios
__POST__:`/api/registry`

Requisição:
```bash
$ curl -X POST -H "Content-Type: application/json" \
	-d '{"Name":"Andre","Username":"andre.ferreira","Email":"andre.ferreira@gmail.com","Password":"asdf12j3h4k1j"}' \
	http://localhost:8000/api/registry
```

Resposta:
```bash
{
  "name": "Andre",
  "username": "andre.ferreira",
  "email": "andre.ferreira@gmail.com",
  "active": false,
  "createat": "2019-10-14T13:21:29.127259892-03:00",
  "updateat": "0001-01-01T00:00:00Z"
}
```

### Handling Errors

Quando estamos olhando para o fluxo entre a geração do erro, manuseio do erro e analise de possiveis falahas.

__O que é um error in GO__
Olhando para o tipo `builtin error` nós podemos tomar algumas conclusões:

`O erro e tipo built-in interface e uma interface convencional`
```
type error interface {
  Error() string
}
```

Um error e uma `interface` que implementa um simples metódo `Error` retornando uma string.



## Referencias
[Handling Errors - Stupid Gopher](https://medium.com/hackernoon/golang-handling-errors-gracefully-8e27f1db729f)
[Manage Config - Felipe Dutra Tine e Silva](https://medium.com/@felipedutratine/manage-config-in-golang-to-get-variables-from-file-and-env-variables-33d876887152)