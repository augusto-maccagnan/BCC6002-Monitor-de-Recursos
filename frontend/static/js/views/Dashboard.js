import AbstractView from "./AbstractView.js";
const url = "http://localhost:8080/resources";

const total_cpu = document.querySelector("#total_cpu");
const core_count_label = document.querySelector("#core_count_label");
const core_count = document.getElementById("core_count");
const max_frequency_label = document.querySelector("#max_frequency_label");
const max_frequency = document.getElementById("max_frequency");
const min_frequency_label = document.querySelector("#min_frequency_label");
const min_frequency = document.getElementById("min_frequency");
const total_use_label = document.querySelector("#total_use_label");
const total_use = document.getElementById("total_use");
const percentage_label = document.querySelector("#percentage_label");
const percentage = document.getElementById("percentage");

const cores_cpu = document.querySelector("#cores_cpu");

const all_disk = document.querySelector("#disks");

const memory_component = document.querySelector("#memory");
const total_memory_label = document.querySelector("#total_memory_label");
const total_memory = document.getElementById("total_memory");
const free_memory_label = document.querySelector("#free_memory_label");
const free_memory = document.getElementById("free_memory");
const used_memory_label = document.querySelector("#used_memory_label");
const used_memory = document.getElementById("used_memory");
const percentage_memory_label = document.querySelector("#percentage_memory_label");
const percentage_memory = document.getElementById("percentage_memory");
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
                        // const temp = JSON.parse(JSON.stringify(data));
                        let dataItem = data[0];
                        cpu_params = dataItem.cpu[0];
                        cpu_cores = dataItem.cpu_cores;
                        disks = dataItem.disk;
                        memory = dataItem.memory[0];

                        // Atualizando valores Totais de CPU
                        core_count.textContent = cpu_params.core_number;
                        max_frequency.textContent = cpu_params.max_frequency + " MHz";
                        min_frequency.textContent = cpu_params.min_frequency + " MHz";
                        total_use.textContent = cpu_params.total_use + " MHz";
                        percentage.textContent = cpu_params.percentage + "%";
                        // Atualizando valores de Memória
                        total_memory.textContent = memory.total / 1000 + " MB";
                        free_memory.textContent = memory.free / 1000 + " MB";
                        used_memory.textContent = memory.used / 1000 + " MB";
                        percentage_memory.textContent = memory.percentage + "%";
                        // Atualizando valores de Cores de CPU
                        cpu_cores.forEach((core, number) => {
                            const freq = document.getElementById(`core-${number}-frequency`);
                            freq.textContent = `${core.frequency} MHz`;
                            const percentage = document.getElementById(`core-${number}-percentage`);
                            percentage.textContent = `${core.percentage}%`;
                        });
                        disks.forEach((disk, number) => {
                            const total = document.getElementById(`disk-${number}-total`);
                            const free = document.getElementById(`disk-${number}-free`);
                            const used = document.getElementById(`disk-${number}-used`);
                            if (disk.total - 1000 < 0) {
                                total.textContent = `${disk.total} Bytes`;
                            } else if (disk.total - 1000000 < 0) {
                                total.textContent = `${Math.trunc(disk.total / 1000)} MB`;
                            } else {
                                total.textContent = `${Math.trunc(disk.total / 1000000)} GB`;
                            }
                            if (disk.free - 1000 < 0) {
                                free.textContent = `${disk.free} Bytes`;
                            } else if (disk.free - 1000000 < 0) {
                                free.textContent = `${Math.trunc(disk.free / 1000)} MB`;
                            } else {
                                free.textContent = `${Math.trunc(disk.free / 1000000)} GB`;
                            }
                            if (disk.used - 1000 < 0) {
                                used.textContent = `${disk.used} Bytes`;
                            } else if (disk.used - 1000000 < 0) {
                                used.textContent = `${Math.trunc(disk.used / 1000)} MB`;
                            } else {
                                used.textContent = `${Math.trunc(disk.used / 1000000)} GB`;
                            }
                        });
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
                max_frequency.textContent = cpu_params.max_frequency + " MHz";
                min_frequency.textContent = cpu_params.min_frequency + " MHz";
                total_use.textContent = cpu_params.total_use + " MHz";
                percentage.textContent = cpu_params.percentage + "%";

                total_memory.textContent = memory.total / 1000 + " MB";
                free_memory.textContent = memory.free / 1000 + " MB";
                used_memory.textContent = memory.used / 1000 + " MB";
                percentage_memory.textContent = memory.percentage + "%";

                console.log(disks);

                cpu_cores.forEach((core, number) => {
                    const elem = elementFromHtml(`
                        <div class="core-${number}">
                            <p>CPU ${number}</p>
                            <p>Frequência: <strong id="core-${number}-frequency">${core.frequency} Mhz</strong></p>
                            <p>Porcentagem de uso: <strong id="core-${number}-percentage">${core.percentage}%</strong></p>
                        </div>
                        `);
                    cores_cpu.appendChild(elem);
                });

                disks.forEach((disk, number) => {
                    let total, free, used;
                    if (disk.total - 1000 < 0) {
                        total = `${disk.total} Bytes`;
                    } else if (disk.total - 1000000 < 0) {
                        total = `${Math.trunc(disk.total / 1000)} MB`;
                    } else {
                        total = `${Math.trunc(disk.total / 1000000)} GB`;
                    }
                    if (disk.free - 1000 < 0) {
                        free = `${disk.free} Bytes`;
                    } else if (disk.free - 1000000 < 0) {
                        free = `${Math.trunc(disk.free / 1000)} MB`;
                    } else {
                        free = `${Math.trunc(disk.free / 1000000)} GB`;
                    }
                    if (disk.used - 1000 < 0) {
                        used = `${disk.used} Bytes`;
                    } else if (disk.used - 1000000 < 0) {
                        used = `${Math.trunc(disk.used / 1000)} MB`;
                    } else {
                        used = `${Math.trunc(disk.used / 1000000)} GB`;
                    }
                    const elem = elementFromHtml(`
                        <div class="disk-${number}">
                            <p>Disco ${number}</p>
                            <p>Nome: <strong>${disk.name}</strong></p>
                            <p>Capacidade: <strong id="disk-${number}-total">${total}</strong></p>
                            <p>Disponível: <strong id="disk-${number}-free">${free}</strong></p>
                            <p>Usado: <strong id="disk-${number}-used">${used}</strong></p>
                            <p>Porcentagem de uso: <strong>${disk.percentage}%</strong></p>
                        </div>
                        `);
                    all_disk.appendChild(elem);
                });
            })
            .catch((err) => console.log(err));

        total_cpu.insertAdjacentHTML("afterbegin", "Total CPU");
        core_count_label.insertAdjacentHTML("afterbegin", "Número de cores: ");
        max_frequency_label.insertAdjacentHTML("afterbegin", "Frequência máx: ");
        min_frequency_label.insertAdjacentHTML("afterbegin", "Frequência mín: ");
        total_use_label.insertAdjacentHTML("afterbegin", "Frequência atual: ");
        percentage_label.insertAdjacentHTML("afterbegin", "Porcentagem de uso: ");

        memory_component.insertAdjacentHTML("afterbegin", "Memória");
        total_memory_label.insertAdjacentHTML("afterbegin", "Total: ");
        free_memory_label.insertAdjacentHTML("afterbegin", "Livre: ");
        used_memory_label.insertAdjacentHTML("afterbegin", "Usado: ");
        percentage_memory_label.insertAdjacentHTML("afterbegin", "Portecentagem: ");
        return ``;
    }
}

function elementFromHtml(html) {
    const template = document.createElement("template");
    template.innerHTML = html.trim();
    return template.content.firstChild;
}
