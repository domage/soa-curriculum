const express = require('express')
const Redis = require('ioredis');
const app = express()
const port = 3000

// const redis = new Redis({
//   pkg: 'ioredis',
//   host: 'redis',
//   port: 6379,
//   database: 0
// });

app.get('/', (req, res) => {
  req.smth = "smth";
  res.send('Hello World! - Express version!')
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})

// setInterval(async () => {
//   const values = await redis.get(`qwe`);

//   console.log('Values from redis: ', values);
// }, 1000);
