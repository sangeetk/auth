package middleware

import (
	"github.com/codegangsta/negroni"
	_ "github.com/unrolled/render"
	"net/http"
	//"os"
)

func Database() negroni.HandlerFunc {
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
