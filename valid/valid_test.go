package valid

import (
	"testing"
)

type Name string
type Email string
type Cash int

// User is user
type User struct {
	Name
	Email
	Cash
}

// Team is team
type Team struct {
	Boss       User
	Leader     *User
	Staffs     []User
	Assistants []*User
	Trainees   *[]User
	Interns    *[]*User
}

func (src Name) IsValid() bool {
	return src != ""
}

func (src Cash) IsValid() bool {
	return src >= 0
}

func TestIsValidCheck(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"string is always Valid", args{"string"}, true},
		{"int is always Valid", args{42}, true},
		{"Valid Name", args{Name("Alice")}, true},
		{"InValid Name", args{Name("")}, false},
		{"Valid Cash", args{Cash(83)}, true},
		{"Invalid Cash", args{Cash(-27)}, false},
		{"Valid User", args{User{Name: "Alice", Cash: 302}}, true},
		{"Invalid User", args{User{Name: "Bob", Cash: -76}}, false},
		{"Valid Users", args{[]User{{Name: "Alice", Cash: 302}}}, true},
		{"Invalid Users", args{[]User{{Name: "Bob", Cash: -76}}}, false},
		{"Valid Users Ptr", args{&[]User{{Name: "Alice", Cash: 302}}}, true},
		{"Invalid Users Ptr", args{&[]User{{Name: "Bob", Cash: -76}}}, false},
		{"Valid Team", args{
			Team{
				Boss: User{
					Name: "Alice",
					Cash: 76,
				},
				Leader: &User{
					Name: "Bob",
					Cash: 27,
				},
				Staffs: []User{
					{
						Name: "Carol",
						Cash: 38,
					},
				},
				Assistants: []*User{
					{
						Name: "Dennis",
						Cash: 82,
					},
				},
				Trainees: &[]User{
					{
						Name: "Eric",
						Cash: 35,
					},
				},
				Interns: &[]*User{
					{
						Name: "Fred",
						Cash: 12,
					},
				},
			},
		}, true},
		{"Invalid Team", args{
			Team{
				Boss: User{
					Name: "Alice",
					Cash: 76,
				},
				Leader: &User{
					Name: "Bob",
					Cash: 27,
				},
				Staffs: []User{
					{
						Name: "Carol",
						Cash: 38,
					},
				},
				Assistants: []*User{
					{
						Name: "Dennis",
						Cash: 82,
					},
				},
				Trainees: &[]User{
					{
						Name: "Eric",
						Cash: -35,
					},
				},
				Interns: &[]*User{
					{
						Name: "Fred",
						Cash: 12,
					},
				},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidCheck(tt.args.value); got != tt.want {
				t.Errorf("IsValidCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
