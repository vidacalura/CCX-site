//const API = "https://afeb-api.onrender.com/api";
const API = "http://127.0.0.1:4000/api";

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
    dataPublicacao = new Date(dataPublicacao);

    const mesesMap = new Map();
    mesesMap.set(1, "janeiro");
    mesesMap.set(2, "fevereiro");
    mesesMap.set(3, "março");
    mesesMap.set(4, "abril");
    mesesMap.set(5, "maio");
    mesesMap.set(6, "junho");
    mesesMap.set(7, "julho");
    mesesMap.set(8, "agosto");
    mesesMap.set(9, "setembro");
    mesesMap.set(10, "outubro");
    mesesMap.set(11, "novembro");
    mesesMap.set(12, "dezembro");

    const dia = dataPublicacao.getDay();
    const mes = mesesMap.get(dataPublicacao.getMonth() + 1);

    const hora = dataPublicacao.getHours();
    const minuto = dataPublicacao.getMinutes();

    return `${dia} de ${mes} às ${hora}:${minuto}`;
}

/**
 * Mostra uma mensagem de erro ao usuário.
 * @param {string} erro - Mensagem de erro.
 */
function mostrarMensagemErro(erro) {
    alert(erro);
}