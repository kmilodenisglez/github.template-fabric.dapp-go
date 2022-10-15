package repo

import (
	"dapp/lib"
	"dapp/schema/dto"
	"dapp/service/utils"
	"fmt"
	"sync"
)

// region ======== SETUP =================================================================

type RepoUser struct {
	DBLocation string
}

var singletonRU *RepoUser

// using Go sync package to invoke a method exactly only once
var onceRU sync.Once

// endregion =============================================================================

func NewRepoUser(svcConf *utils.SvcConfig) *RepoUser {
	onceRU.Do(func() {
		singletonRU = &RepoUser{DBLocation: svcConf.StoreDBPath}

		// TODO: "fakeUsers" is only for demo purpose. Save users in In-memory.
		fakeUsers()
	})
	return singletonRU
}

// In-memory storage
// replace later with some db

var UsersById map[string]any

// GetUser get the user from the DB
func (r *RepoUser) GetUser(userID string) (dto.User, error) {

	if user, exists := UsersById[userID]; exists  {
		return user.(dto.User), nil
	}
	return dto.User{}, fmt.Errorf("user not exist")
}

// GetUsers return a list of dto.User
func (r *RepoUser) GetUsers() ([]any, error) {
	res := lib.MapToSliceOfValues(UsersById)
	return res, nil
}

func fakeUsers()  {
	if len(UsersById) != 0 {
		return
	}
	p1, _ := lib.Checksum("SHA256", []byte("password1"))
	p2, _ := lib.Checksum("SHA256", []byte("password2"))

	users := []dto.User{
		{
			Username:   "richard.sargon@meinermail.com",
			Passphrase: p1,
			FirstName:  "Richard",
			LastName:   "Sargon",
			Email:      "richard.sargon@meinermail.com",
		},
		{
			Username:   "tom.carter@meinermail.com",
			Passphrase: p2,
			FirstName:  "Tom",
			LastName:   "Carter",
			Email:      "tom.carter@meinermail.com",
		},
	}

	UsersById = make(map[string]any)
	for _, user := range users {
		UsersById[user.Email] = user
	}
}

//func fakeDrones() []dto.Drone {
//	uuid := "123e4567-e89b-12d3-a456-4266141740"
//	var drones = []dto.Drone{{
//		SerialNumber:    uuid + "10",
//		Model:           dto.Lightweight,
//		WeightLimit:     lib.CalculateDroneWeightLimit(dto.Lightweight),
//		BatteryCapacity: 25,
//		State:           dto.IDLE,
//	}}
//
//	var medications = []dto.Medication{{
//		Name:   gofakeit.Password(true, true, true, false, false, 12),
//		Weight: 700,
//		Code:   gofakeit.Password(false, true, true, false, false, 10),
//		Image:  base64.StdEncoding.EncodeToString([]byte("fake_image")),
//	}}
//
//	return drones
//}
