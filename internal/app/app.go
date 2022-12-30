package app

import (
	"gorm.io/gorm"
)

// AppName is the app's name .
const AppName string = "hetab"

const asciiArt = `
                                                                                   bbbbbbbb            
HHHHHHHHH     HHHHHHHHH                             tttt                           b::::::b            
H:::::::H     H:::::::H                          ttt:::t                           b::::::b            
H:::::::H     H:::::::H                          t:::::t                           b::::::b            
HH::::::H     H::::::HH                          t:::::t                            b:::::b            
  H:::::H     H:::::H      eeeeeeeeeeee    ttttttt:::::ttttttt      aaaaaaaaaaaaa   b:::::bbbbbbbbb    
  H:::::H     H:::::H    ee::::::::::::ee  t:::::::::::::::::t      a::::::::::::a  b::::::::::::::bb  
  H::::::HHHHH::::::H   e::::::eeeee:::::eet:::::::::::::::::t      aaaaaaaaa:::::a b::::::::::::::::b 
  H:::::::::::::::::H  e::::::e     e:::::etttttt:::::::tttttt               a::::a b:::::bbbbb:::::::b
  H:::::::::::::::::H  e:::::::eeeee::::::e      t:::::t              aaaaaaa:::::a b:::::b    b::::::b
  H::::::HHHHH::::::H  e:::::::::::::::::e       t:::::t            aa::::::::::::a b:::::b     b:::::b
  H:::::H     H:::::H  e::::::eeeeeeeeeee        t:::::t           a::::aaaa::::::a b:::::b     b:::::b
  H:::::H     H:::::H  e:::::::e                 t:::::t    tttttta::::a    a:::::a b:::::b     b:::::b
HH::::::H     H::::::HHe::::::::e                t::::::tttt:::::ta::::a    a:::::a b:::::bbbbbb::::::b
H:::::::H     H:::::::H e::::::::eeeeeeee        tt::::::::::::::ta:::::aaaa::::::a b::::::::::::::::b 
H:::::::H     H:::::::H  ee:::::::::::::e          tt:::::::::::tt a::::::::::aa:::ab:::::::::::::::b  
HHHHHHHHH     HHHHHHHHH    eeeeeeeeeeeeee            ttttttttttt    aaaaaaaaaa  aaaabbbbbbbbbbbbbbbb
`

type Application struct {
	DB           *gorm.DB
	Repositories *Repositories
}

func Banner() string {
	return asciiArt
}
