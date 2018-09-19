'use strict';
const http = require('http');
const immutable = require('immutable');

const spoonerisms = immutable.List([
  'Cradley Booper',
  'Tame of Groans',
  'Jeef Berky',
  'Rot Poast',
  'Dive Follar Lootfong, from WubSay!',
  'Bustin Jieber',
  'Jilly Boel',
  'Boot rear',
  'Man diller',
])

http.createServer(function (req, res) {
    res.writeHead(200, {'Content-Type': 'text/plain'});
    let item = spoonerisms.get(Math.floor(Math.random()*spoonerisms.size))
    res.end(item);
}).listen(5000, "0.0.0.0");

console.log('Server running at http://127.0.0.1:5000/');
