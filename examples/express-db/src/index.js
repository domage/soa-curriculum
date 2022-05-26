const express = require('express')
const Redis = require('ioredis');
const app = express()

const port = Number(process.env.PORT)

const redis = new Redis({
  pkg: 'ioredis',
  host: process.env.REDIS_HOST,
  port: 6379,
  database: 0
});

app.get('/', (req, res) => {
  res.send('Hello World! Secret is ' + process.env.SECRET)
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})

setInterval(async () => {
  const values = await redis.get(`qwe`);

  console.log('Values from redis: ', values);
}, 1000);
