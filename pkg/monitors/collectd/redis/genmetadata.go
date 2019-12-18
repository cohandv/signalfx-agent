// Code generated by monitor-code-gen. DO NOT EDIT.

package redis

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "collectd/redis"

var groupSet = map[string]bool{}

const (
	bytesUsedMemory                 = "bytes.used_memory"
	bytesUsedMemoryLua              = "bytes.used_memory_lua"
	bytesUsedMemoryPeak             = "bytes.used_memory_peak"
	bytesUsedMemoryRss              = "bytes.used_memory_rss"
	counterCommandsProcessed        = "counter.commands_processed"
	counterConnectionsReceived      = "counter.connections_received"
	counterEvictedKeys              = "counter.evicted_keys"
	counterExpiredKeys              = "counter.expired_keys"
	counterLruClock                 = "counter.lru_clock"
	counterRejectedConnections      = "counter.rejected_connections"
	counterTotalNetInputBytes       = "counter.total_net_input_bytes"
	counterTotalNetOutputBytes      = "counter.total_net_output_bytes"
	counterUsedCPUSys               = "counter.used_cpu_sys"
	counterUsedCPUSysChildren       = "counter.used_cpu_sys_children"
	counterUsedCPUUser              = "counter.used_cpu_user"
	counterUsedCPUUserChildren      = "counter.used_cpu_user_children"
	deriveKeyspaceHits              = "derive.keyspace_hits"
	deriveKeyspaceMisses            = "derive.keyspace_misses"
	gaugeBlockedClients             = "gauge.blocked_clients"
	gaugeChangesSinceLastSave       = "gauge.changes_since_last_save"
	gaugeClientBiggestInputBuf      = "gauge.client_biggest_input_buf"
	gaugeClientLongestOutputList    = "gauge.client_longest_output_list"
	gaugeConnectedClients           = "gauge.connected_clients"
	gaugeConnectedSlaves            = "gauge.connected_slaves"
	gaugeDb0AvgTTL                  = "gauge.db0_avg_ttl"
	gaugeDb0Expires                 = "gauge.db0_expires"
	gaugeDb0Keys                    = "gauge.db0_keys"
	gaugeInstantaneousOpsPerSec     = "gauge.instantaneous_ops_per_sec"
	gaugeKeyLlen                    = "gauge.key_llen"
	gaugeLatestForkUsec             = "gauge.latest_fork_usec"
	gaugeMasterLastIoSecondsAgo     = "gauge.master_last_io_seconds_ago"
	gaugeMasterReplOffset           = "gauge.master_repl_offset"
	gaugeMemFragmentationRatio      = "gauge.mem_fragmentation_ratio"
	gaugeRdbBgsaveInProgress        = "gauge.rdb_bgsave_in_progress"
	gaugeReplBacklogFirstByteOffset = "gauge.repl_backlog_first_byte_offset"
	gaugeSlaveReplOffset            = "gauge.slave_repl_offset"
	gaugeUptimeInDays               = "gauge.uptime_in_days"
	gaugeUptimeInSeconds            = "gauge.uptime_in_seconds"
)

var metricSet = map[string]monitors.MetricInfo{
	bytesUsedMemory:                 {Type: datapoint.Gauge},
	bytesUsedMemoryLua:              {Type: datapoint.Gauge},
	bytesUsedMemoryPeak:             {Type: datapoint.Gauge},
	bytesUsedMemoryRss:              {Type: datapoint.Gauge},
	counterCommandsProcessed:        {Type: datapoint.Counter},
	counterConnectionsReceived:      {Type: datapoint.Counter},
	counterEvictedKeys:              {Type: datapoint.Counter},
	counterExpiredKeys:              {Type: datapoint.Counter},
	counterLruClock:                 {Type: datapoint.Counter},
	counterRejectedConnections:      {Type: datapoint.Counter},
	counterTotalNetInputBytes:       {Type: datapoint.Counter},
	counterTotalNetOutputBytes:      {Type: datapoint.Counter},
	counterUsedCPUSys:               {Type: datapoint.Counter},
	counterUsedCPUSysChildren:       {Type: datapoint.Counter},
	counterUsedCPUUser:              {Type: datapoint.Counter},
	counterUsedCPUUserChildren:      {Type: datapoint.Counter},
	deriveKeyspaceHits:              {Type: datapoint.Counter},
	deriveKeyspaceMisses:            {Type: datapoint.Counter},
	gaugeBlockedClients:             {Type: datapoint.Gauge},
	gaugeChangesSinceLastSave:       {Type: datapoint.Gauge},
	gaugeClientBiggestInputBuf:      {Type: datapoint.Gauge},
	gaugeClientLongestOutputList:    {Type: datapoint.Gauge},
	gaugeConnectedClients:           {Type: datapoint.Gauge},
	gaugeConnectedSlaves:            {Type: datapoint.Gauge},
	gaugeDb0AvgTTL:                  {Type: datapoint.Gauge},
	gaugeDb0Expires:                 {Type: datapoint.Gauge},
	gaugeDb0Keys:                    {Type: datapoint.Gauge},
	gaugeInstantaneousOpsPerSec:     {Type: datapoint.Gauge},
	gaugeKeyLlen:                    {Type: datapoint.Gauge},
	gaugeLatestForkUsec:             {Type: datapoint.Gauge},
	gaugeMasterLastIoSecondsAgo:     {Type: datapoint.Gauge},
	gaugeMasterReplOffset:           {Type: datapoint.Gauge},
	gaugeMemFragmentationRatio:      {Type: datapoint.Gauge},
	gaugeRdbBgsaveInProgress:        {Type: datapoint.Gauge},
	gaugeReplBacklogFirstByteOffset: {Type: datapoint.Gauge},
	gaugeSlaveReplOffset:            {Type: datapoint.Gauge},
	gaugeUptimeInDays:               {Type: datapoint.Gauge},
	gaugeUptimeInSeconds:            {Type: datapoint.Gauge},
}

var defaultMetrics = map[string]bool{
	bytesUsedMemory:            true,
	bytesUsedMemoryRss:         true,
	counterCommandsProcessed:   true,
	counterEvictedKeys:         true,
	counterExpiredKeys:         true,
	counterRejectedConnections: true,
	counterTotalNetInputBytes:  true,
	counterTotalNetOutputBytes: true,
	counterUsedCPUSys:          true,
	counterUsedCPUUser:         true,
	deriveKeyspaceHits:         true,
	deriveKeyspaceMisses:       true,
	gaugeBlockedClients:        true,
	gaugeConnectedClients:      true,
	gaugeMasterReplOffset:      true,
	gaugeSlaveReplOffset:       true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:       "collectd/redis",
	DefaultMetrics:    defaultMetrics,
	Metrics:           metricSet,
	MetricsExhaustive: false,
	Groups:            groupSet,
	GroupMetricsMap:   groupMetricsMap,
	SendAll:           false,
}