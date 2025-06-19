package main

import (
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	clientID := 1

	client, err := selectClient(db, clientID)
	require.NoError(t, err)
	assert.Equal(t, clientID, client.ID)
	assert.NotEmpty(t, client.FIO)
	assert.NotEmpty(t, client.Login)
	assert.NotEmpty(t, client.Email)
	assert.NotEmpty(t, client.Birthday)

}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	clientID := -1

	client, err := selectClient(db, clientID)
	require.Error(t, err)
	require.Equal(t, sql.ErrNoRows, err)
	assert.Empty(t, client.ID)
	assert.Empty(t, client.FIO)
	assert.Empty(t, client.Login)
	assert.Empty(t, client.Email)
	assert.Empty(t, client.Birthday)
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	cl.ID, err = insertClient(db, cl)
	require.NoError(t, err)
	require.NotEmpty(t, cl.ID)

	client, err := selectClient(db, cl.ID)
	require.NoError(t, err)
	assert.Equal(t, cl.FIO, client.FIO)
	assert.Equal(t, cl.Login, client.Login)
	assert.Equal(t, cl.Email, client.Email)
	assert.Equal(t, cl.Birthday, client.Birthday)

}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	id, err := insertClient(db, cl)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	_, err = selectClient(db, id)
	require.NoError(t, err)

	err = deleteClient(db, id)
	require.NoError(t, err)

	_, err = selectClient(db, id)
	require.Error(t, err)
	require.Equal(t, sql.ErrNoRows, err)
}
