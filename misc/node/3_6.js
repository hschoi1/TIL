//p109
const fs = require('fs');
fs.readFile('./example.txt',(err,data) => {
    if (err) {
        throw err;
    }
    console.log(data); // <Buffer 61 62 63 64 65 20 72 61 6e 64 6f 6d 20 74 78 74 20>
    console.log(data.toString()); //abcde random txt
}) 


fs.writeFile('./writeExample.txt', 'write anything', (err) => {
    if (err) {
        throw err;
    }
    fs.readFile('./writeExample.txt', (err, data) => {
        if (err) {
            throw err;
        }
        console.log(data.toString());  //write anything
    });
});