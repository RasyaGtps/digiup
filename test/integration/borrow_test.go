package integration_test

import (
	"base-gin/domain/dao"
	"base-gin/domain/dto"
	"base-gin/server"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func ptrToTime(t time.Time) *time.Time {
	return &t
}

func TestBorrowing_Create_Success(t *testing.T) {
	params := dto.BorrowingDTO{
		BorrowDate: time.Now(),
		BookID:     1,
		PersonID:   1,
	}

	w := doTest(
		"POST",
		server.RootBorrowing,
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 201, w.Code)
}

func TestBorrowing_Update_Success(t *testing.T) {
    // Setup: Pastikan menggunakan data yang sudah ada di database
    b := dao.Borrowing{
        BorrowDate: time.Now(),
        BookID:     1,  // Gunakan ID book yang sudah ada
        PersonID:   1,  // Gunakan ID person yang sudah ada
    }
    err := borrowingRepo.Create(&b)
    assert.Nil(t, err)
    assert.NotZero(t, b.ID)

    // Siapkan data update
    returnDate := time.Now().AddDate(0, 0, 14)
    params := dto.BorrowingUpdate{
        ID:         b.ID,
        ReturnDate: &returnDate,
    }

    // Kirim request update
    w := doTest(
        "PUT",
        fmt.Sprintf("%s/%d", server.RootBorrowing, b.ID),
        params,
        createAuthAccessToken(dummyAdmin.Account.Username),
    )

    // Debug jika error
    if w.Code != 200 {
        t.Logf("Response Body: %s", w.Body.String())
    }

    assert.Equal(t, 200, w.Code)

    // Verifikasi update
    item, err := borrowingRepo.GetByID(b.ID)
    assert.Nil(t, err)
    
    // Bandingkan dengan presisi detik
    if item.ReturnDate != nil {
        assert.Equal(t, 
            returnDate.Unix(), 
            item.ReturnDate.Unix(), 
            "Return dates should match",
        )
    } else {
        t.Fatalf("Return date should not be nil")
    }
}

func TestBorrowing_Delete_Success(t *testing.T) {
    // Setup: Pastikan menggunakan data yang sudah ada di database
    b := dao.Borrowing{
        BorrowDate: time.Now(),
        BookID:     1,  // Gunakan ID book yang sudah ada
        PersonID:   1,  // Gunakan ID person yang sudah ada
    }
    err := borrowingRepo.Create(&b)
    assert.Nil(t, err)
    assert.NotZero(t, b.ID)

    // Kirim request DELETE
    w := doTest(
        "DELETE",
        fmt.Sprintf("%s/%d", server.RootBorrowing, b.ID),
        nil,
        createAuthAccessToken(dummyAdmin.Account.Username),
    )

    // Debug jika error
    if w.Code != 200 {
        t.Logf("Response Body: %s", w.Body.String())
    }

    assert.Equal(t, 200, w.Code)

    // Verifikasi borrowing sudah terhapus
    _, err = borrowingRepo.GetByID(b.ID)
    assert.NotNil(t, err, "Borrowing should be deleted")
}
func TestBorrowing_GetList_Success(t *testing.T) {
	b1 := dao.Borrowing{
		BorrowDate: time.Now(),
		BookID:     1,
		PersonID:   1,
	}
	_ = borrowingRepo.Create(&b1)

	w := doTest(
		"GET",
		server.RootBorrowing,
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)
}

func TestBorrowing_GetDetail_Success(t *testing.T) {
	b := dao.Borrowing{
		BorrowDate: time.Now(),
		BookID:     1,
		PersonID:   1,
	}
	_ = borrowingRepo.Create(&b)

	w := doTest(
		"GET",
		fmt.Sprintf("%s/%d", server.RootBorrowing, b.ID),
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)
}
