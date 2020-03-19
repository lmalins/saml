package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	l "github.com/lmalins/saml/pkg/models/ldap"
	"gopkg.in/ldap.v3"
)

type application struct {
	resourcesLDAP *l.ResourceModelLDAP
	resourcesAMX  *a.ResourceModelAMX
	cfg           appConfig
	errorLog      *log.Logger
	infoLog       *log.Logger
}

type appConfig struct {
	LdapCfg    ldapCfg
	AmxDRC2Cfg  amxDRC2Cfg
	CertPath   string
	KeyPath    string
	AppContext string
	IdpURL     string
	RootURL    string
	AppPort    string
}

// LdapCfg LDAP cofigurations
type ldapCfg struct {
	LdapServer       string
	LdapBind         string
	LdapPassword     string
	FilterDN         string
	BaseDN           string
	DepartmentNumber string
	GroupsDN         []string
}

// AmxDRCfg AMX DRC2 configurations
type amxDRC2Cfg struct {
	BaseURL     string
	LoginURL    string
	ContainerID string
	LdapAlias   string
	Drc2User    string
	Drc2Pwd     string
	Drc2Sigla   string
	Drc2Ip      string
	GroupsID    []string
}

func main() {
	var cfg appConfig
	if _, err := toml.DecodeFile("config/config.toml", &cfg); err != nil {
		panic(err) // TODO handle error
	}

	/*conn, err := connectLDAP(cfg.LdapCfg.LdapServer, cfg.LdapCfg.LdapBind, cfg.LdapCfg.LdapPassword)
	if err != nil {
		panic(err) // TODO handle error
	}
	defer conn.Close()*/

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		resourcesLDAP: &l.ResourceModelLDAP{
			LDAP:       conn,
			BaseDN:     cfg.LdapCfg.BaseDN,
			DptoNumber: cfg.LdapCfg.DepartmentNumber,
			FilterDN:   cfg.LdapCfg.FilterDN,
			GroupsDN:   cfg.LdapCfg.GroupsDN,
		},
		resourcesAMX: &a.ResourceModelAMX{
			ContainerID: cfg.AmxDRC2Cfg.ContainerID,
			LdapAlias:   cfg.AmxDRC2Cfg.LdapAlias,
			BaseDN:      cfg.LdapCfg.BaseDN,
			BaseURL:     cfg.AmxDRC2Cfg.BaseURL,
			LoginURL:    cfg.AmxDRC2Cfg.LoginURL,
			Drc2User:    cfg.AmxDRC2Cfg.Drc2User,
			Drc2Pwd:     cfg.AmxDRC2Cfg.Drc2Pwd,
			Drc2Sigla:   cfg.AmxDRC2Cfg.Drc2Sigla,
			Drc2Ip:      cfg.AmxDRC2Cfg.Drc2Ip,
			GroupsID:    cfg.AmxDRC2Cfg.GroupsID,
		},
		cfg:      cfg,
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	err = app.start()
	errorLog.Fatal(err)
}

func connectLDAP(ldapServer string, ldapBind string, ldapPwd string) (*ldap.Conn, error) {
	conn, err := ldap.Dial("tcp", ldapServer)

	if err != nil {
		return nil, fmt.Errorf("Failed to connect. %s", err)
	}

	if err := conn.Bind(ldapBind, ldapPwd); err != nil {
		return nil, fmt.Errorf("Failed to bind. %s", err)
	}

	return conn, nil
}
