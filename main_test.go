package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {

	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()
	// настройте подключение к БД

	clientID := 1

	// напиши тест здесь
	cl, err := selectClient(db, clientID)

	require.NoError(t, err)
	assert.Equal(t, clientID, cl.ID)
	assert.NotEmpty(t, cl.Birthday)
	assert.NotEmpty(t, cl.Email)
	assert.NotEmpty(t, cl.FIO)
	assert.NotEmpty(t, cl.Login)

}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()
	// настройте подключение к БД

	clientID := -1

	// напиши тест здесь
	cl, err := selectClient(db, clientID)
	require.Equal(t, sql.ErrNoRows, err)

	assert.Empty(t, cl.Birthday)
	assert.Empty(t, cl.Email)
	assert.Empty(t, cl.FIO)
	assert.Empty(t, cl.Login)
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()
	// настройте подключение к БД

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	id, err := insertClient(db, cl)
	cl.ID = id
	require.NotEmpty(t, cl.ID)
	require.NoError(t, err)

	clSel, err := selectClient(db, cl.ID)
	require.NoError(t, err)

	assert.Equal(t, cl.ID, clSel.ID)
	assert.Equal(t, cl.Birthday, clSel.Birthday)
	assert.Equal(t, cl.Email, clSel.Email)
	assert.Equal(t, cl.Login, clSel.Login)
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()
	// настройте подключение к БД

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	id, err := insertClient(db, cl)

	require.NotEmpty(t, id)
	require.NoError(t, err)

	err = deleteClient(db, id)
	require.NoError(t, err)

	_, err = selectClient(db, id)
	require.Equal(t, sql.ErrNoRows, err)

}
