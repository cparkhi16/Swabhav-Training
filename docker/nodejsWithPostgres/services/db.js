const { Pool } = require('pg');
const config = require('../config');
const pool = new Pool(config.db);

/**
 * Query the database using the pool
 * @param {*} query 
 * @param {*} params 
 * 
 * @see https://node-postgres.com/features/pooling#single-query
 */
async function query(query, params) {
    const {rows, fields} = await pool.query(query, params);

    return rows;
}

module.exports = {
  query
}
//docker exec -it <id> bash
//psql -h localhost -p 5432 -U postgres -W
//\c postgres
//SELECT * from quote;
//INSERT INTO quote  VALUES(22,'Hi','How r u');