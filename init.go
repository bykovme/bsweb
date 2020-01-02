package main

import (
	"flag"
	"log"

	"github.com/bykovme/goconfig"
)

func initApp() error {
	usrHomePath, err := goconfig.GetUserHomePath()
	if err == nil {
		log.Println("User home path: " + usrHomePath)
		log.Println("App config default path: " + usrHomePath + cConfigPath)
		err = goconfig.LoadConfig(usrHomePath+cConfigPath, &loadedConfig)
		if err != nil {
			log.Println(err.Error())
			log.Println("applying defaults and trying to take command line parameters")

			loadedConfig.Port = cDefaultPort
			loadedConfig.DBPath = usrHomePath + cDefaultAppPath + "/" + cDbDefaultFileName
			loadedConfig.AssetsPath = usrHomePath + "bsweb/assets/"
			loadedConfig.TemplatesPath = usrHomePath + "bsweb/templates/"

			dbPathPtr := flag.String(cDbParamKey, loadedConfig.DBPath, cDbParamDescription)
			portPtr := flag.String("port", loadedConfig.Port, "app port")
			assetsPtr := flag.String("assets", loadedConfig.AssetsPath, "path to assets")
			templatesPtr := flag.String("templates", loadedConfig.TemplatesPath, "path to html templates")
			flag.Parse()

			if *dbPathPtr != loadedConfig.DBPath {
				loadedConfig.DBPath = *dbPathPtr
				log.Println("CMD parameter for DB is found: " + loadedConfig.DBPath)
			}
			if *portPtr != loadedConfig.Port {
				loadedConfig.Port = *portPtr
				log.Println("CMD parameter for port is found: " + loadedConfig.Port)
			}
			if *assetsPtr != loadedConfig.AssetsPath {
				loadedConfig.AssetsPath = *assetsPtr
				log.Println("CMD parameter for assets is found: " + loadedConfig.AssetsPath)
			}
			if *templatesPtr != loadedConfig.TemplatesPath {
				loadedConfig.TemplatesPath = *templatesPtr
				log.Println("CMD parameter for port is found: " + loadedConfig.TemplatesPath)
			}
		}

	} else {
		log.Println(err.Error())
		return err
	}
	log.Println("Port: " + loadedConfig.Port)
	log.Println("DB path: " + loadedConfig.DBPath)
	log.Println("Assets path: " + loadedConfig.AssetsPath)
	log.Println("HTML templates path: " + loadedConfig.TemplatesPath)
	return nil
}
