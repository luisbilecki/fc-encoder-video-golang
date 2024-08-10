# FC Encoder GoLang Video

## Como rodar?

1. Inicializar

`docker-compose up`

2. Criar chaves e projeto no Google com Cloud Storage e colocar as credenciais no arquivo `bucket-credential.json`

3. Inicializar aplicação acessando container:
`docker exec -it fc-encoder-video-golang-app-1 sh`
e 
`go run framework/cmd/server/server.go`

4. Quando iniciar acesse o RabbitMQ em `localhost:15672` com rabbitmq de usuário e senha igual.

5. Abra as `queues` e poste na queue de `video` a mensagem:
```
{"resource_id": "id-cliente-123", "file_path":"video.mp4"}
```
*Alterar o resource_id e file_path para o vídeo que está no storage*