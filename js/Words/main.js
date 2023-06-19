const express = require('express')
const server = express()

server.set('view engine', 'ejs');

const wordRouter = require('./routes/word.routes')

const PORT = 3000

server.use(express.json())

server.use('/api', wordRouter)
server.listen(PORT, () => {
    console.log(`Server started: http://localhost:${PORT}`)
})
