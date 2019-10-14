# API Rest in Golang

Uma simples implementação de API Rest usando o `Muxer Pat`.

## Estrutura do projeto

### Resource PING

Primeiro endpoint __GET__:`/api/ping`

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
## Referencias

[Manage Config - Felipe Dutra Tine e Silva](https://medium.com/@felipedutratine/manage-config-in-golang-to-get-variables-from-file-and-env-variables-33d876887152)