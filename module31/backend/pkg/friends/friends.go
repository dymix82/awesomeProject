package friends

import (
	"backend/pkg/db"
	"backend/pkg/user"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/render"
	"github.com/lib/pq"
	"golang.org/x/exp/slices"
	"net/http"
	"strconv"
	"strings"
)

type FriendsH struct {
	ID   int
	Data json.RawMessage
}
type FriendsD struct {
	ID   int    `db:"id"`
	Data string `db:"users"`
}

type FriendRequest struct {
	Source string `json:"source_id"`
	Target string `json:"target_id"`
}

func MaxID() int {
	db.Connect()
	var result int
	e := db.DB.QueryRow(`select max(id) from users`).Scan(&result)
	if e != nil {
		fmt.Println(e.Error())
	}
	return result
}
func makeFriends(id1, id2 int) error {
	if user.CheckUser(id1) != true {
		return errors.New("no such user")
	}
	if user.CheckUser(id2) != true {
		return errors.New("no such user")
	}
	db.Connect()
	result1, _ := UserFriendsToSlice(id1)
	result2, _ := UserFriendsToSlice(id2)
	res1 := append(result2, id1)
	res2 := append(result1, id2)
	addFriendQ := `UPDATE users SET users = users || jsonb_build_object('friends', $1::int[]) WHERE id = $2`
	_, e := db.DB.Exec(addFriendQ, pq.Array(res1), id2)
	if e != nil {
		fmt.Println(e.Error())
	}
	_, e2 := db.DB.Exec(addFriendQ, pq.Array(res2), id1)
	if e2 != nil {
		fmt.Println(e.Error())
	}
	db.Close()
	return nil
}

func ListFriends(w http.ResponseWriter, r *http.Request) error {
	idQuery := r.URL.Query().Get("id")
	if idQuery == "" {
		return errors.New(idQuery)
	}
	id, err := strconv.Atoi(idQuery)
	lS := MaxID()
	if err != nil || id > lS {
		return errors.New(idQuery)
	}
	friendSlice, _ := UserFriendsToSlice(id)
	for _, vol := range friendSlice {
		db.Connect()
		row := FriendsD{}
		err = db.DB.Get(&row, "SELECT * FROM users WHERE id=$1", vol)
		data := FriendsH{}
		data.ID = row.ID
		data.Data = []byte(row.Data)
		j, _ := json.Marshal(&data)
		w.Write(j)
	}
	render.Status(r, http.StatusOK)
	db.Close()
	return nil
}
func MakeFriends(w http.ResponseWriter, r *http.Request) error {
	var p FriendRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	data := &p
	src, _ := strconv.Atoi(data.Source)
	tgt, _ := strconv.Atoi(data.Target)
	if user.CheckUser(src) != true {
		return errors.New("no such user")
	}
	if user.CheckUser(tgt) != true {
		return errors.New("no such user")
	}
	srcSlice, _ := UserFriendsToSlice(src)
	idx := slices.Index(srcSlice, tgt)
	if idx == -1 {
		fmt.Fprintf(w, "%s and %s are friends now \n", user.GetName(src), user.GetName(tgt))
		makeFriends(src, tgt)
		render.Status(r, http.StatusOK)
		return nil
	} else {
		return errors.New("no such user")
	}
}
func UserFriendsToSlice(id int) ([]int, error) {
	db.Connect()
	var result string
	listFriends := `select users->'friends' from users where id = $1;`
	e := db.DB.QueryRow(listFriends, id).Scan(&result)
	if e != nil {
		fmt.Println(e.Error())
	}
	if result != "null" {
		trimmed := strings.Trim(result, "[]")
		strings := strings.Split(trimmed, ", ")
		ints := make([]int, len(strings))
		for i, s := range strings {
			ints[i], _ = strconv.Atoi(s)
		}
		return ints, nil
	} else {
		return nil, errors.New("no friends")
	}
}
