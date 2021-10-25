const express = require('express');
const http = require('http');
const mongodb = require('mongodb');
const randomstring = require('randomstring');

const app = express();
const server = http.createServer(app);
app.use(express.text({type: '*/*'}))
app.use(express.static(__dirname));

let database;
(async () => {
    const mongoClient = mongodb.MongoClient;
    const db = await mongoClient.connect('mongodb://localhost:27017/');
    database = db.db('url-shortener');
})();

app.post('/r', async (req, res) => {
    const url = req.body;
    const code = randomstring.generate(7);
    const obj = { url, code};
    await database.collection('urls').insertOne(obj);
    res.send(obj);
});

app.get('/r/:code', async (req, res) => {
    const code = req.params.code;
    const result = await database.collection('urls').findOne({ code });
    res.redirect(result.url);
});

server.listen(80);