package cmd

import (
	"database/sql"
	api "github.com/dunzoit/Store/Api"
	"github.com/dunzoit/Store/config"
	"net/http"
	"sync"
)

// Execute execution of start here
func Execute() {
	db, err := mustPrepareDB()
	if err != nil {
		panic(err)
	}
	serveApp(db)
}

func serveApp(db *sql.DB) {
	var wg sync.WaitGroup
	app := api.NewApp(db)
	wg.Add(1)
	go func() {
		err := http.ListenAndServe(config.PORT, app.Router())
		if err != nil {
			panic(err)
		}
		wg.Done()
	}()
	wg.Wait()

}
