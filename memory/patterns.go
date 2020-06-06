package memory

//NewPatterns is Base osu signatures stuct (for tdeo's wrapper)
type NewPatterns struct {
	BPM           uint64 `sig:"8B 40 08 89 86 4C 01 00 00 C6"` //-0x4
	Base          uint64 `sig:"F8 01 74 04 83 65"`             //-0xC
	InMenuMods    uint64 `sig:"C8 FF ?? ?? ?? ?? ?? 81 0D ?? ?? ?? ?? 00 08 00 00"`
	PlayTime      uint64 `sig:"5E 5F 5D C3 A1 ?? ?? ?? ?? 89 ?? 04"` //+0x5
	PlayContainer uint64 `sig:"85 C9 74 1F 8D 55 F0 8B 01"`
	LeaderBoard   uint64 `sig:"A1 ?? ?? ?? ?? 8B 50 04 8B 0D"` //+0x1
	SongsFolder   uint64 `sig:"?? ?? 67 ?? 2F 00 28 00"`
}

//StaticAddresses (should be updated every client restart)
type StaticAddresses struct {
	Status        uint32
	BPM           uint32
	Base          uint32
	InMenuMods    uint32
	PlayTime      uint32
	PlayContainer uint32
	LeaderBoard   uint32
	SongsFolder   uint32
}

//osuStaticAddresses (should be updated every client restart)
var osuStaticAddresses = StaticAddresses{}
