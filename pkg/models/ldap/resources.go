package ldap

import (
	"fmt"
	"log"
	"strings"

	"gopkg.in/ldap.v3"
)

// ResourceModelLDAP Define a ResourceModel type wich wraps a ldap.Conn connection.
type ResourceModelLDAP struct {
	LDAP       *ldap.Conn
	BaseDN     string
	DptoNumber string
	FilterDN   string
	GroupsDN   []string
}

// CreateOrUpdateUser This will create a new resource or update an existing resource into the LDAP.
func (r *ResourceModelLDAP) CreateOrUpdateUser(uid string, firstName string, lastName string, pss string, mail string, costCenter string) (string, error) {

	dn := ""

	result, err := r.LDAP.Search(ldap.NewSearchRequest(
		r.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		filter(uid, r.FilterDN),
		[]string{"dn", "businessCategory"},
		nil,
	))

	if err != nil {
		return "", fmt.Errorf("Failed to find user. %s", err)
	}

	if len(result.Entries) < 1 {
		dn, err = r.newUser(firstName, lastName, pss, mail, uid, costCenter)
		if err != nil {
			return "", fmt.Errorf("Failed to create user. %s", err)
		}
	} else if len(result.Entries) > 1 {
		return "", fmt.Errorf("Too many entries returned")
	} else {
		entry := result.Entries[0]
		dn = entry.DN
		log.Printf("Entry: %v", entry)
		bc := entry.GetAttributeValue("businessCategory")
		log.Printf("businessCategory: %s", bc)
		modify := ldap.NewModifyRequest(entry.DN, nil)

		if len(costCenter) == 0 && len(bc) > 0 {
			modify.Delete("businessCategory", nil)
			modify.Delete("roomNumber", nil)
		} else if len(costCenter) > 0 && len(bc) == 0 {
			roomNumber := getPrincipalCC(costCenter)
			modify.Add("businessCategory", []string{costCenter})
			modify.Add("roomNumber", []string{roomNumber})
		} else if len(costCenter) > 0 && len(bc) > 0 && costCenter != bc {
			roomNumber := getPrincipalCC(costCenter)
			modify.Replace("businessCategory", []string{costCenter})
			modify.Replace("roomNumber", []string{roomNumber})
		}

		if len(modify.Changes) > 0 {
			err := r.LDAP.Modify(modify)
			if err != nil {
				return "", fmt.Errorf("Changes could not be applied: %s", err.Error())
			}
		}

		passwordModifyRequest := ldap.NewPasswordModifyRequest(entry.DN, "", pss)
		_, err = r.LDAP.PasswordModify(passwordModifyRequest)

		if err != nil {
			return "", fmt.Errorf("Password could not be changed: %s", err.Error())
		}
	}

	return dn, nil
}

func (r *ResourceModelLDAP) newUser(firstName string, lastName string, pss string, mail string, uid string, costCenter string) (string, error) {
	cn := fmt.Sprintf("%s %s", firstName, lastName)
	dn := fmt.Sprintf("uid=%s,%s", uid, r.BaseDN)
	a := ldap.NewAddRequest(dn, nil)

	a.Attribute("objectClass", []string{"top"})
	a.Attribute("objectClass", []string{"person"})
	a.Attribute("objectClass", []string{"organizationalPerson"})
	a.Attribute("objectClass", []string{"inetOrgPerson"})

	a.Attribute("cn", []string{cn})
	a.Attribute("sn", []string{lastName})
	a.Attribute("departmentNumber", []string{r.DptoNumber})
	a.Attribute("mail", []string{mail})
	a.Attribute("uid", []string{uid})
	a.Attribute("userPassword", []string{pss})
	if len(r.GroupsDN) > 0 {
		a.Attribute("l", r.GroupsDN)
	}

	if len(costCenter) > 0 {
		roomNumber := getPrincipalCC(costCenter)
		a.Attribute("businessCategory", []string{costCenter})
		a.Attribute("roomNumber", []string{roomNumber})
	}

	err := r.LDAP.Add(a)
	if err != nil {
		return "", fmt.Errorf("Entry NOT done: %s. %s", dn, err)
	}

	log.Printf("Entry DONE: %s", dn)
	return dn, nil
}

func filter(needle string, filterDN string) string {
	res := strings.Replace(
		filterDN,
		"{0}",
		needle,
		-1,
	)

	return res
}

func getPrincipalCC(cc string) string {
	tmp := strings.Split(cc, ",")
	return tmp[0]
}
