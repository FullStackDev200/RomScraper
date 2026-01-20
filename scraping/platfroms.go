package scraping

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

type Platform int

const (
	_A26    Platform = iota //Atari 2600
	_A78                    //Atari 7800
	_NES                    //Nintendo Entertainment System
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
	_AL                     //Atari Lynx
	_MOB                    //Mobile
	_NEO                    //Neo Geo
	_PC                     //PC
	_CD_i                   //Philips CD-i
	_TG_16                  //TurboGrafx-16
	_MUL                    //Multiplatform
)
