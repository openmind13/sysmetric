package configurator

type Config struct{}

var (
	ConfigChan = make(chan Config, 1)
)

func ListenConfig() {

}
