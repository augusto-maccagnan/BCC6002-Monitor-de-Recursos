# Monitor de Recursos
> **Disciplina:** BCC_6002 - Aspectos de Linguagens de Programação  
> **Objetivo:** Explorar os aspectos e funcionalidades oferecidas pela linguagem de programação Go (Golang) ao desenvolver um monitor de recursos.

Este projeto tem como objetivo implementar um sistema de monitoramento de recursos, utilizando o Golang para o backend e tecnologias web modernas para o frontend.

### <span style="color:red">***O projeto foi desenvolvido apenas para Linux***</span>
---

## Estrutura do Projeto

O projeto é composto por duas partes principais:
- **Backend:** Responsável pelo processamento e fornecimento de dados de monitoramento.
- **Frontend:** Interface de usuário para exibir os dados de monitoramento de forma interativa.
---

## Como Rodar o Projeto

### Requisitos
Antes de começar, certifique-se de que possui as ferramentas necessárias instaladas em seu ambiente:
- [Go](https://golang.org/doc/install) (para o backend)
- [Node.js](https://nodejs.org/en/) (para o frontend)
### Backend

```
cd /backend
```

Para subir o servidor backend é necessário intalar o pacote de variáveis de ambiente [github.com/joho/godotenv](https://github.com/joho/godotenv.git):

```
go get github.com/joho/godotenv
```

você pode optar por duas formas de inicialização:

#### Usando Makefile (recomendado):

```
make
```

#### Usando o Golang diretamente
Caso prefira rodar diretamente com o Golang, siga os passos abaixo:

```
go run main.go
```

Isso iniciará o servidor backend.

### Frontend

```
cd /frontend
```

O frontend é desenvolvido com tecnologias baseadas em JavaScript, então você precisará instalar as dependências via npm:

#### Instalar dependências

```
npm install
```

#### Iniciar o servidor frontend

```
npm start
```

Isso irá iniciar o servidor do frontend e você poderá acessá-lo no seu navegador.

---

## Tecnologias Utilizadas

- **Backend:** Go (Golang)
- **Frontend:** Node.js
