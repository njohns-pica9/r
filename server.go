package main

import (
	"github.com/RangelReale/osin"
	"html/template"
	"log"
	"net/http"
)

func initServer() {
	sconfig := osin.NewServerConfig()
	sconfig.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE}
	sconfig.AllowedAccessTypes = osin.AllowedAccessType{osin.AUTHORIZATION_CODE}
	sconfig.AllowGetAccessRequest = false
	sconfig.AllowClientSecretInParams = true

	server := osin.NewServer(sconfig, rstore)

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Authorize Endpoint")

		resp := server.NewResponse()
		defer resp.Close()

		if ar := server.HandleAuthorizeRequest(resp, r); ar != nil {
			if r.Method == "GET" {
				t, err := template.ParseFiles("data/index.html")
				if err != nil {
					log.Fatalln(err)
				}

				data := struct {
					Type        osin.AuthorizeRequestType
					Id          string
					RedirectUri string
					State       string
				}{
					Type:        ar.Type,
					Id:          ar.Client.GetId(),
					RedirectUri: ar.RedirectUri,
					State:       ar.State,
				}

				t.Execute(w, data)
				return
			} else {
				r.ParseForm()
				if r.Form.Get("login") == "test" && r.Form.Get("password") == "test" {
					ar.Authorized = true
					ar.UserData = struct{ Login string }{Login: "test"}
					server.FinishAuthorizeRequest(resp, r, ar)
				}
			}
		}

		if resp.IsError && resp.InternalError != nil {
			log.Printf("ERROR: %s\n", resp.InternalError)
		}

		osin.OutputJSON(resp, w, r)
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		resp := server.NewResponse()
		defer resp.Close()

		if ar := server.HandleAccessRequest(resp, r); ar != nil {
			log.Println("Access request Good")
			ar.Authorized = true
			server.FinishAccessRequest(resp, r, ar)
		}

		if resp.IsError && resp.InternalError != nil {
			log.Printf("ERROR: %s\n", resp.InternalError)
		}

		osin.OutputJSON(resp, w, r)
	})

	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		resp := server.NewResponse()
		defer resp.Close()

		if ir := server.HandleInfoRequest(resp, r); ir != nil {
			server.FinishInfoRequest(resp, r, ir)
		}

		if resp.IsError && resp.InternalError != nil {
			log.Printf("ERROR: %s\n", resp.InternalError)
		}

		osin.OutputJSON(resp, w, r)
	})

	http.ListenAndServe(":14000", nil)
}
