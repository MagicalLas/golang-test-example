package basic

// User is simeple user.
type User struct {
	name  string
	age   uint
	point uint
}

// NewUser create user.
func NewUser(name string, age uint) *User {
	return &User{name, age, 0}
}

// IsAdult return user is adult or not.
func (u *User) IsAdult() bool {
	return u.age > 20
}

// Name get user's name.
func (u *User) Name() string {
	return u.name
}

// GetPoint user get more point.
func (u *User) GetPoint(point uint) {
	u.point += point
}

// Purchase user spend some point.
func (u *User) Purchase(item *Item) error {
	u.point -= item.cost
}

// Item is item that can be sale.
type Item struct {
	cost uint
}

// NewItem create new item.
func NewItem(cost uint) *Item {
	return &Item{cost}
}
