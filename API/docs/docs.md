## AFEB API

Esta é a API oficial do website do CCX. Aqui você encontra todos os endpoints
do sistema para colaborar no desenvolvimento deste projeto, ou então conectar-se
a ele para a criação de um projeto seu.

Link da API: <a href="https://afeb-api.onrender.com/">https://afeb-api.onrender.com/</a>

# Endpoints

* [/api/jogadores](#apijogadores)
* [/api/noticias](#apinoticias)
* [/api/torneios](#apitorneios)
* [/api/trofeus](#apitrofeus)
* [/api/usuarios](#apiusuarios)

## /api/jogadores

<details>
  <summary> <code>GET</code> <code>/api/jogadores/ranking</code> </summary>

  ### Descrição

  Retorna um ranking com todos os jogadores registrados na AFEB, em ordem
  decrescente de rating rápido.

  ### Parâmetros

  > Nenhum

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Response
  ```js
  {
    "ranking": []object,
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Ranking encontrado com sucesso!",
    "ranking": [
      {
        "codJog": 1,
        "nome": "Lucas Guedes",
        "apelido": "Guedes",
        "tituloAFEB": "MNB",
        "info": "Fundador do CCX.",
        "eloRapid": 3200,
        "eloBlitz": 3400,
        "jogos": 405,
        "vitorias": 300,
        "derrotas": 100,
        "empates": 5,
        "dataNascimento": "2005-09-02",
        "trofeus": null
      },
      {
        "codJog": 2,
        "nome": "Bye",
        "apelido": null,
        "tituloAFEB": "GMB",
        "info": "O bye",
        "eloRapid": 9999,
        "eloBlitz": 9999,
        "jogos": 3000,
        "vitorias": 3000,
        "derrotas": 0,
        "empates": 0,
        "dataNascimento": "0000-01-01",
        "trofeus": null
      }
    ]
  }
  ```
</details>

<details>
  <summary> <code>GET</code> <code>/api/jogadores/${codJog}</code> </summary>

  ### Descrição

  Retorna os dados de um jogador, juntamente de suas premiações na AFEB.

  ### Parâmetros

  | param.    |  tipo     | data type   | desc.                                                            |
  |-----------|-----------|-------------|------------------------------------------------------------------|
  | codJog    |  required | string      | Código do jogador a ser buscado.                                 |

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Response
  ```js
  {
    "jogador": object,
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "jogador": {
      "codJog": 2,
      "nome": "Bye",
      "apelido": null,
      "tituloAFEB": "GMB",
      "info": "O bye",
      "eloRapid": 9999,
      "eloBlitz": 9999,
      "jogos": 405,
      "vitorias": 300,
      "derrotas": 100,
      "empates": 5,
      "dataNascimento": "0000-01-01",
      "trofeus": [
        {
          "codTrof": 0,
          "codJog": 0,
          "codTorn": 0,
          "torneio": "1º Campeonato Mundial da FIDE",
          "posicao": 1
        }
      ]
    },
    "message": "Jogador encontrado com sucesso!"
  }
  ```
</details>

<details>
  <summary> <code>POST</code> <code>/api/jogadores</code> </summary>

  ### Descrição

  Registra um novo jogador no sistema

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 201 | `CREATED` |
  | 400 | `BAD REQUEST` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Request body
  ```js
  {
    "nome": string,
    "apelido": string || null,
    "tituloAFEB": string || null,
    "info": string || null,
    "eloRapid": Number || null,
    "eloBlitz": Number || null,
    "vitorias": Number,
    "derrotas": Number,
    "empates": Number,
    "dataNascimento": string
  }
  ```

  Exemplo:
  ```js
  {
    "nome": "Bye",
    "apelido": "Ciao",
    "tituloAFEB": "GMB",
    "info": "O bye.",
    "eloRapid": 9999,
    "eloBlitz": 9999,
    "vitorias": 3000,
    "derrotas": 0,
    "empates": 0,
    "dataNascimento": "1287-01-01"
  }
  ```

  ### Response
  ```js
  {
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Jogador registrado com sucesso!"
  }
  ```
</details>

<details>
  <summary> <code>PUT</code> <code>/api/jogadores</code> </summary>

  ### Descrição

  Atualiza os dados de um jogador no sistema

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 400 | `BAD REQUEST` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Request body
  ```js
  {
    "codJog": Number,
    "nome": string,
    "apelido": string || null,
    "tituloAFEB": string || null,
    "info": string || null,
    "eloRapid": Number || null,
    "eloBlitz": Number || null,
    "vitorias": Number,
    "derrotas": Number,
    "empates": Number,
    "dataNascimento": string
  }
  ```

  Exemplo:
  ```js
  {
    "codJog": 0,
    "nome": "Bye",
    "apelido": "Ciao",
    "tituloAFEB": "GMB",
    "info": "O bye.",
    "eloRapid": 9999,
    "eloBlitz": 9999,
    "vitorias": 3000,
    "derrotas": 0,
    "empates": 0,
    "dataNascimento": "1287-01-01"
  }
  ```

  ### Response
  ```js
  {
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Jogador atualizado com sucesso!"
  }
  ```
</details>

<details>
  <summary> <code>DELETE</code> <code>/api/jogadores/${codJog}</code> </summary>

  ### Descrição

  Exclui um jogador permanentemente do sistema.

  ### Parâmetros

  | param.    |  tipo     | data type   | desc.                                                            |
  |-----------|-----------|-------------|------------------------------------------------------------------|
  | codJog    |  required | string      | Código do jogador a ser excluído.                                |

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 400 | `BAD REQUEST` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Response
  ```js
  {
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Jogador excluído com sucesso!"
  }
  ```
</details>

## /api/noticias

<details>
  <summary> <code>GET</code> <code>/api/noticias</code> </summary>

  ### Descrição

  Retorna todas as notícias do CCX.

  > Em desenvolvimento
</details>

<details>
  <summary> <code>GET</code> <code>/api/noticias/feed</code> </summary>

  ### Descrição

  Retorna as 6 últimas notícias do CCX.

  ### Parâmetros

  > Nenhum

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Response
  ```js
  {
    "noticias": []object || null,
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Notícias encontradas com sucesso!",
    "noticias": [
      {
        "codNotc": 1,
        "codAutor": "OGrandePoderosoB",
        "titulo": "AFEB continua existindo",
        "noticia": "Atualização para o dia de hoje: A AFEB ainda existe.",
        "dataPublicacao": "2023-11-24 11:23:50"
      }
    ]
  }
  ```
</details>

<details>
  <summary> <code>GET</code> <code>/api/noticias/${codNotc}</code> </summary>

  ### Descrição

  Retorna uma notícia específica do jornal do CCX a partir de seu código.

  ### Parâmetros

  | param.    |  tipo     | data type   | desc.                                                            |
  |-----------|-----------|-------------|------------------------------------------------------------------|
  | codNotc   |  required | string      | Código de notícia a ser retornada.                               |

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Response
  ```js
  {
    "noticia": object || null,
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Notícia encontrada com sucesso!",
    "noticia": {
      "codNotc": 7,
      "codAutor": "LucasMoraesGuede",
      "titulo": "Murilo Holtz se abstém do Torneio de Chess 960",
      "noticia": "A partir de hoje, quinta-feira (16), Murilo Holtz não irá mais participar do Torneio de Chess 960, pois irá fazer uma viagem aos Estados Unidos, e não será capaz de comparecer ao torneio até sexta-feira da próxima semana (24). Assim, o torneio agora acontecerá apenas entre 8 participantes e com um forte competidor a menos.",
      "dataPublicacao": "2023-11-16 17:38:59"
    }
  }
  ```
</details>

<details>
  <summary> <code>POST</code> <code>/api/noticias</code> </summary>

  ### Descrição

  Cria uma nova notícia.

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 201 | `CREATED` |
  | 400 | `BAD REQUEST` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Request body
  ```js
  {
    "codAutor": string,
    "titulo": string,
    "noticia": string
  }
  ```

  Exemplo:
  ```js
  {
    "codAutor": "OGrandePoderosoB",
    "titulo": "AFEB continua existindo",
    "noticia": "Atualização para o dia de hoje: A AFEB ainda existe."
  }
  ```

  ### Response
  ```js
  {
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Notícia criada com sucesso!"
  }
  ```
</details>

<details>
  <summary> <code>PUT</code> <code>/api/noticias</code> </summary>

  ### Descrição

  Atualiza os dados de uma notícia

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 400 | `BAD REQUEST` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Request body
  ```js
  {
    "codNotc": Number,
    "codAutor": string,
    "titulo": string,
    "noticia": string
  }
  ```

  Exemplo:
  ```js
  {
    "codNotc": 2,
    "codAutor": "OGrandePoderosoB",
    "titulo": "AFEB continua existindo",
    "noticia": "Atualização para o dia de hoje: A AFEB ainda existe."
  }
  ```

  ### Response
  ```js
  {
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Notícia atualizada com sucesso!"
  }
  ```
</details>

<details>
  <summary> <code>DELETE</code> <code>/api/noticias/${codNotc}</code> </summary>

  ### Descrição

  Exclui uma notícia do sistema.

  ### Parâmetros

  | param.    |  tipo     | data type   | desc.                                                            |
  |-----------|-----------|-------------|------------------------------------------------------------------|
  | codNotc   |  required | string      | Código da notícia a ser excluída.                                |

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 400 | `BAD REQUEST` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Response
  ```js
  {
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Notícia excluída com sucesso!"
  }
  ```
</details>

## /api/torneios

<details>
  <summary> <code>GET</code> <code>/api/torneios</code> </summary>

  ### Descrição

  Retorna os dados de todos os torneios do CCX.

  ### Parâmetros

  > Nenhum

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Response
  ```js
  {
    "torneios": []object || null,
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Torneios encontrados com sucesso!",
    "torneios": [
      {
        "codTorn": 0,
        "titulo": "Simultânea Oshiro vs. AFEB",
        "descricao": "Torneio de toda a AFEB contra o Oshiro jogando xadrez Oshiro.",
        "comentarios": "Oshiro destruiu a todos.",
        "dataInicio": "2024-01-01",
        "dataFim": "2024-01-02",
        "modo": "presencial",
        "participantes": 10,
        "placarFinal": "1º. Daniel Oshiro - 10/10\n2º. AFEB - 0/10"
      }
    ]
  }
  ```
</details>

<details>
  <summary> <code>GET</code> <code>/api/torneios/${codTorn}</code> </summary>

  ### Descrição

  Retorna os dados de um torneio específico do CCX.

  ### Parâmetros

  | param.    |  tipo     | data type   | desc.                                                            |
  |-----------|-----------|-------------|------------------------------------------------------------------|
  | codTorn   |  required | string      | Código de torneio a ser retornado.                               |

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Response
  ```js
  {
    "torneio": object || null,
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Torneio encontrado com sucesso!",
    "torneio": {
      "codTorn": 0,
      "titulo": "Simultânea Oshiro vs. AFEB",
      "descricao": "Torneio de toda a AFEB contra o Oshiro jogando xadrez Oshiro.",
      "comentarios": "Oshiro destruiu a todos.",
      "dataInicio": "2024-01-01",
      "dataFim": "2024-01-02",
      "modo": "presencial",
      "participantes": 10,
      "placarFinal": "1º. Daniel Oshiro - 9/9\n2º. AFEB - 0/9"
    }
  }
  ```
</details>

<details>
  <summary> <code>POST</code> <code>/api/torneios</code> </summary>

  ### Descrição

  Registra um novo torneio.

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 201 | `CREATED` |
  | 400 | `BAD REQUEST` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Request body
  ```js
  {
    "titulo": string,
    "descricao": string,
    "comentarios": string || null,
    "dataInicio": string,
    "dataFim": string || null,
    "modo": string,
    "participantes": Number,
    "placarFinal": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "titulo": "Simultânea Oshiro vs. AFEB",
    "descricao": "Torneio de toda a AFEB contra o Oshiro jogando xadrez Oshiro.",
    "comentarios": "Oshiro destruiu a todos.",
    "dataInicio": "2024-01-01",
    "dataFim": "2024-01-02",
    "modo": "presencial",
    "participantes": 10,
    "placarFinal": "1º. Daniel Oshiro - 9/9\n2º. AFEB - 0/9"
  }
  ```

  ### Response
  ```js
  {
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Torneio criado com sucesso!"
  }
  ```
</details>

<details>
  <summary> <code>PUT</code> <code>/api/torneios</code> </summary>

  ### Descrição

  Atualiza os dados de um torneio.

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 400 | `BAD REQUEST` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Request body
  ```js
  {
    "codTorn": Number,
    "titulo": string,
    "descricao": string,
    "comentarios": string || null,
    "dataInicio": string,
    "dataFim": string || null,
    "modo": string,
    "participantes": Number,
    "placarFinal": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "codTorn": 0,
    "titulo": "Simultânea Oshiro vs. AFEB",
    "descricao": "Torneio de toda a AFEB contra o Oshiro jogando xadrez Oshiro.",
    "comentarios": "Oshiro destruiu a todos.",
    "dataInicio": "2024-01-01",
    "dataFim": "2024-01-02",
    "modo": "presencial",
    "participantes": 10,
    "placarFinal": "1º. Daniel Oshiro - 9/9\n2º. AFEB - 0/9"
  }
  ```

  ### Response
  ```js
  {
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Torneio editado com sucesso!"
  }
  ```
</details>

<details>
  <summary> <code>DELETE</code> <code>/api/torneios/${codTorn}</code> </summary>

  ### Descrição

  Exclui um torneio do sistema.

  ### Parâmetros

  | param.    |  tipo     | data type   | desc.                                                            |
  |-----------|-----------|-------------|------------------------------------------------------------------|
  | codTorn   |  required | string      | Código do torneio a ser excluído.                                |

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 400 | `BAD REQUEST` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Response
  ```js
  {
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Torneio excluído com sucesso!"
  }
  ```
</details>

## /api/trofeus

<details>
  <summary> <code>GET</code> <code>/api/trofeus/jogador/${codJog}</code> </summary>

  ### Descrição

  Retorna todas as premiações pela associação de um jogador do CCX.

  ### Parâmetros

  | param.    |  tipo     | data type   | desc.                                                            |
  |-----------|-----------|-------------|------------------------------------------------------------------|
  | codJog    |  required | string      | Código do jogador a ter os troféus retornados.                   |

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 400 | `BAD REQUEST` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Response
  ```js
  {
    "trofeus": []object || null,
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Troféus encontrados com sucesso!",
    "trofeus": [
      {
        "codTrof": 1,
        "codJog": 1,
        "codTorn": 1,
        "torneio": "1º Campeonato Mundial da Biblioteca AFEB",
        "posicao": 1
      },
      {
        "codTrof": 4,
        "codJog": 1,
        "codTorn": 2,
        "torneio": "1º Torneio Suíço de Blitz AFEB",
        "posicao": 1
      }
    ]
  }
  ```
</details>

<details>
  <summary> <code>GET</code> <code>/api/trofeus/${codTrof}</code> </summary>

  ### Descrição

  Retorna os dados de uma premiação a partir de seu código.

  ### Parâmetros

  | param.    |  tipo     | data type   | desc.                                                            |
  |-----------|-----------|-------------|------------------------------------------------------------------|
  | codTrof   |  required | string      | Código de troféu a ser retornado.                                |

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Response
  ```js
  {
    "trofeu": object || null,
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Troféu encontrado com sucesso!",
    "trofeu": {
      "codTrof": 1,
      "codJog": 1,
      "codTorn": 1,
      "torneio": "1º Campeonato Mundial da Biblioteca AFEB",
      "posicao": 1
    }
  }
  ```
</details>

<details>
  <summary> <code>POST</code> <code>/api/trofeus</code> </summary>

  ### Descrição

  Registra uma nova premiação.

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 201 | `CREATED` |
  | 400 | `BAD REQUEST` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Request body
  ```js
  {
    "codJog": Number,
    "codTorn": Number,
    "posicao": Number
  }
  ```

  Exemplo:
  ```js
  {
    "codJog": 1,
    "codTorn": 12,
    "posicao": 1
  }
  ```

  ### Response
  ```js
  {
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Troféu registrado com sucesso!"
  }
  ```
</details>

<details>
  <summary> <code>PUT</code> <code>/api/trofeus</code> </summary>

  ### Descrição

  Atualiza os dados de uma premiação.

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 400 | `BAD REQUEST` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Request body
  ```js
  {
    "codTrof": Number,
    "codJog": Number,
    "codTorn": Number,
    "posicao": Number
  }
  ```

  Exemplo:
  ```js
  {
    "codTrof": 5,
    "codJog": 1,
    "codTorn": 12,
    "posicao": 1
  }
  ```

  ### Response
  ```js
  {
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Troféu atualizado com sucesso!"
  }
  ```
</details>

<details>
  <summary> <code>DELETE</code> <code>/api/trofeus/${codTrof}</code> </summary>

  ### Descrição

  Exclui uma premiação do sistema.

  ### Parâmetros

  | param.    |  tipo     | data type   | desc.                                                            |
  |-----------|-----------|-------------|------------------------------------------------------------------|
  | codTrof   |  required | string      | Código do troféu a ser excluído.                                 |

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 400 | `BAD REQUEST` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Response
  ```js
  {
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Troféu excluído com sucesso!"
  }
  ```
</details>

## /api/usuarios

<details>
  <summary> <code>GET</code> <code>/api/usuarios</code> </summary>

  ### Descrição

  Retorna todos os usuários com acesso ao sistema do CCX.

  ### Parâmetros

  > Nenhum

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Response
  ```js
  {
    "usuarios": []object || null,
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Usuários encontrados com sucesso!",
    "usuarios": [
      {
        "codUsu": "A2rToPKHJe9B5PIOUYKgXQ==",
        "username": "Cléber",
        "senha": "",
        "adm": true,
        "dataReg": "2024-01-02"
      }
    ]
  }
  ```
</details>

<details>
  <summary> <code>GET</code> <code>/api/usuarios/${username}</code> </summary>

  ### Descrição

  Retorna um usuário a partir de seu username.

  ### Parâmetros

  | param.    |  tipo     | data type   | desc.                                                            |
  |-----------|-----------|-------------|------------------------------------------------------------------|
  | username  |  required | string      | Username de usuário a ser retornado.                             |

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Response
  ```js
  {
    "usuario": object || null,
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Usuário encontrado com sucesso!",
    "usuario": {
      "codUsu": "A2rToPKHJe9B5PIOUYKgXQ==",
      "username": "Cléber",
      "senha": "",
      "adm": true,
      "dataReg": "2024-01-02"
    }
  }
  ```
</details>

<details>
  <summary> <code>POST</code> <code>/api/usuarios</code> </summary>

  ### Descrição

  Cria um novo usuário.

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 201 | `CREATED` |
  | 400 | `BAD REQUEST` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Request body
  ```js
  {
    "username": string,
    "senha": string,
    "adm": boolean
  }
  ```

  Exemplo:
  ```js
  {
    "username": "Cléber",
    "senha": "senhaForte",
    "adm": true
  }
  ```

  ### Response
  ```js
  {
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Usuário cadastrado com sucesso!"
  }
  ```
</details>

<details>
  <summary> <code>PUT</code> <code>/api/usuarios</code> </summary>

  ### Descrição

  Atualiza os dados de um usuário.

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 400 | `BAD REQUEST` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Request body
  ```js
  {
    "codUsu": string,
    "username": string,
    "senha": string,
    "adm": boolean
  }
  ```

  Exemplo:
  ```js
  {
    "codUsu": "A2rToPKHJe9B5PIOUYKgXQ==",
    "username": "Cléber",
    "senha": "senhaForte",
    "adm": true
  }
  ```

  ### Response
  ```js
  {
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Usuário atualizada com sucesso!"
  }
  ```
</details>

<details>
  <summary> <code>DELETE</code> <code>/api/usuarios/${username}</code> </summary>

  ### Descrição

  Exclui um usuário do sistema.

  ### Parâmetros

  | param.    |  tipo     | data type   | desc.                                                            |
  |-----------|-----------|-------------|------------------------------------------------------------------|
  | username  |  required | string      | Username do usuário a ser excluído.                              |

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 400 | `BAD REQUEST` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Response
  ```js
  {
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Usuário excluído com sucesso!"
  }
  ```
</details>
