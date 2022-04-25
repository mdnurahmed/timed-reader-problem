const readline = require('readline');


const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});
const input = new Promise((resolve) => {
    rl.question('Type something ', function(input) {
        resolve(input);
    });
})

const timeout = new Promise((resolve, reject) => {
    setTimeout(() => {
        reject('\nTimed out\n');
    }, 3000)
})

const race = Promise.race([input, timeout]);

race
    .then((res) => {
        console.log('user typed - ' + res);
        rl.close();
    })
    .catch((err) => {
        console.log(err);
        rl.close();
    });