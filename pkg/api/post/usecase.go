package post

type UseCase interface {
	GetPostList() []Post
	GetPostById(id string) Post
}

type useCase struct {
	Repo Repository
}

func (u *useCase) GetPostList() []Post {
	return u.Repo.GetList()
}

func (u *useCase) GetPostById(id string) Post {
	return u.Repo.GetById(id)
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}
