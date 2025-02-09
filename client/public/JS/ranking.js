//const API = "https://afeb-api.onrender.com/api";
const API = "http://127.0.0.1:4000/api";

fetchRanking()
.then((ranking) => {
    if (ranking)
        mostrarRanking(ranking);
})
.catch((err) => {
    mostrarMensagemErro(err);
});

/**
 * Retorna o ranking CCX.
 * @returns {Array} - Ranking de jogadores CCX.
 * @throws - Retorna um erro da API.
 */
async function fetchRanking() {
    return await fetch(API + "/jogadores/ranking")
    .then((res) => res.json())
    .then((res) => {
        if (res == null)
            return null;

        if (res.error) 
            return new Error(res.error);

        return res.ranking;
    })
    .catch((err) => {
        console.error(err);
        return new Error("Erro ao conectar com a API.");
    });
}

/**
 * Coloca os dados do ranking na tabela.
 * @param {Array} ranking - Ranking de jogadores CCX.
 */
function mostrarRanking(ranking) {
    if (!ranking)
        return;

    const rankingTabela = document.getElementById("tabela-ranking");

    let i = 1;
    for (const j of ranking) {
        const containerJogadorRanking = document.createElement("a");
        containerJogadorRanking.classList.add("ranking-jogador-div");
        containerJogadorRanking.href = `perfil.html?codJog=${j.codJog}`;

        containerJogadorRanking.innerHTML = `
            <div class="jogador-ranking-nome-container">
                <p> ${i} </p>
                <p> <b>${!j.titulo ? "" : j.titulo}</b> ${j.nome} </p>
            </div>

            <div class="ranking-jogador-rating-div">
                <p> ${(!j.eloClassic ? "Sem" : j.eloClassic)} </p>
                <p> ${(!j.eloRapid ? "Sem" : j.eloRapid)} </p>
                <p> ${(!j.eloBlitz ? "Sem" : j.eloBlitz)} </p>
            </div>
        `;

        rankingTabela.appendChild(containerJogadorRanking);

        i++;
    }
}

/**
 * Mostra uma mensagem de erro ao usuário.
 * @param {string} erro - Mensagem de erro.
 */
function mostrarMensagemErro(erro) {
    alert(erro);
}
