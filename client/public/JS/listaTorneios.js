//const API = "https://afeb-api.onrender.com/api";
const API = "http://127.0.0.1:4000/api";

fetchTorneio()
.then((torn) => {
    mostrarTorneio(torn);
})
.catch((err) => {
    mostrarMensagemErro(err);
});

/**
 * Retorna todos os torneios do CCX.
 * @returns {Array} - Torneios do CCX.
 * @throws - Retorna um erro da API.
 */
async function fetchTorneio() {
    return await fetch(API + "/torneios")
    .then((res) => res.json())
    .then((res) => {
        if (res == null)
            return null;

        if (res.error) 
            return new Error(res.error);

        return res.torneios;
    })
    .catch((err) => {
        console.error(err);
        return new Error("Erro ao conectar com a API.");
    });
}

/**
 * Coloca os dados de torneios na tabela.
 * @param {Array} torneios - Torneios do CCX.
 */
function mostrarTorneio(torneios) {
    if (!torneios)
        return;

    const torneioTabela = document.getElementById("tabela-torneio");

    for (const t of torneios) {
        const containerTorneio = document.createElement("a");
        containerTorneio.classList.add("ranking-torneio-div");
        containerTorneio.href = `torneio.html?torneio=${t.codTorn}`;

        containerTorneio.innerHTML = `
            <div class="ranking-torneio-div-title">
                <p> ${t.titulo} </p>
            </div>
            <div class="ranking-torneio-div-div">
                <p> ${t.modo} </p>
            </div>
            <div class="ranking-torneio-div-div">
                <p> ${formatarData(t.dataInicio)} </p>
            </div>
            <div class="ranking-torneio-div-div">
                <p> ${(!t.dataFim ? "" : formatarData(t.dataFim))} </p>
            </div>
        `;

        torneioTabela.appendChild(containerTorneio);
    }
}

/**
 * Mostra uma mensagem de erro ao usuário.
 * @param {string} erro - Mensagem de erro.
 */
function mostrarMensagemErro(erro) {
    alert(erro);
}

/**
 * Formata uma data em YYYY-MM-DD para DD/MM/YYYY.
 * @param {string} data - Data a ser formatada. Ex: 1970-01-01.
 * @returns {string} - Data formatada em DD/MM/YYYY.
 */
function formatarData(data) {
    data = data.split("T")[0];
    const [ano, mes, dia] = data.split("-");
    return `${dia}/${mes}/${ano}`;
}
