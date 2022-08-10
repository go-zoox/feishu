package config

type Config struct {
	BaseURI    string `env:"BASE_URI" json:"base_url"`
	AppID      string `env:"APP_ID" config:"app_id"`
	ApptSecret string `env:"APP_SECRET" config:"app_t_secret"`
}
