package test1

import "fmt"

type Config struct {
	APIKey   string `json:"api_key"`
	Debug    bool   `json:"debug"`
	Endpoint string `json:"endpoint"`
}

type ConfigEx struct {
	*Config
	name string
}

func Test_embeded_struct() {

	cfgEx1 := ConfigEx{
		Config: &Config{
			APIKey:   "key1",
			Debug:    true,
			Endpoint: "endPt1",
		},
		name: "ddd",
	}

	cfgEx2 := ConfigEx{
		Config: &Config{
			APIKey:   "key1",
			Debug:    true,
			Endpoint: "endPt1",
		},
		name: "ddd",
	}

	fmt.Printf("%p ,%p\n", cfgEx1.Config, cfgEx2.Config)
}

func Test_embeded_struct01() {

	var cfg = Config{
		APIKey:   "key1",
		Debug:    true,
		Endpoint: "endPt1",
	}
	cfgEx1 := ConfigEx{
		Config: &cfg,
		name:   "ddd",
	}

	cfgEx2 := ConfigEx{
		Config: &cfg,
		name:   "ddd",
	}

	fmt.Printf("%p ,%p\n", cfgEx1.Config, cfgEx2.Config)
}
