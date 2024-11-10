package integration_test

import (
	"base-gin/domain/dao"
	"base-gin/domain/dto"
	"base-gin/server"
	"base-gin/util"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func stringPtr(s string) *string {
	return &s
}

func timePtr(t time.Time) *time.Time {
	return &t
}

func TestAuthor_Create_Success(t *testing.T) {
	params := dto.AuthorDTO{
		FullName:  util.RandomStringAlpha(8),
		Gender:    stringPtr("m"),
		BirthDate: timePtr(time.Now().AddDate(-30, 0, 0)),
	}

	w := doTest(
		"POST",
		server.RootAuthor,
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 201, w.Code)
}

func TestAuthor_Update_Success(t *testing.T) {
	a := dao.Author{
		FullName:  util.RandomStringAlpha(8),
		Gender:    *stringPtr("m"),
		BirthDate: *timePtr(time.Now().AddDate(-30, 0, 0)),
	}
	_ = authorRepo.Create(&a)

	params := dto.AuthorUpdate{
		FullName:  util.RandomStringAlpha(10),
		Gender:    stringPtr("f"),
		BirthDate: timePtr(time.Now().AddDate(-28, 0, 0)),
	}

	w := doTest(
		"PUT",
		fmt.Sprintf("%s/%d", server.RootAuthor, a.ID),
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 200, w.Code)

	item, _ := authorRepo.GetByID(a.ID)
	assert.Equal(t, params.FullName, item.FullName)
	assert.Equal(t, *params.Gender, item.Gender)
	assert.WithinDuration(t, *params.BirthDate, item.BirthDate, time.Second)

}

func TestAuthor_Delete_Success(t *testing.T) {
	a := dao.Author{
		FullName:  util.RandomStringAlpha(8),
		Gender:    *stringPtr("f"),
		BirthDate: *timePtr(time.Now().AddDate(-25, 0, 0)),
	}
	_ = authorRepo.Create(&a)

	w := doTest(
		"DELETE",
		fmt.Sprintf("%s/%d", server.RootAuthor, a.ID),
		nil,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 200, w.Code)
}

func TestAuthor_GetList_Success(t *testing.T) {
	a := dao.Author{
		FullName:  util.RandomStringAlpha(8),
		Gender:    *stringPtr("m"),
		BirthDate: *timePtr(time.Now().AddDate(-25, 0, 0)),
	}
	_ = authorRepo.Create(&a)

	w := doTest(
		"GET",
		server.RootAuthor,
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)
}

func TestAuthor_GetDetail_Success(t *testing.T) {
	a := dao.Author{
		FullName:  util.RandomStringAlpha(8),
		Gender:    *stringPtr("f"),
		BirthDate: *timePtr(time.Now().AddDate(-20, 0, 0)),
	}
	_ = authorRepo.Create(&a)

	w := doTest(
		"GET",
		fmt.Sprintf("%s/%d", server.RootAuthor, a.ID),
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)
}
