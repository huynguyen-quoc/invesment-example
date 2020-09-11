package rredis

// OffsetType is a custom type to choose the appropriate offset
type OffsetType string

// KafkaConfig defines the basic config for sk2
type KafkaConfig struct {
	// Cluster and Stream Configurations used by both Producers & Consumers
	Brokers           []string `json:"brokers"`
	ClientID          string   `json:"clientID"`
	ClusterType       string   `json:"clusterType"`
	Enabled           bool     `json:"enabled"`
	KafkaVersion      string   `json:"kafkaVersion"`
	ShortName         string   `json:"shortName"`
	Stream            string   `json:"stream"`

	// Consumer Configurations
	ConsumerGroupID string     `json:"consumerGroupID"`
	OffsetType      OffsetType `json:"offsetType"`

	// Producer Configurations
	CompressionCodec string `json:"compressionCodec"`
	CompressionLevel int    `json:"compressionLevel"`
	EnableRetry      bool   `json:"enableRetry"`
	RequiredAcks     *int   `json:"requiredAcks,omitempty"`

}
