package basic

import "github.com/stretchr/testify/assert"

func TestCreateUser(t *testing.T) {
	user := NewUser("Las", 20)

	assert.NotNil(t, user)
}

func TestUser_GetPoint(t *testing.T) {
	// Given.
	user := NewUser("Las", 20)

	// When.
	user.GetPoint(100)
	
	// Then.
	// 여기에서는 user의 point가 private이기 때문에, package에 suffix로
	// `_test`가 붙지 않았습니다. 이렇게 할 경우 private한 method & func
	// & field에 접근이 가능합니다.
	assert.Equal(t, 100, user.point)
}