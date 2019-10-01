package main
  
import (
        "fmt"
        "net/http"
)

//Home function, writes out HTML form to the user
func Home(w http.ResponseWriter, r *http.Request) {
        html := `<html>
                <head>
                <title>Hi. Who are you?</title>
                <style>
                body {
                        background-color: silver;
                }
                label {
                        font-family: verdana;
                        font-size: 18px;
                }

                </style>
                </head>
                <body>

                <form action="http://54.196.230.194:30002/getname" method="post" enctype="multipart/form-data">
                <label for="text">Please enter your name:</label></br>
                <input type="text" name="name" size="22"></br>
                <input type="submit" name="submit" value="Submit">
                </form>

                </body>
                </html>`

                w.Write([]byte(fmt.Sprintf(html)))
        }

//Form processor, checks for any errors and writes back based on user input
func GetName(w http.ResponseWriter, r *http.Request) {
        err := r.ParseMultipartForm(200000) //Grab and error check text form and send err to browser

        if err != nil {
                fmt.Fprintln(w, err)
                return
        }

        //write out a response with the provided name to both console and browser
        fmt.Fprintln(w, "Hello", r.FormValue("name"), "and thank you very much for having me here today!")
        fmt.Println("We just heard from:", r.FormValue("name"))
}

func main() {
        mux := http.NewServeMux()

        mux.HandleFunc("/", Home)
        mux.HandleFunc("/getname", GetName)

        http.ListenAndServe(":8080", mux)
}
