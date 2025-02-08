package models

import (
	"log"
	"net/http"

	"gopkg.in/guregu/null.v3"
)

type Jogador struct {
	CodJog         int         `json:"codJog"`
	Nome           string      `json:"nome"`
	Apelido        null.String `json:"apelido"`
	Titulo         null.String `json:"titulo"`
	Info           null.String `json:"info"`
	EloClassic     null.Int    `json:"eloClassic"`
	EloRapid       null.Int    `json:"eloRapid"`
	EloBlitz       null.Int    `json:"eloBlitz"`
	Jogos          int         `json:"jogos"`
	Vitorias       int         `json:"vitorias"`
	Derrotas       int         `json:"derrotas"`
	Empates        int         `json:"empates"`
	DataNascimento string      `json:"dataNascimento"`
	Trofeus        Trofeus     `json:"trofeus"`
}

type RankingJogadores []Jogador

// Valida uma instância de Jogador
func (j *Jogador) IsValid() (bool, string) {
	if j.CodJog < 0 || j.CodJog > 99999999 {
		return false, "Código de jogador inválido."
	}

	if len(j.Nome) < 4 || len(j.Nome) > 70 {
		return false, "Nome de jogador deve conter de 4 a 70 caracteres."
	}

	if j.Apelido.Valid {
		if len(j.Apelido.String) > 20 {
			return false, "Apelido deve conter até 20 caracteres."
		}
	}

	if j.Titulo.Valid {
		if j.Titulo.String != "CMN" && j.Titulo.String != "MN" &&
			j.Titulo.String != "CMF" && j.Titulo.String != "MNF" &&
			j.Titulo.String != "ACM" && j.Titulo.String != "AFM" &&
			j.Titulo.String != "AIM" && j.Titulo.String != "AGM" &&
			j.Titulo.String != "CM" && j.Titulo.String != "FM" &&
			j.Titulo.String != "WCM" && j.Titulo.String != "WFM" &&
			j.Titulo.String != "IM" && j.Titulo.String != "GM" &&
			j.Titulo.String != "WIM" && j.Titulo.String != "WGM" &&
			j.Titulo.String != "MNO" {
			return false, "Título de mestre inválido."
		}
	}

	if j.Info.Valid {
		if len(j.Info.String) > 255 {
			return false,
				"Informações de jogador não podem ultrapassar 255 caracteres."
		}
	}

	if j.EloClassic.Valid {
		if j.EloClassic.Int64 < 0 || j.EloClassic.Int64 > 9999 {
			return false, "Rating clássico CCX deve estar entre 0 e 9999."
		}
	}

	if j.EloRapid.Valid {
		if j.EloRapid.Int64 < 0 || j.EloRapid.Int64 > 9999 {
			return false, "Rating rápido CCX deve estar entre 0 e 9999."
		}
	}

	if j.EloBlitz.Valid {
		if j.EloBlitz.Int64 < 0 || j.EloBlitz.Int64 > 9999 {
			return false, "Rating blitz CCX deve estar entre 0 e 9999."
		}
	}

	if j.Jogos == 0 {
		j.Jogos = j.Vitorias + j.Derrotas +
			j.Empates
	}

	if j.Jogos < 0 || j.Jogos > 99999999 {
		return false, "Jogos deve estar entre 0 e 99999999"
	}

	if j.Vitorias < 0 || j.Vitorias > 99999999 {
		return false, "Vitórias deve estar entre 0 e 99999999"
	}

	if j.Derrotas < 0 || j.Derrotas > 99999999 {
		return false, "Derrotas deve estar entre 0 e 99999999"
	}

	if j.Empates < 0 || j.Empates > 99999999 {
		return false, "Empates deve estar entre 0 e 99999999"
	}

	if len(j.DataNascimento) != 10 {
		return false, "Data de nascimento deve estar no padrão: YYYY-MM-DD."
	}

	return true, ""
}

