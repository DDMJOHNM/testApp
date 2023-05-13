package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	test "test/gen"

	_ "github.com/lib/pq"
)

func main() {

	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		GetAsString("DB_USER", "postgres"),
		GetAsString("DB_PASSWORD", "mysecretpassword"),
		GetAsString("DB_HOST", "localhost"),
		GetAsInt("DB_PORT", 5432),
		GetAsString("DB_NAME", "postgres"),
	)

	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalln("Error from database ping: ", err)
	}

	st := test.New(db)
	ctx := context.Background()

	_, err = st.CreateUsers(ctx, test.CreateUsersParams{
		UserName:     "testuser",
		PassWordHash: "hash",
		Name:         "test",
	})

	if err != nil {
		log.Fatalln("Error creating user: ", err)
	}

	eid, err := st.CreateExercise(ctx, "Excerise1")

	if err != nil {
		log.Fatalln("Error creating exercise", err)
	}

	set, err := st.CreateSet(ctx, test.CreateSetParams{
		ExerciseID: eid,
		Weight:     100,
	})

	if err != nil {
		log.Fatalln("Error updating exercise :", err)
	}

	set, err = st.UpdateSet(ctx, test.UpdateSetParams{
		ExerciseID: eid,
		SetID:      set.SetID,
		Weight:     2000,
	})

	if err != nil {
		log.Fatalln("Error updating set :", err)
	}

	log.Println("Done!")

	u, err := st.ListUsers(ctx)

	for _, usr := range u {
		fmt.Println(fmt.Sprintf("Name : %s, ID : %d", usr.Name, usr.UserID))
	}

}
