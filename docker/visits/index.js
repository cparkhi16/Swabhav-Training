const express = require('express');
const redis = require('redis');
//const process = require('process');
const app = express();
const client = redis.createClient({
  host: 'redis-server',
  port: 6379,
});
client.set('visits', 0);

app.get('/', (req, res) => { 
  //process.exit(0); exit this process without any error code , hence this node app container will stop and if restart always is mentioned , this container will be restarted again
  //process.exit(1); exit this process with error code (1), hence this node app container will stio and we need restart on failure policy to again rerun this container 
  client.get('visits', (err, visits) => {
    res.send('Number of visits ' + visits);
    client.set('visits', parseInt(visits) + 1);
  });
});

app.listen(8081, () => {
  console.log('listening on port 8081');
});
