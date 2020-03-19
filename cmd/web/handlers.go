package main

import (
	b64 "encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
	"github.com/crewjam/saml/samlsp"
	"github.com/lmalins/saml/tree/master/pkg/modelsmodels"
	"github.com/lmalins/saml/tree/master/pkg/stringutils"
)

func (app *application) hello(w http.ResponseWriter, r *http.Request) {
	usrValidation := samlsp.Token(r.Context()).Attributes.Get("urn:mace:dir:attribute-def:authresult")
	uid := samlsp.Token(r.Context()).Attributes.Get("urn:mace:dir:attribute-def:NameId")
	bp := samlsp.Token(r.Context()).Attributes.Get("urn:mace:dir:attribute-def:BP")

	pss := stringutils.Random(15)

	m := fmt.Sprintf("BP: %s, NameId: %s, pss: %s", bp, uid, pss)

	if strings.ToLower(usrValidation) != "true" {
		nsec := time.Now().UnixNano()
		tmpl := template.Must(template.ParseFiles("./ui/html/error.html"))
		d := fmt.Sprintf("Ocorreu um erro durante o processo de autenticação: 401-%d", nsec)
		log.Printf("%s | %s", d, m)

		msg := models.Msg{
			Title: "User login has fail by this user.",
			Desc:  d,
		}
		tmpl.Execute(w, msg)
		return
	}

	if ((uid == "" || strings.ToLower(uid) == "null" || uid is nil) && (bp == "" || strings.ToLower(bp) == "null" || bp is nil)) { 
		nsec := time.Now().UnixNano()
		tmpl := template.Must(template.ParseFiles("./ui/html/error.html"))
		d := fmt.Sprintf("Ocorreu um erro durante o processo de autenticação: 401-%d", nsec)
		log.Printf("%s | %s", d, m)

		msg := models.Msg{
			Title: "Oops!",
			Desc:  d,
		}
		tmpl.Execute(w, msg)
		return
	}

	log.Print(m)

	// Esta sendo usado o e-mail como do UID
	/*dn, err := app.resourcesLDAP.CreateOrUpdateUser(mail, firstName, lastName, pss, mail, costCenter)
	if err != nil {
		app.serverError(w, err)
		return
	}*/

	uid64 := b64.StdEncoding.EncodeToString([]byte(user))
	pss64 := b64.StdEncoding.EncodeToString([]byte(pss))	

	//err = app.resourcesAMX.CreateResource(mail, firstName, lastName, pss)
	user, err = app.resourcesAMX.CreateResource(uid, bp, pss64)
	if err != nil {
		app.serverError(w, err)
		return
	}

	tmpl := template.Must(template.ParseFiles("./ui/html/layout.html"))

	data := models.User{
		//Usr: uid64,
		Usr: user
		//Pss: pss64,
		Pss: pss,
		//Sigla: "BR001",
		Sigla: "csl05",
		Ip: "1"
	}
	tmpl.Execute(w, data)
}

func splitName(fullname string) (string, string) {
	words := strings.Fields(fullname)
	return words[0], words[len(words)-1]
}

func (app *application) testAuth(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	key := keys[0]

	log.Println("Url Param 'key' is: " + string(key))

	kDecoded, _ := b64.StdEncoding.DecodeString(key)
	log.Println("Key decoded is: " + string(kDecoded))
	s := strings.Split(string(kDecoded), ":")

	tmpl := template.Must(template.ParseFiles("./ui/html/layout.html"))

	data := models.User{
		Usr: s[0],
		Pss: s[1],
		//Sigla: "BR001",
		Sigla: "csl05",
		Ip: "1"
	}
	tmpl.Execute(w, data)
}
