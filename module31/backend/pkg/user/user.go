package user

import (
	"backend/pkg/db"
	"backend/pkg/friends"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/lib/pq"
	"net/http"
	"strconv"
)

type delRequest struct {
	Source string `json:"source_id"`
}
type AgeChange struct {
	Source string `json:"new age"`
}

func CheckUser(id int) bool {
	db.Connect()
	var result int
	check := `select 1 from users where id = $1 limit 1;`
	e := db.DB.QueryRow(check, id).Scan(&result)
	if e != nil {
		fmt.Println(e.Error())
	}
	if result == 1 {
		return true
	} else {
		return false
	}
}
func GetName(id int) string {
	db.Connect()
	var result string
	listFriends := `select users->'name' from users where id = $1;`
	e := db.DB.QueryRow(listFriends, id).Scan(&result)
	if e != nil {
		fmt.Println(e.Error())
	}
	db.Close()
	return result
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func DeleteUser(w http.ResponseWriter, r *http.Request) error {
	var p delRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	req := &p
	id, _ := strconv.Atoi(req.Source)
	if CheckUser(id) != true {
		return errors.New("no such user")
	}
	InFriends, _ := friends.UserFriendsToSlice(id)

	for _, vol := range InFriends {
		FriendsIn, _ := friends.UserFriendsToSlice(vol)
		index := indexOf(id, FriendsIn)
		removeIndex := RemoveIndex(FriendsIn, index)
		if len(removeIndex) > 0 {
			addFriendQ := `UPDATE users SET users = users || jsonb_build_object('friends', $1::int[]) WHERE id = $2`
			_, e := db.DB.Exec(addFriendQ, pq.Array(removeIndex), vol)
			if e != nil {
				fmt.Println(e.Error())
			}
		} else {
			delFriendQ := `UPDATE users SET users = users || jsonb_build_object('friends', null) WHERE id = $1`
			_, e := db.DB.Exec(delFriendQ, vol)
			if e != nil {
				fmt.Println(e.Error())
			}

		}
	}

	fmt.Fprintf(w, "%s is deleted\n", GetName(id))
	db.Connect()
	deleteStmt := `delete from Users where id=$1`
	_, er := db.DB.Exec(deleteStmt, id)
	if er != nil {
		errors.New("no such user")
	}
	db.Close()
	render.Status(r, http.StatusOK)
	return nil
}
func UpdateAge(w http.ResponseWriter, r *http.Request) error {
	var p AgeChange
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	rep := &p
	idString := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idString)
	if CheckUser(id) != true {
		return errors.New("no such user")
	}
	newAge, _ := strconv.Atoi(rep.Source)
	if newAge > 0 {
		db.Connect()
		updateAge2 := `UPDATE users SET users = users || jsonb_build_object('age', $1::int) WHERE id = $2`
		_, e := db.DB.Exec(updateAge2, newAge, id)
		if e != nil {
			fmt.Println(e.Error())
		}
		db.Close()
		fmt.Fprintf(w, "Age of user %v is update to %+v\n", GetName(id), newAge)
	} else {
		return errors.New("something's wrong with his age")
	}

	render.Status(r, http.StatusOK)
	return nil
}
