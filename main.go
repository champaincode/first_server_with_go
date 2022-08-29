package main 

import ("net/http"

)

func  main()  {
	//routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	//start server
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello world"))
}

func contactHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Contacto perro"))
}
