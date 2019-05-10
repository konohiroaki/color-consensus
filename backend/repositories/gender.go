package repositories

type GenderRepository interface {
	GetAll() []string
	IsPresent(string) bool
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

func (r genderRepository) IsPresent(gender string) bool {
	for _, v := range r.genderList {
		if gender == v {
			return true
		}
	}
	return false
}

func (r *genderRepository) setUpData() {
	r.genderList = []string{"Female", "Male", "Others"}
}
