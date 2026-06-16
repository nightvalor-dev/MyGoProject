package user

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(req CreateUserRequest) error {
	user := User{
		Username:     req.Username,
		Email:        req.Email,
		Phone:        req.Phone,
		Bio:          req.Bio,
		PasswordHash: req.Password,
		AssignedRole: Role(req.Role),
	}
	return s.repo.Create(user)
}

func (s *UserService) GetAll() ([]UserResponse, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var result []UserResponse
	for _, u := range users {
		result = append(result, toUserResponse(u))
	}
	return result, nil
}

func (s *UserService) GetById(id int) (UserResponse, error) {
	u, err := s.repo.GetById(id)
	if err != nil {
		return UserResponse{}, err
	}
	return toUserResponse(u), nil
}

func (s *UserService) Update(id int, req UpdateUserRequest) error {
	user := User{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Bio:      req.Bio,
	}
	return s.repo.Update(id, user)
}

func (s *UserService) Delete(id int) error {
	return s.repo.Delete(id)
}
