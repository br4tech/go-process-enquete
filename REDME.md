# Sistema de Votação em Tempo Real com Go, MongoDB, PostgreSQL e RabbitMQ

Este projeto implementa um sistema de votação em tempo real utilizando as seguintes tecnologias:

- **Go**: Linguagem de programação principal para a API e componentes.
- **Echo**: Framework web para Go, utilizado para construir a API RESTful.
- **GORM**: ORM (Object Relational Mapper) para Go, facilitando a interação com o banco de dados PostgreSQL.
- **MongoDB**: Banco de dados NoSQL utilizado para armazenar os resultados parciais e em tempo real das enquetes.
- **PostgreSQL**: Banco de dados relacional utilizado para armazenar os dados das enquetes (perguntas, opções) e resultados finais.
- **RabbitMQ**: Message broker que permite a comunicação assíncrona entre os componentes do sistema, garantindo a atualização dos resultados em tempo real.

## Funcionalidades

- **Criar enquetes**: Permite criar novas enquetes com uma pergunta e múltiplas opções de resposta.
- **Votar**: Usuários podem votar em uma opção de uma enquete ativa.
- **Visualizar resultados**: Os resultados das enquetes são atualizados em tempo real e podem ser visualizados por meio da API.
- **Resultados consolidados**: Os resultados finais das enquetes são armazenados no PostgreSQL para posterior análise.

## Arquitetura

A arquitetura do sistema é composta pelos seguintes componentes:

- **API**: Responsável por receber as requisições dos usuários, processar os votos e publicar os eventos de votação no RabbitMQ.
- **RabbitMQ**: Message broker que recebe os eventos de votação e os distribui para os consumidores.
- **Consumidor de votos**: Componente responsável por consumir os eventos de votação, atualizar os resultados no MongoDB e no PostgreSQL.
- **MongoDB**: Armazena os resultados parciais e em tempo real das enquetes.
- **PostgreSQL**: Armazena os dados das enquetes e os resultados finais.

## Como executar

### Pré-requisitos:

- Go (versão 1.21)
- Docker e Docker Compose

### Clonar o repositório:

```bash
  git clone  git@github.com:br4tech/goi-process-enquete.git
```

### Iniciar os serviços:

```bash
  docker-compose up -d
```

### Executar a aplicação:

```bash
  go run main.go
```

# Endpoints da API

- **POST /poll**: Cria uma nova enquete.
- **GET /poll/:id**: Obtém os resultados de uma enquete.
- **POST /vote/:id/:option**: Vota em uma opção de uma enquete.
