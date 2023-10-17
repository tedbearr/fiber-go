package service

type AuthService interface {
	CheckUser()
	GenerateAccessToken()
	GenerateRefreshToken()
	SaveUser()
	HashPassword()
	ComparePassword()
}
