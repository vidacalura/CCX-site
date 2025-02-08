package models

import (
	"log"
	"net/http"

	"github.com/vidacalura/AFEB/internal/utils"
)

type Usuario struct {
	CodUsu   []byte `json:"codUsu"`
	Username string `json:"username"`
	Senha    string `json:"senha"`
	Adm      bool   `json:"adm"`
	DataReg  string `json:"dataReg"`
}

// Valida uma instância de Usuario
func (u Usuario) IsValid() (bool, string) {
	if len(u.CodUsu) > 16 {
		return false, "Código de usuário inválido."
	}

	if len(u.Username) < 3 || len(u.Username) > 30 {
		return false, "Username deve conter de 3 a 30 caracteres."
	}

	if len(u.Senha) < 8 || len(u.Senha) > 32 {
		return false, "Senha de usuário deve conter entre 8 e 32 caracteres"
	}

	return true, ""
}

type Usuarios []Usuario

// Retorna todos os usuários do sistema
func (u *Usuarios) GetUsuarios() (int, string) {
	selectUsu := "SELECT cod_usu, username, adm, data_reg FROM Usuarios;"

	rows, err := E.DB.Query(selectUsu)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao receber usuários do banco de dados."
	}

	for rows.Next() {
		var usu Usuario
		err := rows.Scan(&usu.CodUsu, &usu.Username, &usu.Adm, &usu.DataReg)
		if err != nil {
			log.Println(err)
			return http.StatusInternalServerError,
				"Erro ao receber usuários do banco de dados."
		}

		*u = append(*u, usu)
	}

	if err := rows.Err(); err != nil {
		return http.StatusInternalServerError,
			"Erro ao conectar ao banco de dados."
	}

	defer rows.Close()

	return http.StatusOK, ""
}

// Retorna as informações de um só usuário
func (u *Usuario) GetUsuario(username string) (int, string) {
	selectUsu := `
		SELECT cod_usu, username, adm, data_reg FROM Usuarios
		WHERE username = ?;`

	row := E.DB.QueryRow(selectUsu, username)
	err := row.Scan(&u.CodUsu, &u.Username, &u.Adm, &u.DataReg)
	if err != nil {
		log.Println(err)
		return http.StatusNotFound, "Usuário não encontrado."
	}

	return http.StatusOK, ""
}

// Cria um novo usuário
func (u Usuario) CriarUsuario() (int, string) {
	insert := `
		INSERT INTO Usuarios
		VALUES(UUID_TO_BIN(UUID()), ?, ?, ?, CURRENT_DATE());`

	u.Senha = utils.CriptografarSenha(u.Senha)

	_, err := E.DB.Exec(insert, u.Username, u.Senha, u.Adm)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao registrar usuário. Informe o erro a um mantenedor do " +
				"projeto ou abra uma issue em nosso github."
	}

	return http.StatusCreated, ""
}

// Edita um usuário
func (u Usuario) EditarUsuario() (int, string) {
	update := `
		UPDATE Usuarios SET username = ?, senha = ?, adm = ? 
		WHERE cod_usu = ?;`

	u.Senha = utils.CriptografarSenha(u.Senha)

	_, err := E.DB.Exec(update, u.Username, u.Senha, u.Adm, u.CodUsu)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao editar usuário. Informe o erro a um mantenedor do " +
				"projeto ou abra uma issue em nosso github."
	}

	return http.StatusOK, ""
}

// Exclui um usuário do sistema
func (u Usuario) ExcluirUsuario() (int, string) {
	// Verifica se usuário existe
	selectUsu := "SELECT username FROM Usuarios WHERE username = ?;"

	row := E.DB.QueryRow(selectUsu, u.Username)

	var username string
	if err := row.Scan(&username); err != nil {
		log.Println(err)
		return http.StatusNotFound, "Usuário a ser excluído não encontrado."
	}

	// Deleta usuário
	delete := "DELETE FROM Usuarios WHERE username = ?;"

	_, err := E.DB.Exec(delete, u.Username)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao excluir usuário. Informe o erro a um mantenedor do " +
				"projeto ou abra uma issue em nosso github."
	}

	return http.StatusOK, ""
}

// Valida dados de login de usuário
func (u Usuario) ValidarLogin() bool {
	selectUsu := `
		SELECT cod_usu FROM Usuarios
		WHERE BINARY username = ? AND BINARY senha = ?;`

	row := E.DB.QueryRow(selectUsu, u.Username, u.Senha)
	
	var codUsu []byte
	if err := row.Scan(&codUsu); err != nil {
		log.Println(err)
		return false
	}
	
	return true
}

// Valida se usuário existe
func UsuarioExiste(codUsu []byte) bool {
	selectUsu := "SELECT username FROM Usuarios WHERE cod_usu = ?;"

	row := E.DB.QueryRow(selectUsu, codUsu)

	var username string
	if err := row.Scan(&username); err != nil {
		log.Println(err)
		return false
	}

	return true
}
