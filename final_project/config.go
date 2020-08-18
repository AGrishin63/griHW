package main

import (
	"time"
)

//Конфигурируемые переменные и декларативные ресурсы

// Общие
//var ServiceStarted bool = false
// Чёрные и белые листы

var BlackList map[string]subnet
var WhiteList map[string]subnet

// Слайс возможных масок подсети
var masks = [33][4]byte{
	{0, 0, 0, 0}, {128, 0, 0, 0}, {192, 0, 0, 0}, {224, 0, 0, 0}, {240, 0, 0, 0}, {248, 0, 0, 0}, {252, 0, 0, 0}, {254, 0, 0, 0}, {255, 0, 0, 0},
	{255, 128, 0, 0}, {255, 192, 0, 0}, {255, 224, 0, 0}, {255, 240, 0, 0}, {255, 248, 0, 0}, {255, 252, 0, 0}, {255, 254, 0, 0}, {255, 255, 0, 0},
	{255, 255, 128, 0}, {255, 255, 192, 0}, {255, 255, 224, 0}, {255, 255, 240, 0}, {255, 255, 248, 0}, {255, 255, 252, 0}, {255, 255, 254, 0}, {255, 255, 255, 0},
	{255, 255, 255, 128}, {255, 255, 255, 192}, {255, 255, 255, 224}, {255, 255, 255, 240}, {255, 255, 255, 248}, {255, 255, 255, 252}, {255, 255, 255, 254}, {255, 255, 255, 255},
}

//Основные ограничители

// Максимальное количество попыток аутентификации под указанным параметром в мируту (значения по умолчанию)
type Config struct {
	N       int    `yaml:"N"`
	M       int    `yaml:"M"`
	K       int    `yaml:"K"`
	Port    string `yaml:"Port"`
	LogPath string `yaml:"LogPath"`
}

var Cfg Config

// Максимальное количество элементов КЭШ для:
var MaxBucketsInCache int = 1000 //Логинов

// Максимальное число горутин
//var MaxGoroutin int = 30
//var goCount = 0

// Контейнеры подсчёта ограничивающих событий
var TimestampList []time.Time

//LRU Кэши для хранения контейнеров
var LoginLru Cache
var PasswLru Cache
var IpLru Cache

// Каналы для передачи данных, возврата результата и для останова
var inChan chan Auth
var outChan chan bool
var stopChan chan bool

type Auth struct {
	l string
	p string
	i string
}