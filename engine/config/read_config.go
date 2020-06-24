package config

var Center =  &CenterConfig{
	ListenAddr:"127.0.0.1:9999",
	HTTPAddr:"127.0.0.1:9999",
	LogFile:"center.log",
	LogLevel:"debug",
}
var Login = &LoginConfig{
	ListenAddr:"127.0.0.1:11001",
	HTTPAddr:"127.0.0.1:11001",
	LogFile:"dispatcher.log",
	LogLevel:"debug",
}
var Dispatcher = &DispatcherConfig{
	ListenAddr:"127.0.0.1:12001",
	HTTPAddr:"127.0.0.1:12001",
	LogFile:"dispatcher.log",
	LogLevel:"debug",
}
var Gate = &GateConfig{
	ListenAddr:"127.0.0.1:13001",
	HTTPAddr:"127.0.0.1:13001",
	LogFile:"dispatcher.log",
	LogLevel:"debug",
}
var Game = &GameConfig{
	ListenAddr:"127.0.0.1:14001",
	HTTPAddr:"127.0.0.1:14001",
	LogFile:"dispatcher.log",
	LogLevel:"debug",
}
func init(){

}