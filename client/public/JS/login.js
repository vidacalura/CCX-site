const btnLogin = document.getElementById("login-btn");
btnLogin.addEventListener("click", (e) => {
    e.preventDefault();
    validarLogin();
});

/**
 * Permite o usuário visualizar sua senha na textbox de login.
 */
function verCertinho() {
    const passwordInput = document.getElementById('password');
    const passwordIcon = document.querySelector('.icon-right');

    if (passwordInput.type === 'password') {
        passwordInput.type = 'text';
        passwordIcon.classList.remove('fa-eye');
        passwordIcon.classList.add('fa-eye-slash');
    } else {
        passwordInput.type = 'password';
        passwordIcon.classList.remove('fa-eye-slash');
        passwordIcon.classList.add('fa-eye');
    }
}

/**
 * Envia os dados de login à API para validação e recebimento
 * de token JWT para criação de sessão.
 */
function validarLogin() {
    const usernameInput = document.getElementById('username');
    const passwordInput = document.getElementById('password');

    fetch("http://localhost:4000/api/auth/login", {
        method: "POST",
        headers: {
            "Authorization": `Bearer ${getCookieToken()}`,
            "Content-type": "Application/JSON" 
        },
        body: JSON.stringify({
            username: usernameInput.value.trim(),
            senha: passwordInput.value.trim()
        })
    })
    .then((res) => res.json())
    .then((res) => {
        if (res.error) {
            mostrarMensagemErroLogin(res.error);
        }

        document.cookie = "SameSite=strict";
        document.cookie = "HttpOnly";
        document.cookie = `token=${res.token}`
        document.cookie = `username=${res.username}`;
        window.location.href = "index.html";
    });
}

/**
 * Pega o token dos cookies.
 * @returns {string|null} - Valor do token.
 */
function getCookieToken() {
    for (const cookieValue of document.cookie.split(";")) {
        if (cookieValue.includes("token=")) {
            return cookieValue.split("=")[1];
        }
    }

    return null;
}

/**
 * Mostra um erro ao realizar login ao usuário.
 * @param {string} erro - Erro a ser mostrado.
 */
function mostrarMensagemErroLogin(erro) {
    alert(erro);
}
