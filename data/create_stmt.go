package data

var TableinfoStmt = `
CREATE TABLE IF NOT EXISTS public.tableinfo (			
	_tablename varchar(64) NOT NULL PRIMARY KEY,
	_version integer NULL,		
	_createdtime int8 NULL,		
	_updatetime int8 NULL,		
	_durationmin integer NULL		
);			
CREATE INDEX tableinfo_idx ON public.tableinfo USING btree (_createdtime);			
`

var OntuneinfoStmt = `
CREATE TABLE IF NOT EXISTS public.ontuneinfo (			
	_id integer NOT NULL PRIMARY KEY,
	_majorversion integer NULL,		
	_minorversion integer NULL,		
	_releaseversion integer NULL,		
	_time %s NULL,		
	_processedpacket integer NULL,		
	_processedpackettime integer NULL,		
	_newdata integer NULL,		
	_inserteddata integer NULL,		
	_queueddata integer NULL,		
	_status text NULL,		
	_address text NULL,		
	_data text NULL,		
	_fdata integer NULL,		
	_eventqueue integer NULL,		
	_standardbias integer NULL,		
	_daylight integer NULL,		
	_daylightbias integer NULL,		
	_tablemode integer NULL,		
	_bias integer NULL,		
	_alterdateidtable int8 NULL,		
	_build integer NULL,		
	_alterdateagenttime int8 NULL,		
	_daylightstarttime int8 NULL,		
	_standardstarttime int8 NULL,		
	_alterdatepidtable int8 NULL,		
	_alterhourperftable int8 NULL		
);			
`

var AgentinfoStmt = `
CREATE TABLE IF NOT EXISTS public.agentinfo (			
	_agentid integer NOT NULL PRIMARY KEY,
	_agentname text NULL,		
	_enabled integer NULL,		
	_connected integer NULL,		
	_updated integer NULL,		
	_shorttermbasic integer NULL,
	_shorttermproc integer NULL,		
	_shorttermio integer NULL,		
	_shorttermcpu integer NULL,		
	_longtermbasic integer NULL,		
	_longtermproc integer NULL,		
	_longtermio integer NULL,		
	_longtermcpu integer NULL,		
	_model text NULL,		
	_serial text NULL,		
	_group text NULL,		
	_ipaddress text NULL,		
	_pscommand text NULL,		
	_logevent text NULL,		
	_processevent text NULL,		
	_timecheck integer NULL,		
	_disconnectedtime int8 NULL,		
	_skipdatatypes integer NULL,		
	_virbasicperf integer NULL,		
	_hypervisor integer NULL,		
	_serviceevent text NULL,		
	_installdate int8 NULL,		
	_lastconnectedtime int8 NULL		
);
`

var HostinfoStmt = `
CREATE TABLE IF NOT EXISTS public.hostinfo (			
	_agentid integer NOT NULL PRIMARY KEY,
	_hostname text NULL,		
	_hostnameext text NULL,
	_os text NULL,		
	_fw text NULL,		
	_agentversion text NULL,		
	_model text NULL,		
	_serial text NULL,		
	_processorcount integer NULL,		
	_processorclock integer NULL,		
	_memorysize integer NULL,		
	_swapsize integer NULL,		
	_poolid integer NULL,		
	_replication integer NULL,		
	_smt integer NULL,		
	_micropar integer NULL,		
	_capped integer NULL,		
	_ec integer NULL,		
	_virtualcpu integer NULL,		
	_weight integer NULL,		
	_cpupool integer NULL,		
	_ams integer NULL,		
	_allip text NULL,		
	_numanodecount integer NULL DEFAULT 0
);
`

var LastrealtimeperfStmt = `
CREATE UNLOGGED TABLE IF NOT EXISTS public.lastrealtimeperf (			
	_ontunetime int8 NOT NULL,		
	_agentid integer NOT NULL,		
	_hostname text NULL,		
	_user integer NULL,		
	_sys integer NULL,		
	_wait integer NULL,		
	_idle integer NULL,		
	_memoryused integer NULL,		
	_filecache integer NULL,		
	_memorysize integer NULL,		
	_avm integer NULL,		
	_swapused integer NULL,		
	_swapsize integer NULL,		
	_diskiorate integer NULL,		
	_networkiorate integer NULL,		
	_topproc text NULL,		
	_topuser text NULL,		
	_topproccount integer NULL,		
	_topcpu integer NULL,		
	_topdisk text NULL,		
	_topvg text NULL,		
	_topbusy integer NULL,		
	_maxcpu integer NULL,		
	_maxmem integer NULL,		
	_maxswap integer NULL,		
	_maxdisk integer NULL,		
	_diskiops integer NULL,		
	_networkiops integer NULL
);			
CREATE INDEX lastrealtimeperf_idx ON public.lastrealtimeperf USING btree (_agentid);
`

