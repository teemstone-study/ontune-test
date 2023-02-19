package data

var UpdateOntuneinfoStmt = `
UPDATE ontuneinfo set _time=$1
`

var UpdateAgentinfoReset = `
UPDATE agentinfo 
   set _connected=1, _updatedtime=$1
 where _enabled=1 
   and _connected=0
`

var UpdateAgentinfoState = `
UPDATE agentinfo 
   set _connected=%d, _updatedtime=%d 
 where _agentid in (%s) and _enabled=1 
   and _connected=%d
   and _updatedtime<%d
`
