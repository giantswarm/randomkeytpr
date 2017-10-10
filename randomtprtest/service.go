package randomkeytprtest

import "github.com/giantswarm/randomkeytpr"

type Service struct {
}

func NewService() Service {
	return Service{}
}

func (s Service) SearchKeys(clusterID string) (randomkeytpr.AssetsBundle, error) {
	return randomkeytpr.AssetsBundle{}, nil
}

func (s *Service) SearchKeysForComponent(clusterID, componentName string) (randomkeytpr.AssetsBundle, error) {
	return randomkeytpr.AssetsBundle{}, nil
}