var LastperfStmt = `
CREATE UNLOGGED TABLE IF NOT EXISTS public.lastperf (			
	_ontunetime int8 NOT NULL,		
	_hostname varchar(255) NULL,		
	_user integer NULL,		
	_sys integer NULL,		
	_wait integer NULL,		
	_idle integer NULL,		
	_avm integer NULL,		
	_memoryused integer NULL,		
	_filecache integer NULL,		
	_swapused integer NULL,		
	_diskiorate integer NULL,		
	_networkiorate integer NULL,		
	_topproc text NULL,		
	_topcpu integer NULL,		
	_topdisk text NULL,		
	_topbusy integer NULL,		
	_filesystem text NULL,
	_runqueue integer NULL,
	_blockqueue integer NULL,
	_pagingspacein integer NULL,
	_pagingspaceout integer NULL,
	_filesystemin integer NULL,
	_filesystemout integer NULL,
	_memoryscan integer NULL,
	_memoryfreed integer NULL,
	_swapactive integer NULL,
	_fork integer NULL,
	_exec integer NULL,
	_interupt integer NULL,
	_systemcall integer NULL,
	_contextswitch integer NULL,
	_semaphore integer NULL,
	_msg integer NULL,
	_diskiops integer NULL,		
	_networkiops integer NULL,
	_networkreadrate integer NULL,
	_networkwriterate integer NULL
);			
CREATE INDEX IF NOT EXISTS lastperf_idx ON public.lastperf USING btree (_hostname);
`

var DeviceidStmt = `
CREATE TABLE IF NOT EXISTS public.deviceid (	
	_id serial NOT NULL PRIMARY KEY,
	_name text NULL
);	
`

var DescidStmt = `
CREATE TABLE IF NOT EXISTS public.descid (	
	_id serial NOT NULL PRIMARY KEY,
	_name text NULL
);	
`

var MetricPerfStmt = `
CREATE TABLE IF NOT EXISTS public.%[1]s (			
	_ontunetime	%[2]s NOT NULL,
	_agenttime	integer NULL,
	_agentid	integer NULL,
	_user	integer NULL,
	_sys	integer NULL,
	_wait	integer NULL,
	_idle	integer NULL,
	_processorcount	integer NULL,
	_runqueue	integer NULL,
	_blockqueue	integer NULL,
	_waitqueue	integer NULL,
	_pqueue	integer NULL,
	_pcrateuser	integer NULL,
	_pcratesys	integer NULL,
	_memorysize	integer NULL,
	_memoryused	integer NULL,
	_memorypinned	integer NULL,
	_memorysys	integer NULL,
	_memoryuser	integer NULL,
	_memorycache	integer NULL,
	_avm	integer NULL,
	_pagingspacein	integer NULL,
	_pagingspaceout	integer NULL,
	_filesystemin	integer NULL,
	_filesystemout	integer NULL,
	_memoryscan	integer NULL,
	_memoryfreed	integer NULL,
	_swapsize	integer NULL,
	_swapused	integer NULL,
	_swapactive	integer NULL,
	_fork	integer NULL,
	_exec	integer NULL,
	_interupt	integer NULL,
	_systemcall	integer NULL,
	_contextswitch	integer NULL,
	_semaphore	integer NULL,
	_msg	integer NULL,
	_diskreadwrite	integer NULL,
	_diskiops	integer NULL,
	_networkreadwrite	integer NULL,
	_networkiops	integer NULL,
	_topcommandid	integer NULL,
	_topcommandcount	integer NULL,
	_topuserid	integer NULL,
	_topcpu	integer NULL,
	_topdiskid	integer NULL,
	_topvgid	integer NULL,
	_topbusy	integer NULL,
	_maxpid	integer NULL,
	_threadcount	integer NULL,
	_pidcount	integer NULL,
	_linuxbuffer	integer NULL,
	_linuxcached	integer NULL,
	_linuxsrec	integer NULL,
	_memused_mb	integer NULL,
	_irq	integer NULL,
	_softirq	integer NULL,
	_swapused_mb	integer NULL,
	_dusm	integer NULL
);
CREATE INDEX IF NOT EXISTS %[1]s_idx ON public.%[1]s USING btree (_ontunetime, _agentid);
`

var MetricPerfHypertable = `
select create_hypertable('%s','_ontunetime', chunk_time_interval => interval '1 day');
`
