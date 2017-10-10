package randomkeytpr

// ClusterComponent represents the individual component of a k8s cluster, e.g.
// the API server, or etcd These are used when getting a secret from the k8s
// API, to identify the component the secret belongs to.
type ClusterComponent string

func (c ClusterComponent) String() string {
	return string(c)
}

// RandomKeyAssetType represents the type of Random Key asset, e.g. a encryption key.
// These are used when getting a secret from the k8s API, to
// identify the specific type of Random Key asset that is contained in the secret.
type RandomKeyAssetType string

func (t RandomKeyAssetType) String() string {
	return string(t)
}

// These constants are used to match each asset in the secret.
const (
	// EncryptionKey is the key for the kubernetes encryption.
	EncryptionKey RandomKeyAssetType = "encryption"
)

// These constants are used to match different components of the cluster when
// parsing a secret received from the API.
const (
	// APIComponent is the API server component.
	APIComponent ClusterComponent = "api-keys"
)

// These constants are used when filtering the secrets, to only retrieve the
// ones we are interested in.
const (
	// ComponentLabel is the label used in the secret to identify a cluster
	// component.
	ComponentLabel string = "clusterComponent"
	// ClusterIDLabel is the label used in the secret to identify a cluster.
	ClusterIDLabel string = "clusterID"
)

// AssetsBundleKey is a struct key for an AssetsBundle cfr.
// https://blog.golang.org/go-maps-in-action
type AssetsBundleKey struct {
	Component ClusterComponent
	Type      RandomKeyAssetType
}

// AssetsBundle is a structure that contains all the assets for all the
// components.
type AssetsBundle map[AssetsBundleKey][]byte

// ClusterComponents is a slice enumerating all the components that make up the
// cluster.
var ClusterComponents = []ClusterComponent{
	APIComponent,
}

// RandomKeyAssetTypes is a slice enumerating all the Random Key assets we need to boot the
// cluster.
var RandomKeyAssetTypes = []RandomKeyAssetType{EncryptionKey}

// ValidComponent looks for el among the components.
func ValidComponent(el ClusterComponent, components []ClusterComponent) bool {
	for _, v := range components {
		if el == v {
			return true
		}
	}
	return false
}

// CompactRandomKeyAssets is a struct used by operators to store stringified Random Key assets.
type CompactRandomKeyAssets struct {
	APIServerEncryptionKey string
}
