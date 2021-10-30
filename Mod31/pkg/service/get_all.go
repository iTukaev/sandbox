package service
//
//import (
//	"net/http"
//	"strconv"
//)
//
//func (cl *MongoClient) GetAllUsers(w http.ResponseWriter, r *http.Request) {
//	if r.Method == http.MethodGet {
//		w.WriteHeader(http.StatusOK)
//		for key, val := range d.department.Users{
//			resp := "User ID: " + strconv.Itoa(key) + "\t" + val.ToString() + "\n"
//			w.Write([]byte(resp))
//		}
//		return
//	}
//	w.WriteHeader(http.StatusBadRequest)
//}