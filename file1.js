const readline = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout
});

function soma(a, b){
    return Number(a)+Number(b);
}

readline.question(`Qual o primeiro numero ->`, num1 => {
    readline.question(`Qual o segundo numero ->`, num2 => {
        let res = soma(num1, num2)
        console.log(`a soma entre ${num1} e ${num2} é = ${res}`)
    });
});