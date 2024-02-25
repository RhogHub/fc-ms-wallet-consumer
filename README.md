# Microsservices with Go and Event Driven Architecture (EDA) - FC Challenge

This repository contains two microservices: **walletcore** and **walletconsumer**.

## Usage

To run the microservices and associated services, follow these instructions:

1. Ensure you have Docker installed on your system.

2. Run the following command to start the services:

   ```bash
   docker-compose up -d

This command will start the microservices along with Kafka and other required services.

## Informações dos Serviços

- **Kafka**: O Kafka está sendo executado na porta 9021.
- **walletcore**: Acesse o microsserviço walletcore na porta 8080.
- **walletconsumer**: Acesse o microsserviço walletconsumer na porta 3003.

## Testes

- Para cada microsserviço:
  1. Navegue até o diretório do microsserviço (`walletcore` ou `walletconsumer`).
  2. Execute os testes unitários executando o seguinte comando:

     ```bash
     go test ./...
     ```

  3. Para testes de API, utilize os arquivos `api.http` e `client.http` localizados na pasta `api` de cada microsserviço.

---

Autor: Rodrigo Godoi
Email: rhog.dev@gmail.com