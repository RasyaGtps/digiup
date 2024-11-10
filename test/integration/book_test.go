package integration_test

import (
	"base-gin/domain/dao"
	"base-gin/domain/dto"
	"base-gin/server"
	"base-gin/util"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ptrToString(s string) *string {
	return &s
}

func TestBook_Create_Success(t *testing.T) {
	// Membuat data Publisher dan Author jika belum ada
	publisher := dao.Publisher{Name: "Sample Publisher"}
	author := dao.Author{FullName: "Sample Author"}
	_ = publisherRepo.Create(&publisher)
	_ = authorRepo.Create(&author)

	// Gunakan ID dari Publisher dan Author yang baru dibuat
	params := dto.BookDTO{
		Title:       util.RandomStringAlpha(6),
		Subtitle:    ptrToString(util.RandomStringAlpha(10)),
		PublisherID: publisher.ID,
		AuthorID:    author.ID,
	}

	// Lakukan pengujian POST untuk membuat data buku
	w := doTest(
		"POST",
		server.RootBook,
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 201, w.Code)
}


func TestBook_Update_Success(t *testing.T) {
    // 1. Buat Publisher
    publisher := dao.Publisher{
        Name: "Test Publisher " + util.RandomStringAlpha(4),
        City: "Test City",
    }
    err := publisherRepo.Create(&publisher)
    assert.Nil(t, err)
    assert.NotZero(t, publisher.ID)

    // 2. Buat Author
    author := dao.Author{
        FullName: "Test Author " + util.RandomStringAlpha(4),
        Gender:   "m",
    }
    err = authorRepo.Create(&author)
    assert.Nil(t, err)
    assert.NotZero(t, author.ID)

    // 3. Buat buku awal
    initialBook := dao.Book{
        Title:       util.RandomStringAlpha(6),
        Subtitle:    ptrToString(util.RandomStringAlpha(8)),
        PublisherID: publisher.ID,
        AuthorID:    author.ID,
    }
    err = bookRepo.Create(&initialBook)
    assert.Nil(t, err)
    assert.NotZero(t, initialBook.ID)

    // 4. Param update
    updateParams := dto.BookUpdate{
        Title:       util.RandomStringAlpha(7),
        Subtitle:    ptrToString(util.RandomStringAlpha(10)),
        PublisherID: publisher.ID,
        AuthorID:    author.ID,
    }

    // 5. Kirim request PUT
    w := doTest(
        "PUT",
        fmt.Sprintf("%s/%d", server.RootBook, initialBook.ID),
        updateParams,
        createAuthAccessToken(dummyAdmin.Account.Username),
    )

    assert.Equal(t, 200, w.Code)

    // 6. Verifikasi update
    updatedBook, err := bookRepo.GetByID(initialBook.ID)
    assert.Nil(t, err)
    assert.Equal(t, updateParams.Title, updatedBook.Title)
    assert.Equal(t, updateParams.Subtitle, updatedBook.Subtitle)
    assert.Equal(t, updateParams.PublisherID, updatedBook.PublisherID)
    assert.Equal(t, updateParams.AuthorID, updatedBook.AuthorID)
}

func TestBook_Delete_Success(t *testing.T) {
    // 1. Buat Publisher
    publisher := dao.Publisher{
        Name: "Test Publisher " + util.RandomStringAlpha(4),
        City: "Test City",
    }
    err := publisherRepo.Create(&publisher)
    assert.Nil(t, err)

    // 2. Buat Author
    author := dao.Author{
        FullName: "Test Author " + util.RandomStringAlpha(4),
        Gender:   "m",
    }
    err = authorRepo.Create(&author)
    assert.Nil(t, err)

    // 3. Buat buku yang akan dihapus
    b := dao.Book{
        Title:       util.RandomStringAlpha(6),
        Subtitle:    ptrToString(util.RandomStringAlpha(8)),
        PublisherID: publisher.ID,
        AuthorID:    author.ID,
    }
    err = bookRepo.Create(&b)
    assert.Nil(t, err)
    assert.NotZero(t, b.ID)

    // 4. Kirim request DELETE
    w := doTest(
        "DELETE",
        fmt.Sprintf("%s/%d", server.RootBook, b.ID),
        nil,
        createAuthAccessToken(dummyAdmin.Account.Username),
    )

    assert.Equal(t, 200, w.Code)

    // 5. Verifikasi book sudah terhapus
    deletedBook, err := bookRepo.GetByIDUnscoped(b.ID)
    assert.Nil(t, err)
    assert.NotNil(t, deletedBook.DeletedAt) // Verifikasi bahwa DeletedAt tidak nil
}
func TestBook_GetList_Success(t *testing.T) {
	b1 := dao.Book{
		Title:    util.RandomStringAlpha(6),
		Subtitle: ptrToString(util.RandomStringAlpha(8)),
	}
	_ = bookRepo.Create(&b1)

	w := doTest(
		"GET",
		server.RootBook,
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)
}

func TestBook_GetDetail_Success(t *testing.T) {
    // 1. Buat Publisher
    publisher := dao.Publisher{
        Name: "Test Publisher " + util.RandomStringAlpha(4),
        City: "Test City",
    }
    err := publisherRepo.Create(&publisher)
    assert.Nil(t, err)

    // 2. Buat Author
    author := dao.Author{
        FullName: "Test Author " + util.RandomStringAlpha(4),
        Gender:   "m",
    }
    err = authorRepo.Create(&author)
    assert.Nil(t, err)

    // 3. Buat buku
    b := dao.Book{
        Title:       util.RandomStringAlpha(6),
        Subtitle:    ptrToString(util.RandomStringAlpha(8)),
        PublisherID: publisher.ID,  // Tambahkan PublisherID
        AuthorID:    author.ID,     // Tambahkan AuthorID
    }
    err = bookRepo.Create(&b)
    assert.Nil(t, err)
    assert.NotZero(t, b.ID)

    // 4. Get detail buku
    w := doTest(
        "GET",
        fmt.Sprintf("%s/%d", server.RootBook, b.ID),
        nil,
        "",
    )

    // Debug response jika error
    if w.Code != 200 {
        t.Logf("Response Body: %s", w.Body.String())
    }

    assert.Equal(t, 200, w.Code)
}