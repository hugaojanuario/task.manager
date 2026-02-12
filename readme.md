# Task Manager API

Projeto de API REST para gerenciamento de tarefas desenvolvido em Go.  
A aplicação permite criar, listar, atualizar e remover tarefas por meio de endpoints HTTP, utilizando persistência em banco de dados e execução via Docker Compose.

---

## Visão Geral

Este projeto implementa uma API REST para gerenciamento de tarefas com foco em:

- Estrutura modular
- Separação de responsabilidades
- Organização por camadas
- Execução simplificada com Docker

A aplicação foi desenvolvida com o objetivo de consolidar fundamentos de desenvolvimento backend utilizando Go.

---

## Funcionalidades

- CRUD completo de tarefas:
  - Criar nova tarefa
  - Listar todas as tarefas
  - Buscar tarefa por ID
  - Atualizar tarefa existente
  - Remover tarefa
- Estrutura organizada por handlers, rotas e modelos
- Testes automatizados
- Configuração via Docker e Docker Compose

---

## Tecnologias Utilizadas

- Go
- Gin
- Gorm
- PostgreSQL
- Docker Compose

---

## Estrutura do Projeto

```
task.manager/
│
├── cmd/api                # Ponto de entrada da aplicação
├── internal/handler       # Controllers e lógica de requisições HTTP
├── http/routes            # Definição das rotas
├── model                  # Estruturas e modelos
├── database               # Configuração e conexão com banco
├── test                   # Testes automatizados
├── docker-compose.yml     # Orquestração dos serviços
└── go.mod                 # Gerenciamento de dependências
```

---

## Pré-requisitos

Para executar o projeto localmente:

- Go instalado (1.25.1)
- Docker Compose

---

## Como Executar

1. Clone o repositório:

```bash
git clone https://github.com/hugaojanuario/task.manager.git
cd task.manager
```

2. Suba os serviços com Docker Compose:

```bash
docker compose up --build
```

3. A aplicação estará disponível em:

```
http://localhost:8080
```

(Confirme a porta no arquivo `docker-compose.yml`.)

---

## Endpoints

### Listar todas as tarefas
```
GET /task
```

### Buscar tarefa por ID
```
GET /task/{id}
```

### Criar nova tarefa
```
POST /task
Content-Type: application/json
```

Exemplo de corpo:

```json
{
  "title": "Estudar Go",
  "description": "Revisar conceitos de API REST"
}
```

### Atualizar tarefa
```
PUT /task/{id}
```

### Remover tarefa
```
DELETE /task/{id}
```

---

## Exemplos de Uso

Criar tarefa:

```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Exemplo","description":"Descrição"}'
```

Listar tarefas:

```bash
curl http://localhost:8080/tasks
```

---

## Testes

Para rodar os testes:

```bash
go test ./...
```

---

## Objetivo do Projeto

Este projeto foi desenvolvido como prática de backend com Go, aplicando conceitos de:

- API REST
- Estrutura modular
- Organização de código
- Integração com banco de dados
- Ambiente containerizado
- Testes automatizados

---

## Licença

Projeto para fins educacionais e de estudo.
