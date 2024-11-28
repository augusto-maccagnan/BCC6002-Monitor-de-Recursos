import AbstractView from "./AbstractView.js";
const url = "http://localhost:8080/resources";
const max_frequency_label = document.querySelector("#max_frequency_label");
const min_frequency_label = document.querySelector("#min_frequency_label");
const total_use_label = document.querySelector("#total_use_label");
const percentage_label = document.querySelector("#percentage_label");
const core_count_label = document.querySelector("#core_count_label");
const total_cpu = document.querySelector("#total_cpu");
const core_count = document.getElementById("core_count");
const max_frequency = document.getElementById("max_frequency");
const min_frequency = document.getElementById("min_frequency");
const total_use = document.getElementById("total_use");
const percentage = document.getElementById("percentage");
const app = document.getElementById("app");

export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("Dashboard");
    }

    async getHtml() {
        app.style.display = "none";
        let cpu_params, cpu_cores, disks, memory;
        let json = JSON.parse(localStorage.getItem("timer"));
        if (json === null) {
            localStorage.setItem("timer", JSON.stringify({ tempoAtualizacao: 2 }));
        }
        json = JSON.parse(localStorage.getItem("timer"));
        setInterval(
            () => {
                fetch(url)
                    .then((response) => response.json())
                    .then((data) => {
                        const temp = JSON.parse(JSON.stringify(data));
                        let dataItem = data[0];
                        cpu_params = dataItem.cpu[0];
                        cpu_cores = dataItem.cpu_cores;
                        disks = dataItem.disk;
                        memory = dataItem.memory[0];
                        core_count.textContent = cpu_params.core_number;
                        max_frequency.textContent = cpu_params.max_frequency;
                        min_frequency.textContent = cpu_params.min_frequency;
                        total_use.textContent = cpu_params.total_use;
                        percentage.textContent = cpu_params.percentage + "%";
                    })
                    .catch((err) => console.log(err));
            },
            json.tempoAtualizacao ? json.tempoAtualizacao * 1000 : 2 * 1000
        );
        fetch(url)
            .then((response) => response.json())
            .then((data) => {
                const temp = JSON.parse(JSON.stringify(data));
                let dataItem = data[0];
                cpu_params = dataItem.cpu[0];
                cpu_cores = dataItem.cpu_cores;
                disks = dataItem.disk;
                memory = dataItem.memory[0];
                core_count.textContent = cpu_params.core_number;
                max_frequency.textContent = cpu_params.max_frequency;
                min_frequency.textContent = cpu_params.min_frequency;
                total_use.textContent = cpu_params.total_use;
                percentage.textContent = cpu_params.percentage + "%";
            })
            .catch((err) => console.log(err));

        total_cpu.insertAdjacentHTML("afterbegin", "Total CPU");
        core_count_label.insertAdjacentHTML("afterbegin", "Número de cores: ");
        max_frequency_label.insertAdjacentHTML("afterbegin", "Frequência máx: ");
        min_frequency_label.insertAdjacentHTML("afterbegin", "Frequência mín: ");
        total_use_label.insertAdjacentHTML("afterbegin", "Frequência atual: ");
        percentage_label.insertAdjacentHTML("afterbegin", "Porcentagem de uso: ");
        return ``;
    }
}
