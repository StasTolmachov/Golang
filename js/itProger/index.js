const express = require('express')

const server = express()

const PORT = 3000

const home = '/'
const word = '/word'


server.get(home, (req, res) => {
    res.sendFile(__dirname + '/index.html')
})

server.get(word, (req, res) => {
    res.sendFile(__dirname + '/word.html')
    con()
})



server.listen(PORT, () => {
    console.log(`Server started: http://localhost:${PORT}`)
})

function con() {
    console.log('function con')
}