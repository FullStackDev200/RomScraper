package scraping

type Platform int

// Retroachievements Supported Systems
//
//
// | System / Platform                       | Hash Type          |
// | --------------------------------------- | ------------------ |
// | Atari 2600                              | MD5                |
// | Atari 7800                              | MD5 (header-aware) |
// | Atari Jaguar                            | MD5                |
// | Atari Jaguar CD                         | Custom             |
// | Atari Lynx                              | MD5 (header-aware) |
// | Bandai WonderSwan / Color               | MD5                |
// | Channel F                               | MD5                |
// | ColecoVision                            | MD5                |
// | Dreamcast                               | Custom             |
// | Famicom Disk System                     | MD5 (header-aware) |
// | Game Boy / Color / Advance              | MD5                |
// | Game Gear                               | MD5                |
// | GameCube                                | Custom             |
// | Intellivision                           | MD5                |
// | MSX / MSX2                              | MD5                |
// | Master System                           | MD5                |
// | Mega Drive / Genesis                    | MD5                |
// | NEC PC-8001 / 8801                      | MD5                |
// | NES / Famicom                           | Custom             |
// | Neo Geo CD                              | Custom             |
// | Neo Geo Pocket / Color                  | MD5                |
// | Nintendo 64                             | MD5                |
// | Nintendo DS                             | Custom             |
// | Odyssey²                                | MD5                |
// | PC Engine / TurboGrafx / SuperGrafx     | MD5 (header-aware) |
// | PC Engine CD / TurboGrafx-CD            | Custom             |
// | PC-FX                                   | Custom             |
// | PSP                                     | Custom             |
// | PlayStation                             | Custom             |
// | PlayStation 2                           | Custom             |
// | Pokémon Mini                            | MD5                |
// | SG-1000                                 | MD5                |
// | SNES / SFC / Satellaview / Sufami Turbo | MD5 (header-aware) |
// | Saturn                                  | Custom             |
// | Sega 32X                                | MD5                |
// | Sega CD                                 | Custom             |
// | Vectrex                                 | MD5                |
// | Virtual Boy                             | MD5                |
// | WASM-4                                  | MD5                |
// | Watara Supervision                      | MD5                |
// | Wellback Mega Duck                      | MD5                |

const (
	_NES    Platform = iota //Nintendo Entertainment System
	_GB                     //Game Boy
	_SNES                   //Super Nintendo Entertainment System
	_BSX                    //Satellaview
	_VB                     //Virtual Boy
	_N64                    //Nintendo 64
	_GBC                    //Game Boy Color
	_GBA                    //Game Boy Advance
	_GCN                    //Nintendo GameCube
	_NDS                    //Nintendo DS
	_Wii                    //Nintendo Wii
	_DSi                    //Nintendo DSi
	_3DS                    //Nintendo 3DS
	_WiiU                   //Nintendo Wii
	_SW                     //Nintendo Switch
	_SMS                    //Sega Master System
	_GEN                    //Sega Genesis
	_GG                     //Sega Game Gear
	_SCD                    //Sega CD
	_32X                    //Sega 32X
	_SAT                    //Sega Saturn
	_DC                     //Dreamcast
	_PS1                    //PlayStation
	_PS2                    //PlayStation 2
	_PS3                    //PlayStation 3
	_PS4                    //PlayStation 4
	_PS5                    //PlayStation 5
	_PSP                    //PlayStation Portable
	_PSVita                 //PlayStation Vita
	_XBOX                   //Xbox
	_XB360                  //Xbox 360
	_XBLA                   //Xbox Live Arcade
	_XBONE                  //Xbox One
	_3DO                    //3DO Interactive Multiplayer
	_ARCADE                 //Arcade
	_A2600                  //Atari 2600
	_A5200                  //Atari 5200
	_A7800                  //Atari 7800
	_AJ                     //Atari Jaguar
	_JCD                    //Atari Jaguar CD
	_AL                     //Atari Lynx
	_MOB                    //Mobile
	_NEO                    //Neo Geo
	_PC                     //PC
	_CD_i                   //Philips CD-i
	_TG_16                  //TurboGrafx-16
	_TGCD                   // PC Engine CD/TurboGrafx-CD
	_MUL                    //Multiplatform
)

var AllPlatforms = []struct {
	Value  Platform
	TSName string
}{
	{_NES, "NES"},
	{_GB, "_GB"},
	{_SNES, "SNES"},
	{_BSX, "_BSX"},
	{_VB, "_VB"},
	{_N64, "_N64"},
	{_GBC, "_GBC"},
	{_GBA, "_GBA"},
	{_GCN, "_GCN"},
	{_NDS, "_NDS"},
	{_Wii, "_Wii"},
	{_DSi, "_DSi"},
	{_3DS, "_3DS"},
	{_WiiU, "_WiiU"},
	{_SW, "_Switch"},
	{_SMS, "_SMS"},
	{_GEN, "_GEN"},
	{_GG, "_GG"},
	{_SCD, "_SCD"},
	{_32X, "_32X"},
	{_SAT, "_SAT"},
	{_DC, "_DC"},
	{_PS1, "_PS1"},
	{_PS2, "_PS2"},
	{_PS3, "_PS3"},
	{_PS4, "_PS4"},
	{_PS5, "_PS5"},
	{_PSP, "_PSP"},
	{_PSVita, "_PSVita"},
	{_XBOX, "_XBOX"},
	{_XB360, "_XB360"},
	{_XBLA, "_XBLA"},
	{_XBONE, "_XBONE"},
	{_3DO, "_3DO"},
	{_ARCADE, "_ARCADE"},
	{_A2600, "_A2600"},
	{_A5200, "_A5200"},
	{_A7800, "_A7800"},
	{_AJ, "_AJ"},
	{_JCD, "_JCD"},
	{_AL, "_AL"},
	{_MOB, "_MOB"},
	{_NEO, "_NEO"},
	{_PC, "_PC"},
	{_CD_i, "_CD_i"},
	{_TG_16, "_TG_16"},
	{_TGCD, "_TGCD"},
	{_MUL, "_MUL"},
}

