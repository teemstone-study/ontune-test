package data

var InsertTableinfoStmt = `
INSERT INTO tableinfo values ($1, 0, $2, $2, 0);
`

var InsertOntuneinfoStmt = `
INSERT INTO ontuneinfo values (
	0, $1, $2, $3, $4, 0, 0, 0, 0, 0, null, $5, $6, $7, 0, $8, $9, $10, $11, $12, $13, 0, $14, $15, $16, $17, $18
)
`

var InsertAgentinfo = `
INSERT INTO agentinfo 
(select * from unnest($1::int[], $2::text[], $3::int[], $4::int[], $5::int[], $6::int[],
	$7::int[], $8::int[], $9::int[], $10::int[], $11::int[], $12::int[], $13::int[], 
	$14::text[], $15::text[], $16::text[], $17::text[], $18::text[], $19::text[], $20::text[],
	$21::int[], $22::int[], $23::int[], $24::int[], $25::int[],
	$26::text[], $27::int[], $28::int[]))
`

var InsertHostinfo = `
INSERT INTO hostinfo 
(select * from unnest($1::int[], $2::text[], $3::text[], $4::text[], $5::text[], $6::text[],
	$7::text[], $8::text[], $9::int[], $10::int[], $11::int[], $12::int[], $13::int[], 
	$14::int[], $15::int[], $16::int[], $17::int[], $18::int[], $19::int[], $20::int[],
	$21::int[], $22::int[], $23::text[], $24::int[]))
`

var TruncateLastperf = `
TRUNCATE lastperf
`

var InsertLastperf = `
INSERT INTO lastperf
(select * from unnest($1::int[], $2::text[], $3::int[], $4::int[], $5::int[], $6::int[], $7::int[], 
	$8::int[], $9::int[], $10::int[], $11::int[], $12::int[], $13::text[], $14::int[], $15::text[], $16::int[], $17::text[],
	$18::int[], $19::int[], $20::int[], $21::int[], $22::int[], $23::int[], $24::int[], $25::int[], $26::int[], 
	$27::int[], $28::int[], $29::int[], $30::int[], $31::int[], $32::int[], $33::int[], $34::int[], $35::int[], $36::int[], $37::int[]))
`

var TruncateLastrealtimeperf = `
TRUNCATE lastrealtimeperf
`

var InsertLastrealtimeperf = `
INSERT INTO lastrealtimeperf
(select * from unnest($1::int[], $2::int[], $3::text[], $4::int[], $5::int[], $6::int[], $7::int[], $8::int[], 
	$9::int[], $10::int[], $11::int[], $12::int[], $13::int[], $14::int[], $15::int[], $16::text[], $17::text[],
	$18::int[], $19::int[], $20::text[], $21::text[], $22::int[], $23::int[], $24::int[], $25::int[], $26::int[], 
	$27::int[], $28::int[]))
`

var InsertRealtimeperfPg = `
INSERT INTO %s
(select * from unnest($1::int[],$2::int[],$3::int[],$4::int[],$5::int[],$6::int[],$7::int[],$8::int[],$9::int[],$10::int[],
	$11::int[],$12::int[],$13::int[],$14::int[],$15::int[],$16::int[],$17::int[],$18::int[],$19::int[],$20::int[],
	$21::int[],$22::int[],$23::int[],$24::int[],$25::int[],$26::int[],$27::int[],$28::int[],$29::int[],$30::int[],
	$31::int[],$32::int[],$33::int[],$34::int[],$35::int[],$36::int[],$37::int[],$38::int[],$39::int[],$40::int[],
	$41::int[],$42::int[],$43::int[],$44::int[],$45::int[],$46::int[],$47::int[],$48::int[],$49::int[],$50::int[],
	$51::int[],$52::int[],$53::int[],$54::int[],$55::int[],$56::int[],$57::int[],$58::int[],$59::int[]
	))
`

var InsertRealtimeperfTs = `
INSERT INTO %s
(select * from unnest($1::timestamptz[],$2::int[],$3::int[],$4::int[],$5::int[],$6::int[],$7::int[],$8::int[],$9::int[],$10::int[],
	$11::int[],$12::int[],$13::int[],$14::int[],$15::int[],$16::int[],$17::int[],$18::int[],$19::int[],$20::int[],
	$21::int[],$22::int[],$23::int[],$24::int[],$25::int[],$26::int[],$27::int[],$28::int[],$29::int[],$30::int[],
	$31::int[],$32::int[],$33::int[],$34::int[],$35::int[],$36::int[],$37::int[],$38::int[],$39::int[],$40::int[],
	$41::int[],$42::int[],$43::int[],$44::int[],$45::int[],$46::int[],$47::int[],$48::int[],$49::int[],$50::int[],
	$51::int[],$52::int[],$53::int[],$54::int[],$55::int[],$56::int[],$57::int[],$58::int[],$59::int[]
	))
`
