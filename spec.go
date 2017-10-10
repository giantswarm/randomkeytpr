package randomkeytpr

type Searcher interface {
	SearchKeys(clusterID string) (AssetsBundle, error)
}

type Spec struct {
	ClusterComponent string `json:"clusterComponent" yaml:"clusterComponent"`
	ClusterID        string `json:"clusterID" yaml:"clusterID"`
}
