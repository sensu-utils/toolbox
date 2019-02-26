# sensu-go-utils

A set of simple go utilities designed to be used for building sensu handlers.

## API

### Plugin

```
func InitPlugin(name string, event *Event, config interface{})
```

### Events

```
type Event struct {
	types.Event
}

func ReadEvent(event *Event) error
```

### Config

```
func ReadConfig(configName string, target interface{}) error
```

## Projects Using Sensu-Go-Utils

- [sensu-irc-handler](https://github.com/belak/sensu-irc-handler)