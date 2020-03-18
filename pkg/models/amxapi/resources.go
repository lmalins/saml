package amxapi

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
	"github.com/delley/saml-auth/pkg/models"
)

// ResourceModelAMX Define a ResourceModel type wich wraps a ldap.Conn connection.
type ResourceModelAMX struct {
	ContainerID string
	LdapAlias   string
	BaseDN      string
	BaseURL     string
	DrcUser     string
	DrcPwd      string
	GroupsID    []string
}

// CreateResource Create a resource in AMX DRC2
func (r *ResourceModelAMX) CreateResource(uid string, bp string, pss string) (string, error) {
	lbl := fmt.Sprintf("%s %s", uid, bp)

	//Verify the user exist by Latam BP
	drcURL := fmt.Sprintf("%s/api/CentralSchema/User/ByLogin/" + bp, r.BaseURL)
	req, err := http.NewRequest("GET", drcURL)
	if err != nil {
		return "", fmt.Errorf("HTTP request error\n%s", err)
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(14951)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error making a request to verify user exists on DRC2 '%s'\n%s", drcURL, err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response from verify user exists on DRC2 '%s'\n%s", drcURL, err)
	}
	
	log.Printf("%s", string(body))

	user := models.UserByPerson{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return "", fmt.Errorf("Error parsing response from '%s' - '%s'\n%s", drcURL, string(body), err)
	}

	log.Printf("User: %v", user.Content)

	if(user.Content.ID <= 0){
	{		
		//TODO //Verify the user exist by Name ID
		drcURL := fmt.Sprintf("%s/api/CentralSchema/User/ByLogin/" + uid, r.BaseURL)
		req, err := http.NewRequest("GET", drcURL)
		if err != nil {
			return "", fmt.Errorf("HTTP request error\n%s", err)
		}
		req.Header.Set("Accept", "application/json")
		req.SetBasicAuth(14951)

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return "", fmt.Errorf("Error making a request to verify user exists on DRC2 '%s'\n%s", drcURL, err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("Error reading response from verify user exists on DRC2 '%s'\n%s", drcURL, err)
		}
		
		log.Printf("%s", string(body))

		user := models.UserByPerson{}
		err = json.Unmarshal(body, &user)
		if err != nil {
			return "", fmt.Errorf("Error parsing response from '%s' - '%s'\n%s", drcURL, string(body), err)
		}

		log.Printf("User: %v", user.Content)

		if(user.Content.ID <= 0)
		{
			return "", fmt.Errorf("User not exists on DRC2, please contact the administration of system.")
		}
	}

	u := &models.UserModel{}
	u.ID  = user.Content.ID                    
	u.LdapCustomerConfigurationID = user.Content.LdapCustomerConfigurationID 
	u.PersonID = user.Content.PersonID                 
	u.AccessTypeID = user.Content.AccessTypeID               
	u.CustomerID = user.Content.CustomerID                   
	u.AuthenticationTypeID = user.Content.AuthenticationTypeID        
	u.Login = user.Content.Login                       
	u.Password = pss                    
	u.PasswordExpiration = "2120-12-31 23:59:59"          
	u.Enabled = user.Content.Enabled                     
	u.ValidateLogin = user.Content.ValidateLogin               
	u.SecurityQuestion = user.Content.SecurityQuestion            
	u.SecurityAnswer = user.Content.SecurityAnswer              
	u.AuditUser = user.Content.AuditUser                   
	u.AuditUpdateDate = time.Now().Format("2006-01-02 15:04:05")            
	u.Sha = user.Content.Sha                         
	u.LdapLogin = user.Content.LdapLogin                   
	u.PasswordAttempts = 0            
	u.UseIMConnect = user.Content.UseIMConnect                

	payload, err := json.Marshal(u)
	if err != nil {
		return "", fmt.Errorf("JSON encoding error\n%s", err)
	}
	log.Printf("DRC user: %s", payload)

	drcURL = fmt.Sprintf("%s/api/CentralSchema/User/Update", r.BaseURL)
	req, err = http.NewRequest("PUT", drcURL, bytes.NewBuffer(payload))
	if err != nil {
		return "", fmt.Errorf("HTTP new request Put user error error\n%s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(14951)

	client = http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error making a request to '%s'\n%s", drcURL, err)
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response from '%s' - '%s'\n%s", drcURL, string(body), err)
	}

	log.Printf("update body: %s", string(body))

	return u.Login, nil
}

func makeGroups(g []string) []models.AddGroup {
	var groups []models.AddGroup
	for _, gid := range g {
		groups = append(groups, models.AddGroup{
			GUID: gid,
		})
	}
	return groups
}
