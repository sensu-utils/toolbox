package toolbox

import "log"

func InitPlugin(name string, event *Event, config interface{}) {
	if event != nil {
		err := ReadEvent(event)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if config != nil {
		err := ReadConfig(name, config)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
