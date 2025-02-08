const API = "https://afeb-api.onrender.com/api";

const codJog = new URL(window.location.href).searchParams.get("codJog");

fetchPerfil(codJog)
.then((jogador) => {
    if (!jogador) {
        mostrarMensagemErro("Dados de jogador não cadastrados.");
        return;
    }

    mostrarJogador(jogador);
})
.catch((err) => {
    mostrarMensagemErro(err);
});

/**
 * Retorna um jogador da API.
 * @param {string} codJog - Número do código do jogador na AFEB.
 * @returns {object} - Dados do jogador do CCX.
 * @throws - Retorna um erro da API.
 */
async function fetchPerfil(codJog) {
    return await fetch(API + `/jogadores/${codJog}`)
    .then((res) => res.json())
    .then((res) => {
        if (res == null)
            return null;

        if (res.error) 
            return new Error(res.error);

        return res.jogador;
    })
    .catch((err) => {
        console.error(err);
        return new Error("Erro ao conectar com a API.");
    });
}

/**
 * Mostra o perfil do jogador com seus dados.
 * @param {object} jogador - Jogador a ser mostrado.
 */
function mostrarJogador(jogador) {
    document.getElementById("username-perfil-jogador").textContent = jogador.nome;
    document.getElementById("informacoes-jogador").textContent = jogador.info;
    
    if (jogador.apelido) {
        document.getElementById("apelido-perfil-jogador").textContent = `(${jogador.apelido})`;
    } else {
        document.getElementById("apelido-perfil-jogador").style.display = "none";
    }

    if (jogador.tituloAFEB) {
        document.getElementById("titulo-jogador").textContent = jogador.tituloAFEB;
    } else {
        const tituloContainer = document.getElementById("titulo-perfil-jogador-container");
        tituloContainer.style.display = "none";
    }

    document.getElementById("vitorias-jogador").textContent = jogador.vitorias;
    document.getElementById("derrotas-jogador").textContent = jogador.derrotas;
    document.getElementById("empates-jogador").textContent = jogador.empates;

    document.getElementById("rating-rapido-p").textContent = (
        jogador.eloRapid != null
        ? jogador.eloRapid
        : "-"
    );

    document.getElementById("rating-blitz-p").textContent = (
        jogador.eloBlitz != null
        ? jogador.eloBlitz
        : "-"
    );

    mostrarPremiacoes(jogador.trofeus);
}

/**
 * Mostra os torneios vencidos pelo jogador na AFEB.
 * @param {Array} trofeus - Lista de torneios vencidos pelo jogador.
 */
function mostrarPremiacoes(trofeus) {
    if (trofeus == null)
        return;

    const premiacoesContainer = document.getElementById("premiacoes-container");

    for (const t of trofeus) {
        const trofeuContainer = document.createElement("a");
        trofeuContainer.href = `torneio.html?torneio=${t.codTorn}`;

        const trofeuDiv = document.createElement("div");
        trofeuDiv.classList.add("trofeu-div");

        trofeuDiv.innerHTML = `
            <p class="font-bold"> ${t.torneio} - ${t.posicao}º Lugar </p>`;

        trofeuContainer.appendChild(trofeuDiv);
        premiacoesContainer.appendChild(trofeuContainer);
    }
}

/**
 * Mostra uma mensagem de erro ao usuário.
 * @param {string} erro - Mensagem de erro.
 */
function mostrarMensagemErro(erro) {
    alert(erro);
}
