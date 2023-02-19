package data

var SELECT_COUNT_TABLEINFO = `
SELECT count(*) from TABLEINFO where _tablename=$1
`
