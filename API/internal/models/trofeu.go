package models

import (
	"log"
	"net/http"
)

type Trofeu struct {
	CodTrof int    `json:"codTrof"`
	CodJog  int    `json:"codJog"`
	CodTorn int    `json:"codTorn"`
	Torneio string `json:"torneio"`
	Posicao int    `json:"posicao"`
}

type Trofeus []Trofeu

// Valida uma instância de Trofeu
func (t Trofeu) IsValid() (bool, string) {
	if t.CodTrof < 0 || t.CodTrof > 99999999 {
		return false, "Código de troféu inválido."
	}

	if t.CodJog <= 0 || t.CodJog > 99999999 {
		return false, "Código de jogador inválido."
	}

	if t.CodTorn <= 0 || t.CodTorn > 99999999 {
		return false, "Código de torneio inválido."
	}

	if t.Posicao <= 0 || t.Posicao > 3 {
		return false, "Posição do torneio deve ser de 1º, 2º ou 3º lugar."
	}

	return true, ""
}

// Retorna os troféus de um jogador específico
func (t *Trofeus) GetTrofeusJogador(codJog int) (int, string) {
	selectTrof := `
		SELECT Torneios.titulo, Trofeus.*
		FROM Trofeus
		INNER JOIN Torneios
		ON Trofeus.cod_torn = Torneios.cod_torn
		WHERE cod_jog = ?;`

	rows, err := E.DB.Query(selectTrof, codJog)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao retornar prêmios do jogador. Informe o erro a um" +
				"mantenedor do projeto ou abra uma issue em nosso github."
	}

	for rows.Next() {
		var trof Trofeu
		err := rows.Scan(&trof.Torneio, &trof.CodTrof, &trof.CodJog,
			&trof.CodTorn, &trof.Posicao)
		if err != nil {
			log.Println(err)
			return http.StatusInternalServerError,
				"Erro ao receber prêmios do jogador."
		}

		*t = append(*t, trof)
	}

	if err := rows.Err(); err != nil {
		return http.StatusInternalServerError,
			"Erro ao conectar ao banco de dados."
	}

	defer rows.Close()

	return http.StatusOK, ""
}

// Mostra os dados de um Trofeu
func (t *Trofeu) GetDadosTrofeu(codTrof string) (int, string) {
	selectTrof := `
		SELECT Torneios.titulo, Trofeus.*
		FROM Trofeus
		INNER JOIN Torneios
		ON Trofeus.cod_torn = Torneios.cod_torn
		WHERE Trofeus.cod_trof = ?;`

	row := E.DB.QueryRow(selectTrof, codTrof)
	err := row.Scan(&t.Torneio, &t.CodTrof, &t.CodJog, &t.CodTorn, &t.Posicao)
	if err != nil {
		log.Println(err)
		return http.StatusNotFound, "Troféu não encontrado."
	}

	return http.StatusOK, ""
}

// Registra uma nova premiação no sistema
func (t Trofeu) CriarTrofeu() (int, string) {
	insert := "INSERT INTO Trofeus (cod_jog, cod_torn, posicao) VALUES(?, ?, ?);"

	_, err := E.DB.Exec(insert, t.CodJog, t.CodTorn, t.Posicao)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao registrar troféu. Informe o erro a um mantenedor do " +
				"projeto ou abra uma issue em nosso github."
	}

	return http.StatusCreated, ""
}

// Atualiza os dados de uma premiação no banco de dados
func (t Trofeu) EditarTrofeu() (int, string) {
	update := `
		UPDATE Trofeus SET cod_jog = ?, cod_torn = ?, posicao = ?
		WHERE cod_trof = ?;`

	_, err := E.DB.Exec(update, t.CodJog, t.CodTorn, t.Posicao, t.CodTrof)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao atualizar troféu. Informe o erro a um mantenedor do " +
				"projeto ou abra uma issue em nosso github."
	}

	return http.StatusOK, ""
}

// Exclui uma premiação do sistema
func (t Trofeu) ExcluirTrofeu() (int, string) {
	delete := "DELETE FROM Trofeus WHERE cod_trof = ?;"
	_, err := E.DB.Exec(delete, t.CodTrof)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao excluir troféu. Informe o erro a um mantenedor do " +
				"projeto ou abra uma issue em nosso github."
	}

	return http.StatusOK, ""
}

// Verifica se um troféu existe a partir de seu código
func TrofeuExiste(codTrof int) bool {
	selectTrof := "SELECT cod_jog FROM Trofeus WHERE cod_trof = ?;"
	row := E.DB.QueryRow(selectTrof, codTrof)

	var codJog int
	if err := row.Scan(&codJog); err != nil {
		log.Println(err)
		return false
	}

	return true
}
