const express = require('express')

const server = express()

const home = '/'
const word = '/word'

server.get(home, (req, res) => {
    res.sendFile(__dirname + '/index.html')
})

server.get(word, (req, res) => {
    res.sendFile(__dirname + '/word.html')
    con()
})

const PORT = 3000

server.listen(PORT, () => {
    console.log(`Server started: http://localhost:${PORT}`)
})

function con() {
    console.log('function con')
}