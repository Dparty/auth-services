package authservices

import (
	"github.com/Dparty/common/utils/random"
	"github.com/Dparty/dao/verification"
)

var verificationCodeServices *VerificationCodeService

func GetVerificationCodeService() *VerificationCodeService {
	if verificationCodeServices == nil {
		verificationCodeServices = NewVerificationCodeService()
	}
	return verificationCodeServices
}

func NewVerificationCodeService() *VerificationCodeService {
	return &VerificationCodeService{verificationCodeRepository: verification.GetVerificationCodeRepository()}
}

type VerificationCodeService struct {
	verificationCodeRepository *verification.VerificationCodeRepository
}

const VerificationCodeDigits = 6

func (s *VerificationCodeService) CreateEmailVerificationCode(purpose, email string) (string, error) {
	code := random.RandomNumberString(VerificationCodeDigits)
	verificationCode := verification.VerificationCode{
		Purpose: &purpose,
		Email:   &email,
		Code:    code,
	}
	s.verificationCodeRepository.Create(&verificationCode)
	return code, nil
}

func (s *VerificationCodeService) CreatePhoneVerificationCode(purpose, areaCode, phone string) (string, error) {
	code := random.RandomNumberString(VerificationCodeDigits)
	verificationCode := verification.VerificationCode{
		Purpose:     &purpose,
		AreaCode:    &areaCode,
		PhoneNumber: &phone,
		Code:        code,
	}
	s.verificationCodeRepository.Create(&verificationCode)
	return code, nil
}

func (s *VerificationCodeService) VerifyEmailVerificationCode(email, code string) (bool, error) {
	// TODO: verify email code
	return true, nil
}

func (s *VerificationCodeService) VerifyPhoneVerificationCode(areaCode, phone, code string) (bool, error) {
	// TODO: verify phone code
	return true, nil
}
