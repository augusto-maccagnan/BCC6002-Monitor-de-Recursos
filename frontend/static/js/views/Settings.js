import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("Settings");
    }

    async getHtml() {
        let json = JSON.parse(localStorage.getItem("timer"));
        if (json === null) {
            localStorage.setItem("timer", JSON.stringify({ tempoAtualizacao: 2 }));
        }
        json = JSON.parse(localStorage.getItem("timer"));
        return `
            <p>
                Alterar tempo de atualização de dados: &nbsp
                <form id="formTempoAtualizacao">
                <input type="number" id="tempoAtualizacao" name="tempoAtualizacao" min="1" max="15" value="${
                    json.tempoAtualizacao ? json.tempoAtualizacao : 2
                }">
                </form>
                &nbsp segundo(s)
            </p>
        `;
    }
}
