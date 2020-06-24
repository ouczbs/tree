package config
type CenterConfig struct{
	ListenAddr    string
	AdvertiseAddr string
	HTTPAddr      string
	LogFile       string
	LogStderr     bool
	LogLevel      string
}
// DispatcherConfig defines fields of dispatcher config
type DispatcherConfig struct {
	ListenAddr    string
	AdvertiseAddr string
	HTTPAddr      string
	LogFile       string
	LogStderr     bool
	LogLevel      string
}
type GateConfig struct {
	ListenAddr    string
	AdvertiseAddr string
	HTTPAddr      string
	LogFile       string
	LogStderr     bool
	LogLevel      string
}
type LoginConfig struct {
	ListenAddr    string
	AdvertiseAddr string
	HTTPAddr      string
	LogFile       string
	LogStderr     bool
	LogLevel      string
}
type GameConfig struct {
	ListenAddr    string
	AdvertiseAddr string
	HTTPAddr      string
	LogFile       string
	LogStderr     bool
	LogLevel      string
}