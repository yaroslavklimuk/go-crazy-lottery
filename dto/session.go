package dto

type (
	Session interface {
		GetToken() string
		GetUserId() int64
		GetExpiredAt() int64
	}

	sessionImpl struct {
		Token     string
		UserId    int64
		ExpiredAt int64
	}
)

func (s sessionImpl) GetToken() string {
	return s.Token
}

func (s sessionImpl) GetUserId() int64 {
	return s.UserId
}

func (s sessionImpl) GetExpiredAt() int64 {
	return s.ExpiredAt
}

func NewSession(token string, userId int64, expiredAt int64) Session {
	return &sessionImpl{
		Token:     token,
		UserId:    userId,
		ExpiredAt: expiredAt,
	}
}
