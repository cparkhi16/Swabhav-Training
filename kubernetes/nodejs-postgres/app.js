var express = require('express');
var path = require('path');
const db = require('./services/db');
var cookieParser = require('cookie-parser');
var logger = require('morgan');
var quotes =require('./services/quotes')
var indexRouter = require('./routes/index');
var quotesRouter = require('./routes/quotes');

var app = express();

app.use(logger('dev'));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));

app.use('/', indexRouter);
app.use('/quotes', quotesRouter);
const res=db.createTable("CREATE TABLE quote (id SERIAL PRIMARY KEY,quote character varying(255) NOT NULL UNIQUE,author character varying(255) NOT NULL,created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL)")
module.exports = app;
