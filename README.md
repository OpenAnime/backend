<div align="center">
  
# Constani Go Backend 

<a href="https://www.codefactor.io/repository/github/openanime/backend"><img src="https://www.codefactor.io/repository/github/openanime/backend/badge" alt="CodeFactor" /></a>
  <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/Constani/superapi">

**Constani GoLang Backend to Manage Database Services and API Endpoints**

</div>

# Installation and running

Clone the repository and run `go run main.go` in your terminal to open the API server.

# Note 

Default port number is **3001**

# Endpoints

**GET** `/api/anime` Gives a response as all of the stored anime data

**GET** `/api/anime/home` Gives a avatar, banner, name, id and info all anime.

**GET** `/api/anime/get/:id` Gives a response as the data of the specified anime ID's anime serie

**GET** `/api/anime/get/:id/home` Gives a avatar, banner, name, id and info.

**POST** `/api/anime/create` requires a JSON body with keys that are given below

| Keys      | Type             | Required |
|-----------|------------------|----------|
| Id        | string           | true     |
| name      | string           | true     |
| avatar    | string           | true     |
| totalLike | int              | true     |
| episodes  | array of objects | true     |


