package resourcemanager

type ClusterInformationWrapper struct {
	ClusterInformation `json:"clusterInfo"`
}

type ClusterInformation struct {
	ID                            int    `json:"id"`
	StartedOn                     int    `json:"startedOn"`
	State                         string `json:"state"`
	HaState                       string `json:"haState"`
	ResourceManagerVersion        string `json:"resourceManagerVersion"`
	ResourceManagerBuildVersion   string `json:"ResourceManagerBuildVersion"`
	ResourceManagerVersionBuiltOn string `json:"ResourceManagerVersionBuiltOn"`
	HadoopVersion                 string `json:"hadoopVersion"`
	HadoopBuildVersion            string `json:"hadoopBuildVersion"`
	HadoopVersionBuiltOn          string `json:"hadoopVersionBuiltOn"`
}

type ClusterMetricsWrapper struct {
	ClusterMetrics `json:"clusterMetrics"`
}

type ClusterMetrics struct {
	AppsSubmitted         int `json:"appsSubmitted"`
	AppsCompleted         int `json:"appsCompleted"`
	AppsPending           int `json:"appsPending"`
	AppsRunning           int `json:"appsRunning"`
	AppsFailed            int `json:"appsFailed"`
	AppsKilled            int `json:"appsKilled"`
	ReservedMB            int `json:"ReservedMB"`
	AvailableMB           int `json:"availableMB"`
	ReservedVirtualCores  int `json:"reservedVirtualCores"`
	AvailableVirtualCores int `json:"availableVirtualCores"`
	AllocatedVirtualCores int `json:"allocatedVirtualCores"`
	ContainersAllocated   int `json:"containersAllocated"`
	ContainersReserved    int `json:"containersReserved"`
	ContainersPending     int `json:"containersPending"`
	TotalMB               int `json:"totalMB"`
	TotalVirtualCores     int `json:"totalVirtualCores"`
	TotalNodes            int `json:"totalNodes"`
	LostNodes             int `json:"lostNodes"`
	UnhealthyNodes        int `json:"unhealthyNodes"`
	DecommissionedNodes   int `json:"decommissionedNodes"`
	RebootedNodes         int `json:"rebootedNodes"`
	ActiveNodes           int `json:"activeNodes"`
}
