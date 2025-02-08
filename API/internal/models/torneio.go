package models

import (
	"log"
	"net/http"

	"gopkg.in/guregu/null.v3"
)

type Modo string

const (
	Online     Modo = "online"
	Presencial Modo = "presencial"
)

type Torneio struct {
	CodTorn       int         `json:"codTorn"`
	Titulo        string      `json:"titulo"`
	Descricao     string      `json:"descricao"`
	Comentarios   null.String `json:"comentarios"`
	DataInicio    string      `json:"dataInicio"`
	DataFim       null.String `json:"dataFim"`
	Modo          Modo        `json:"modo"`
	Participantes int         `json:"participantes"`
	PlacarFinal   null.String `json:"placarFinal"`
}

type Torneios []Torneio

// Valida uma instância de Torneio
func (t Torneio) IsValid() (bool, string) {
	if t.CodTorn < 0 || t.CodTorn > 99999999 {
		return false, "Código de torneio inválido."
	}

	if len(t.Titulo) < 7 || len(t.Titulo) > 100 {
		return false, "Título de torneio deve conter de 7 a 100 caracteres."
	}

	if len(t.Descricao) < 10 || len(t.Descricao) > 255 {
		return false, "Descrição do torneio deve conter de 10 a 255 caracteres."
	}

	if len(t.DataInicio) != 10 {
		return false, "Data de início do torneio deve estar no modelo YYYY-MM-DD."
	}

	if t.DataFim.Valid {
		if len(t.DataFim.String) != 10 {
			return false, "Data de encerramento do torneio deve estar no modelo YYYY-MM-DD."
		}
	}

	if t.Modo != Online && t.Modo != Presencial {
		return false, "Modo de torneio deve ser: online ou presencial."
	}

	if t.Participantes <= 1 || t.Participantes > 999 {
		return false, "Torneio não pode ter menos de 2 participantes ou mais que 999."
	}

	return true, ""
}

// Retorna todos torneios
func (t *Torneios) GetTorneios() (int, string) {
	selectTorn := "SELECT * FROM Torneios ORDER BY data_inicio DESC;"
	rows, err := E.DB.Query(selectTorn)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Não foi possível retornar dados de torneios."
	}

	for rows.Next() {
		var torn Torneio
		err := rows.Scan(&torn.CodTorn, &torn.Titulo, &torn.Descricao,
			&torn.Comentarios, &torn.DataInicio, &torn.DataFim, &torn.Modo,
			&torn.Participantes, &torn.PlacarFinal)
		if err != nil {
			log.Println(err)
			return http.StatusInternalServerError,
				"Erro ao receber torneios do banco de dados."
		}

		*t = append(*t, torn)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao conectar ao banco de dados."
	}

	defer rows.Close()

	return http.StatusOK, ""
}

// Retorna dados de um Torneio específico
func (t *Torneio) GetTorneio(codTorn string) (int, string) {
	selectTorn := "SELECT * FROM Torneios WHERE cod_torn = ?;"
	row := E.DB.QueryRow(selectTorn, codTorn)

	err := row.Scan(&t.CodTorn, &t.Titulo, &t.Descricao,
		&t.Comentarios, &t.DataInicio, &t.DataFim, &t.Modo,
		&t.Participantes, &t.PlacarFinal)
	if err != nil {
		log.Println(err)
		return http.StatusNotFound, "Torneio não encontrado."
	}

	return http.StatusOK, ""
}

// Registra novo Torneio
func (t Torneio) RegistrarTorneio() (int, string) {
	insert := `
		INSERT INTO Torneios (titulo, descricao, comentarios, data_inicio,
		data_fim, modo, participantes, placar_final) VALUES(?, ?, ?, ?, ?,
		?, ?, ?);`

	_, err := E.DB.Exec(insert, t.Titulo, t.Descricao, t.Comentarios,
		t.DataInicio, t.DataFim, t.Modo, t.Participantes, t.PlacarFinal)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao registrar torneio. Informe o erro a um mantenedor do " +
				"projeto ou abra uma issue em nosso github."
	}

	return http.StatusCreated, ""
}

// Edita um Torneio
func (t Torneio) EditarTorneio() (int, string) {
	update := `
		UPDATE Torneios SET titulo = ?, descricao = ?, comentarios = ?,
		data_inicio = ?, data_fim = ?, modo = ?, participantes = ?,
		placar_final = ? WHERE cod_torn = ?;`

	_, err := E.DB.Exec(update, t.Titulo, t.Descricao, t.Comentarios,
		t.DataInicio, t.DataFim, t.Modo, t.Participantes, t.PlacarFinal,
		t.CodTorn)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao registrar torneio. Informe o erro a um mantenedor do " +
				"projeto ou abra uma issue em nosso github."
	}

	return http.StatusOK, ""
}

// Exclui um Torneio
func (t Torneio) ExcluirTorneio() (int, string) {
	delete := "DELETE FROM Torneios WHERE cod_torn = ?;"
	_, err := E.DB.Exec(delete, t.CodTorn)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao excluir torneio. Informe o erro a um mantenedor do " +
				"projeto ou abra uma issue em nosso github."
	}

	return http.StatusOK, ""
}

// Verifica se Torneio existe
func TorneioExiste(codTorn int) bool {
	selectTorn := "SELECT titulo FROM Torneios WHERE cod_torn = ?;"
	row := E.DB.QueryRow(selectTorn, codTorn)

	var titulo string
	if err := row.Scan(&titulo); err != nil {
		log.Println(err)
		return false
	}

	return true
}