func platfromForRA(platform Platform) int {
	var RAconsoles = map[Platform]int{
		_A2600: 25, //MD5
		_A5200: 50, //?
		_A7800: 51, //MD5
		_CD_i:  42, //No achiements
		_DC:    40, //Custom
		_GBA:   5,  //MD5
		_GBC:   6,  //MD5
		_GB:    4,  //MD5
		_GG:    15, //MD5
		_GCN:   16, //Custom
		_GEN:   1,  //MD5
		_JCD:   77, //Custom
		_AJ:    17, //MD5
		_AL:    13, //MD5
		_SMS:   11, //MD5
		_3DS:   62, //No achiements
		_N64:   2,  //MD5
		_NDS:   18, //Custom
		_PSP:   41, //Custom
		_PS2:   21, //Custom
		_PS1:   41, //Custom
		_SAT:   39, //Custom
		_SCD:   9,  //Custom
		_SNES:  3,  //MD5
		_TG_16: 8,  //MD5
		_TGCD:  76, //Custom
		_VB:    28, //MD5j
		_Wii:   19, //No achiements
		_XBOX:  22, //No achiements
	}
	return RAconsoles[platform]
}

func invertMap[K comparable, V comparable](m map[K]V) map[V]K {
	inv := make(map[V]K, len(m))
	for k, v := range m {
		inv[v] = k
	}
	return inv
}

var platformMap = map[string]Platform{
	// Nintendo
	"Nintendo Entertainment System":       _NES,
	"NES":                                 _NES,
	"Game Boy":                            _GB,
	"Super Nintendo":                      _SNES,
	"Super Nintendo Entertainment System": _SNES,
	"Satellaview":                         _BSX,
	"Virtual Boy":                         _VB,
	"Nintendo 64":                         _N64,
	"Game Boy Color":                      _GBC,
	"Game Boy Advance":                    _GBA,
	"Nintendo GameCube":                   _GCN,
	"GameCube":                            _GCN,
	"Nintendo DS":                         _NDS,
	"Nintendo DSi":                        _DSi,
	"Nintendo 3DS":                        _3DS,
	"Wii":                                 _Wii,
	"Wii U":                               _WiiU,
	"Nintendo Switch":                     _SW,

	// Sega
	"Sega Master System": _SMS,
	"Master System":      _SMS,
	"Sega Genesis":       _GEN,
	"Genesis":            _GEN,
	"Sega Game Gear":     _GG,
	"Game Gear":          _GG,
	"Sega CD":            _SCD,
	"Sega 32X":           _32X,
	"Sega Saturn":        _SAT,
	"Saturn":             _SAT,
	"Dreamcast":          _DC,

	// PlayStation
	"PlayStation":          _PS1,
	"PlayStation 2":        _PS2,
	"PlayStation 3":        _PS3,
	"PlayStation 4":        _PS4,
	"PlayStation 5":        _PS5,
	"PlayStation Portable": _PSP,
	"PS Portable":          _PSP,
	"PlayStation Vita":     _PSVita,

	// Xbox
	"Xbox":             _XBOX,
	"Xbox 360":         _XB360,
	"Xbox Live Arcade": _XBLA,
	"Xbox One":         _XBONE,

	// Atari
	"Atari 2600":      _A2600,
	"Atari 5200":      _A5200,
	"Atari 7800":      _A7800,
	"Atari Jaguar":    _AJ,
	"Jaguar":          _AJ,
	"Atari Jaguar CD": _JCD,
	"Jaguar CD":       _JCD,
	"Atari Lynx":      _AL,
	"Lynx":            _AL,

	// NEC
	"TurboGrafx-16":              _TG_16,
	"TG-16":                      _TG_16,
	"PC Engine":                  _TG_16,
	"TurboGrafx-CD":              _TGCD,
	"PC Engine CD":               _TGCD,
	"PC Engine CD/TurboGrafx-CD": _TGCD,

	// Other
	"Philips CD-i":                _CD_i,
	"CD-i":                        _CD_i,
	"Neo Geo":                     _NEO,
	"3DO":                         _3DO,
	"3DO Interactive Multiplayer": _3DO,
	"Arcade":                      _ARCADE,
	"PC":                          _PC,
	"Mobile":                      _MOB,
	"Multiplatform":               _MUL,
}

var nameMap = invertMap(platformMap)

func stringToPlatform(name string) Platform {
	return platformMap[name]
}

func PlatformToString(platform Platform) string {
	return nameMap[platform]
}
