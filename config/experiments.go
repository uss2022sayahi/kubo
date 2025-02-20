package config

type Experiments struct {
	FilestoreEnabled     bool
	UrlstoreEnabled      bool
	ShardingEnabled      bool `json:",omitempty"` // deprecated by autosharding: https://github.com/uss2022sayahi/kubo/pull/8527
	GraphsyncEnabled     bool
	Libp2pStreamMounting bool
	P2pHttpProxy         bool
	StrategicProviding   bool
	AcceleratedDHTClient bool
}
