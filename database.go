package main

const SCHEMA_CLIENT = `CREATE TABLE client (
	identifier text,
	rescue text,
	password text);`

const SCHEMA_CHANNEL = `CREATE TABLE channel (
	identifier text,
	topic text);`

const SCHEMA_MODE = `CREATE TABLE mode (
	client text,
	channel text,
	mode text,
	value text);`

const SCHEMA_ROLE = `CREATE TABLE role (
	identifier text,
	client text,
	channel text,
	access integer);`
