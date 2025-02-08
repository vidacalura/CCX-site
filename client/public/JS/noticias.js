const API = "https://afeb-api.onrender.com/api";

fetchNoticias()
.then((noticias) => {
    if (noticias)
        mostrarNoticias(noticias);
})
.catch((err) => {
    mostrarMensagemErro(err);
});

/**
 * Retorna o feed de notícias da API.
 * @returns {Array} - Array com últimas notícias do CCX.
 * @throws - Retorna um erro da API.
 */
async function fetchNoticias() {
    return await fetch(API + "/noticias/feed")
    .then((res) => res.json())
    .then((res) => {
        if (res == null)
            return null;

        if (res.error) 
            return new Error(res.error);

        return res.noticias;
    })
    .catch((err) => {
        console.error(err);
        return new Error("Erro ao conectar com a API.");
    });
}

/**
 * Mostra as últimas notícias no feed.
 * @param {Array} noticias - Array com últimas notícias do CCX.
 */
function mostrarNoticias(noticias) {
    const noticiasSection = document.getElementById("noticias-section");

    for (const n of noticias) {
        const noticiaDiv = document.createElement("div");
        noticiaDiv.classList.add("noticia-container");

        noticiaDiv.innerHTML = `
            <h4 class="header-noticia">
                ${n.autor} - ${formatarDataPublicacao(n.dataPublicacao)}
            </h4>
            <h2 class="titulo-noticia">
                ${n.titulo}
            </h2>
            <p class="corpo-noticia">
                ${n.noticia}
            </p>`;

        noticiasSection.appendChild(noticiaDiv);
    }
}

/**
 * Formata uma data vinda da API.
 * @param {string} dataPublicacao - Data de publicação de notícia.
 * @returns {string} - Data formatada.
 */
function formatarDataPublicacao(dataPublicacao) {
    dataPublicacao = dataPublicacao.split("-");
    dataPublicacao[2] = dataPublicacao[2].split(" ");

    const mesesMap = new Map();
    mesesMap.set("01", "janeiro");
    mesesMap.set("02", "fevereiro");
    mesesMap.set("03", "março");
    mesesMap.set("04", "abril");
    mesesMap.set("05", "maio");
    mesesMap.set("06", "junho");
    mesesMap.set("07", "julho");
    mesesMap.set("08", "agosto");
    mesesMap.set("09", "setembro");
    mesesMap.set("10", "outubro");
    mesesMap.set("11", "novembro");
    mesesMap.set("12", "dezembro");

    const dia = dataPublicacao[2][0];
    const mes = mesesMap.get(dataPublicacao[1]);

    return `${dia} de ${mes}`;
}

/**
 * Mostra uma mensagem de erro ao usuário.
 * @param {string} erro - Mensagem de erro.
 */
function mostrarMensagemErro(erro) {
    alert(erro);
}