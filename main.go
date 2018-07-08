package main
/*
* Developed and published by Freaxbot <support@magiccoder.de>
* This program is free software. It comes without any warranty, to
* the extent permitted by applicable law. You can redistribute it
* and/or modify it under the terms of the Do What The Fuck You Want
* To Public License, Version 2, as published by Sam Hocevar. See
* http://www.wtfpl.net/ for more details.
*/

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
)

func main()  {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Println("Fail to read file: %v", err)
		os.Exit(1)
	}
	StartTCP( cfg.Section("server").Key("port").MustInt(41144));
}