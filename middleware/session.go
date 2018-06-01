package middleware

import (
	"github.com/codegangsta/negroni"
	"net/http"
	//"os"
)

// Session - Session Middle to extend session timeouts
func Session() negroni.HandlerFunc {
	//database := os.Getenv("DB_NAME")
	/*session, err := mgo.Dial("127.0.0.1:27017")

	if err != nil {
		panic(err)
	}*/

	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		//reqSession := session.Clone()
		//defer reqSession.Close()
		//db := reqSession.DB(database)
		//SetDb(r, db)
		next(rw, r)
	})
}
