package post

type Repository interface {
	GetList() []Post
	GetById(id string) Post
}

type repository struct {
	Ds DataSource
}

func (r *repository) GetList() []Post {
	return r.Ds.GetList()
}

func (r *repository) GetById(id string) Post {
	return r.Ds.GetById(id)
}

func NewRepository(ds DataSource) Repository {
	return &repository{
		Ds: ds,
	}
}
