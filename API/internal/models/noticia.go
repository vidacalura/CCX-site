package models

import (
	"log"
	"net/http"
	"time"
)

type Noticia struct {
	CodNotc        int    `json:"codNotc"`
	CodAutor       []byte `json:"codAutor"`
	Autor          string `json:"autor"`
	Titulo         string `json:"titulo"`
	Noticia        string `json:"noticia"`
	DataPublicacao string `json:"dataPublicacao"`
}

type Feed []Noticia

// Valida uma instância de Noticia
func (n Noticia) IsValid() (bool, string) {
	if n.CodNotc < 0 || n.CodNotc > 99999999 {
		return false, "Código de notícia inválido."
	}

	if len(n.CodAutor) != 16 {
		return false, "Código de autor inválido."
	}

	if len(n.Titulo) < 4 || len(n.Titulo) > 255 {
		return false, "Título de notícia deve conter de 4 a 255 caracteres."
	}

	if n.Noticia == "" {
		return false, "Corpo da notícia não pode estar vazio."
	}

	return true, ""
}

// Retorna um feed com as 6 últimas notícias
func (f *Feed) GetFeed() (int, string) {
	selectFeed := `
		SELECT Noticias.*, Usuarios.username
		FROM Noticias
		INNER JOIN Usuarios
		ON Noticias.cod_autor = Usuarios.cod_usu
		ORDER BY data_publicacao DESC LIMIT 6;`

	rows, err := E.DB.Query(selectFeed)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao retornar últimas notícias. Tente novamente mais tarde."
	}

	for rows.Next() {
		var n Noticia
		err := rows.Scan(&n.CodNotc, &n.CodAutor, &n.Titulo, &n.Noticia,
			&n.DataPublicacao, &n.Autor)
		if err != nil {
			log.Println(err)
			return http.StatusInternalServerError,
				"Erro ao receber notícias do banco de dados."
		}

		*f = append(*f, n)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao conectar ao banco de dados."
	}

	defer rows.Close()

	return http.StatusOK, ""
}

// Retorna o feed completo de notícias, com todas as notícias do CCX.
func (f *Feed) GetNoticias() (int, string) {
	selectFeed := "SELECT * FROM Noticias ORDER BY data_publicacao DESC;"
	rows, err := E.DB.Query(selectFeed)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao receber notícias do banco de dados."
	}

	defer rows.Close()

	for rows.Next() {
		var n Noticia
		err := rows.Scan(&n.CodNotc, &n.CodAutor, &n.Titulo, &n.Noticia,
			&n.DataPublicacao)
		if err != nil {
			log.Println(err)
			return http.StatusInternalServerError,
				"Erro ao receber notícias do banco de dados."
		}

		*f = append(*f, n)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao conectar ao banco de dados."
	}

	return http.StatusOK, ""
}

// Recebe uma notícia especificada pelo seu código
func (n *Noticia) GetNoticia(codNotc string) (int, string) {
	selectNotc := `
		SELECT Noticias.*, Usuarios.username
		FROM Noticias
		INNER JOIN Usuarios
		ON Noticias.cod_autor = Usuarios.cod_usu
		WHERE cod_notc = ?;`

	row := E.DB.QueryRow(selectNotc, codNotc)

	err := row.Scan(&n.CodNotc, &n.CodAutor, &n.Titulo, &n.Noticia,
		&n.DataPublicacao, &n.Autor)
	if err != nil {
		log.Println(err)
		return http.StatusNotFound, "Notícia não encontrada."
	}

	return http.StatusOK, ""
}

// Cria uma nova notícia
func (n Noticia) CriarNoticia() (int, string) {
	insert := `
		INSERT INTO Noticias (cod_autor, titulo, noticia, data_publicacao)
		VALUES(?, ?, ?, ?);`

	_, err := E.DB.Exec(insert, n.CodAutor, n.Titulo, n.Noticia,
		time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao registrar notícia. Informe o erro a um mantenedor do " +
				"projeto ou abra uma issue em nosso github."
	}

	return http.StatusCreated, ""
}

// Edita o título e/ou corpo de uma notícia
func (n Noticia) EditarNoticia() (int, string) {
	update := `
		UPDATE Noticias SET titulo = ?, noticia = ? WHERE cod_notc = ?;`

	_, err := E.DB.Exec(update, n.Titulo, n.Noticia, n.CodNotc)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao editar notícia. Informe o erro a um mantenedor do " +
				"projeto ou abra uma issue em nosso github."
	}

	return http.StatusOK, ""
}

// Exclui uma notícia do sistema
func (n Noticia) ExcluirNoticia() (int, string) {
	delete := "DELETE FROM Noticias WHERE cod_notc = ?;"
	_, err := E.DB.Exec(delete, n.CodNotc)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao excluir notícia. Informe o erro a um mantenedor do " +
				"projeto ou abra uma issue em nosso github."
	}

	return http.StatusOK, ""
}

// Verifica se notícia com um dado código existe
func NoticiaExiste(codNotc int) bool {
	selectNotc := "SELECT titulo FROM Noticias WHERE cod_notc = ?;"
	row := E.DB.QueryRow(selectNotc, codNotc)

	var titulo string
	if err := row.Scan(&titulo); err != nil {
		log.Println(err)
		return false
	}

	return true
}
