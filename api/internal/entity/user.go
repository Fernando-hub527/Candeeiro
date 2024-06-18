package entity

type User struct {
	userName string
	plans    []Plan
}

func (u *User) GetUserName() string {
	return u.userName
}

func (u *User) GetPlans() []Plan {
	return u.plans
}
