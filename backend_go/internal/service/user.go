package service

import (
	"backend_go/internal/entity"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo entity.UserRepository
}

func NewUserService(userRepo entity.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetAll() (*[]entity.User, error) {
	usersDB, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	} else {
		for i := 0; i < len(*usersDB); i++ {
			(*usersDB)[i].OmitPassword()
		}
		return usersDB, nil
	}
}

func (s *UserService) Get(id uint) (*entity.User, error) {
	userDB, err := s.userRepo.Get(id)
	if err != nil {
		return nil, err
	} else {
		userDB.OmitPassword()
		return userDB, nil
	}
}

func (s *UserService) Update(user *entity.User) error {
	userDB, err := s.userRepo.Get(user.ID)
	if err != nil {
		return err
	}

	var newUserPswdHash string
	if comparePaaswordWithHash(user.Password, userDB.Password) != nil {
		newUserPswdHash, err = generatePasswordHash(user.Password)
		if err != nil {
			return err
		}
	}

	user.Password = newUserPswdHash
	err = s.userRepo.Update(user)
	return err
}

func comparePaaswordWithHash(pswdFromInput, pswdHashFromDB string) error {
	err := bcrypt.CompareHashAndPassword([]byte(pswdHashFromDB), []byte(pswdFromInput))
	return err
}

func generatePasswordHash(pswd string) (string, error) {
	pswdHash, err := bcrypt.GenerateFromPassword([]byte(pswd), bcrypt.DefaultCost)
	return string(pswdHash), err
}

func (s *UserService) Delete(user *entity.User) error {
	err := s.userRepo.Delete(user.ID)
	return err
}

func (s *UserService) Register(userReg *entity.UserRegister) error {
	if userReg.Email == "" {
		return entity.ErrInvalidEmail
	}

	if len(userReg.Password) < 8 {
		return entity.ErrInvalidPassword
	}

	pswdHash, err := generatePasswordHash(userReg.Password)
	if err != nil {
		return err
	}

	userReg.Password = pswdHash
	user := entity.User{UserRegister: *userReg}
	err = s.userRepo.Create(&user)
	return err
}

func (s *UserService) Login(userLogin *entity.UserLogin) (uint, error) {
	userDB, err := s.userRepo.GetByEmail(userLogin.Email)
	if err != nil {
		return 0, err
	}

	err = comparePaaswordWithHash(userLogin.Password, userDB.Password)
	if err != nil {
		return 0, err
	}

	return userDB.ID, nil
}
