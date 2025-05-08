package main

import (
	// "encoding/json"
	"fmt"
	// "io"
	"log"
	"net/http"
)

type user struct{
	Name string `json:"name"`
	Age string `json:"age"`
	City string `json:"city"`
}


func main() {

	port := ":3000"



	http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {

		// fmt.Fprintf(w , "Hello there is more to be done on the portal")
		// fmt.Println(r.Method)
		// if r.Method == http.MethodGet{
        // w.Write([]byte ("now we are into the get request "))
		// return
		// }

		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("hello get Post"))
		case http.MethodPut:
			w.Write([]byte("hello get Put"))    
		case http.MethodPatch:
			w.Write([]byte("hello get Patch"))
		case http.MethodDelete:
			w.Write([]byte("hello get Delete"))
		default:
			w.Write([]byte("hello you are into the main route"))
		}
	})

    http.HandleFunc( "/student", func(w http.ResponseWriter, r *http.Request) {

	    // fmt.Fprintf(w, "hello students route")

		// w.Write([]byte ("hello student route"))

		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("hello get Post"))
		case http.MethodPut:
			w.Write([]byte("hello get Put"))    
		case http.MethodPatch:
			w.Write([]byte("hello get Patch"))
		case http.MethodDelete:
			w.Write([]byte("hello get Delete"))
		default:
			w.Write([]byte("hello you are into the student route"))
		}
   })


   http.HandleFunc("/teacher", func(w http.ResponseWriter, r *http.Request) {
    // fmt.Println("Received method:", r.Method) // Debugging line

    // switch r.Method {
    // case http.MethodGet:
    //     w.Write([]byte("hello get method"))
    // case http.MethodPost:
	// 	// parse form data (necessary for x-ww-form-urlencoded)
	// 	err := r.ParseForm()
	// 	if err!=nil {
	// 		http.Error(w,"Error parsing form",http.StatusBadRequest)
	// 	}
	// 	fmt.Println("Form:" , r.Form)
   
		  
	// 	// perpare response 

	// 	response := make(map[string]interface{})
	// 	for key , value := range r.Form{
	// 		response[key] = value[0]
	// 	}

	// 	fmt.Println("Processed Response map", response)


	// 	// Raw Body 
	// 	body, err := io.ReadAll(r.Body)
	// 	if err != nil {
	// 		return
	// 	}
	// 	defer r.Body.Close() 

	// 	fmt.Println("Raw Body:", body)
    //     fmt.Println("Raw Body:", string(body))


	// 	// If expect data unmarshal it 
	// 	var userInstance user
	// 	err = json.Unmarshal(body, &userInstance)
	// 	if err != nil {
	// 		return
	// 	}
	// 	w.Write([]byte ("its send the raw data"))
    //     fmt.Println(userInstance)

	 switch r.Method {
    case http.MethodGet:
        w.Write([]byte("hello get Post"))
    case http.MethodPut:
        w.Write([]byte("hello get Put"))    
    case http.MethodPatch:
        w.Write([]byte("hello get Patch"))
    case http.MethodDelete:
        w.Write([]byte("hello get Delete"))
    default:
        w.Write([]byte("hello you are into the teacher route"))
    }
})

	fmt.Println("Server is running on the port :" , port)

	err := http.ListenAndServe(port , nil)

	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}