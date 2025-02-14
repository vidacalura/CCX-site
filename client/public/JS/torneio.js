//const API = "https://afeb-api.onrender.com/api";
const API = "http://127.0.0.1:4000/api";

const codTorn = new URL(window.location.href).searchParams.get("torneio");

fetchTorneio(codTorn)
.then((torneio) => {
    if (!torneio) {
        mostrarMensagemErro("Dados de torneio não cadastrados.");
        return;
    }

    mostrarTorneio(torneio);
})
.catch((err) => {
    mostrarMensagemErro(err);
});

/**
 * Retorna os dados de um torneio pela API.
 * @param {string} codTorn - Código do torneio a ser retornado.
 * @returns {object} - Dados do torneio do CCX.
 * @throws - Retorna um erro da API.
 */
async function fetchTorneio(codTorn) {
    return await fetch(API + `/torneios/${codTorn}`)
    .then((res) => res.json())
    .then((res) => {
        if (res == null)
            return null;

        if (res.error) 
            return new Error(res.error);

        return res.torneio;
    })
    .catch((err) => {
        console.error(err);
        return new Error("Erro ao conectar com a API.");
    });
}

/**
 * Mostra os dados do torneio ao usuário.
 * @param {object} torneio - Dados do torneio a ser mostrado.
 */
function mostrarTorneio(torneio) {
    document.getElementById("titulo-torneio").textContent = torneio.titulo;
    document.getElementById("descricao-torneio").innerHTML = torneio.descricao;

    if (torneio.comentarios) {
        document.getElementById("comentarios-torneio").innerHTML = 
            torneio.comentarios.replaceAll("\n", "<br />");
    }

    document.getElementById("torneio-participantes").textContent = torneio.participantes;
    document.getElementById("torneio-resultados").innerHTML =
        torneio.placarFinal.replaceAll("\n", "<br />");
}
