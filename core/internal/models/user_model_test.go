package models

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// var UserTester = struct {
// 	user     *User
// 	userRepo *UserRepository
// }{
// 	user: &User{
// 		Type:       "individual",
// 		FirstName:  "First",
// 		MiddleName: "Middle",
// 		LastName:   "Last",
// 		MobileNum:  "0912XXXXXXX",
// 		WalletID:   "",
// 	},
// }

// func TestUserModel(t *testing.T) {
// 	// Setup dependency can be replaced with mocked db.DBClient
// 	UserTester.userRepo = NewUserRepository(Tester.dbClient)

// 	t.Run("Test create user", func(t *testing.T) {
// 		// Create separate connection for the purpose of assertions
// 		db, close := UserTester.userRepo.Connect()
// 		defer close()

// 		assert.True(t, db.NewRecord(UserTester.user))
// 		UserTester.userRepo.Create(UserTester.user)
// 		assert.False(t, db.NewRecord(UserTester.user))
// 	})
// }
