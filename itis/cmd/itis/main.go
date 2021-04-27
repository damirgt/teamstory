package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main()  {
	logrus.Info("Hello world")

	var port = os.Getenv("PORT")

	if len(port) == 0{
		logrus.Fatal("Port is not set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//w.WriteHeader(http.StatusOK)
		w.Write([]byte("This is a sample Page."))
	})

	http.ListenAndServe(":"+ port, nil)
}