package repositories

type GenderRepository interface {
	GetAll() []string
}

type genderRepository struct {
	genderList []string
}

func NewGenderRepository() GenderRepository {
	repository := genderRepository{}
	repository.setUpData()

	return repository
}

func (r genderRepository) GetAll() []string {
	return r.genderList
}

func (r *genderRepository) setUpData() {
	r.genderList = []string{"Female", "Male", "Others"}
}
