package data

var DeleteAgentinfoDummy = `
DELETE FROM agentinfo where _agentname like 'Dummy%';
`

var DeleteHostinfoDummy = `
DELETE FROM hostinfo where _hostname like 'Dummy%';
`

var TruncateLastperf = `
TRUNCATE lastperf
`

var TruncateLastrealtimeperf = `
TRUNCATE lastrealtimeperf
`

var TruncateRef = `
TRUNCATE %s
`
