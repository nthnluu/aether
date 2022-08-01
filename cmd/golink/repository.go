package golink

type Repository interface {
	LookupBySlug(slug string) (string, error)
	CreateLink(slug string, destinationUrl string) error
}

type repository struct {
	slugMapping map[string]string
}

func (r *repository) LookupBySlug(slug string) (string, error) {
	destinationUrl, _ := r.slugMapping[slug]
	return destinationUrl, nil
}

func (r *repository) CreateLink(slug string, destinationUrl string) error {
	r.slugMapping[slug] = destinationUrl
	return nil
}

func NewRepository() *repository {
	return &repository{slugMapping: make(map[string]string)}
}