// Retorna os jogadores do CCX em ordem de rating clássico.
func (r *RankingJogadores) GetTop50CCX() (int, string) {
	selectRanking :=
		`SELECT cod_jog, nome, apelido, titulo, info, elo_classico, elo_rapido,
		elo_blitz, jogos, vitorias, derrotas, empates, data_nascimento
		FROM Jogadores ORDER BY elo_classico DESC, elo_rapido DESC, elo_blitz DESC;`

	rows, err := E.DB.Query(selectRanking)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao receber ranking do banco de dados."
	}

	for rows.Next() {
		var j Jogador
		err := rows.Scan(&j.CodJog, &j.Nome, &j.Apelido, &j.Titulo, &j.Info,
			&j.EloClassic, &j.EloRapid, &j.EloBlitz, &j.Jogos, &j.Vitorias,
			&j.Derrotas, &j.Empates, &j.DataNascimento)
		if err != nil {
			log.Println(err)
			return http.StatusInternalServerError,
				"Erro ao receber jogadores do banco de dados."
		}

		*r = append(*r, j)
	}

	if err := rows.Err(); err != nil {
		return http.StatusInternalServerError,
			"Erro ao conectar ao banco de dados."
	}

	defer rows.Close()

	return http.StatusOK, ""
}

// Encontra dados de um jogador
func (j *Jogador) GetJogador(codJog string) (int, string) {
	selectJog := `
		SELECT cod_jog, nome, apelido, titulo, info, elo_classico, elo_rapido,
		elo_blitz, jogos, vitorias, derrotas, empates, data_nascimento
		FROM Jogadores WHERE cod_jog = ?;`
	row := E.DB.QueryRow(selectJog, codJog)

	err := row.Scan(&j.CodJog, &j.Nome, &j.Apelido, &j.Titulo, &j.Info,
		&j.EloClassic, &j.EloRapid, &j.EloBlitz, &j.Jogos, &j.Vitorias,
		&j.Derrotas, &j.Empates, &j.DataNascimento)
	if err != nil {
		log.Println(err)
		return http.StatusNotFound, "Jogador não encontrado."
	}

	statusCode, msgErro := j.Trofeus.GetTrofeusJogador(j.CodJog)

	return statusCode, msgErro
}

// Salva um jogador novo no banco de dados
func (j Jogador) RegistrarJogador() (int, string) {
	insert := `
		INSERT INTO Jogadores (nome, apelido, titulo, info, elo_classico, 
		elo_rapido, elo_blitz, jogos, vitorias, derrotas, empates, 
		data_nascimento) 
		VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	_, err := E.DB.Exec(insert, j.Nome, j.Apelido, j.Titulo, j.Info, j.EloClassic,
		j.EloRapid, j.EloBlitz, j.Jogos, j.Vitorias, j.Derrotas, j.Empates,
		j.DataNascimento)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao registrar usuário. Informe o erro a um mantenedor do " +
				"projeto ou abra uma issue em nosso github."
	}

	return http.StatusCreated, ""
}

// Atualiza os dados de um jogador no sistema
func (j Jogador) EditarJogador() (int, string) {
	update := `
		UPDATE Jogadores SET nome = ?, apelido = ?, titulo = ?, info = ?,
		elo_classico = ?, elo_rapido = ?, elo_blitz = ?, jogos = ?, vitorias = ?, 
		derrotas = ?, empates = ?, data_nascimento = ? WHERE cod_jog = ?;`

	_, err := E.DB.Exec(update, j.Nome, j.Apelido, j.Titulo, j.Info, j.EloClassic,
		j.EloRapid, j.EloBlitz, j.Jogos, j.Vitorias, j.Derrotas, j.Empates,
		j.DataNascimento, j.CodJog)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao editar jogador. Informe o erro a um mantenedor do " +
				"projeto ou abra uma issue em nosso github."
	}

	return http.StatusOK, ""
}

// Exclui um jogador do sistema
func (j Jogador) ExcluirJogador() (int, string) {
	delete := "DELETE FROM Jogadores WHERE cod_jog = ?;"
	_, err := E.DB.Exec(delete, j.CodJog)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao excluir jogador. Informe o erro a um mantenedor do " +
				"projeto ou abra uma issue em nosso github."
	}

	return http.StatusOK, ""
}

// Verifica se um jogador existe a partir de seu código de jogador CCX
func JogadorExiste(codJog int) bool {
	selectJog := "SELECT nome FROM Jogadores WHERE cod_jog = ?;"
	row := E.DB.QueryRow(selectJog, codJog)

	var nome string
	if err := row.Scan(&nome); err != nil {
		log.Println(err)
		return false
	}

	return true
}
