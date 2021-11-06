package datastructures

type Itemrequest struct {
	Items   []Item        `json:"items"`
	Request RequestStruct `json:"request"`
}

type Item struct {
	Name                  string   `json:"name"`
	Tier                  string   `json:"tier"`
	Type                  string   `json:"type"`
	Category              string   `json:"category"`
	Set                   *string  `json:"set"`
	Restrict              string   `json:"restrict"`
	Restrictions          *string  `json:"restrictions"`
	Drop                  string   `json:"drop"`
	HP                    int      `json:"hp"`
	Lvl                   int      `json:"lvl"`
	StrReq                int      `json:"strReq"`
	AgiReq                int      `json:"agiReq"`
	DefReq                int      `json:"defReq"`
	DexReq                int      `json:"dexReq"`
	IntReq                int      `json:"intReq"`
	Material              *string  `json:"material"`
	ArmorType             *string  `json:"armorType"`
	ArmorColor            *string  `json:"armorColor"`
	AddedLore             *string  `json:"addedLore"`
	Sockets               int      `json:"slots"`
	EarthDefense          int      `json:"eDef"`
	ThunderDefense        int      `json:"tDef"`
	WaterDefense          int      `json:"wDef"`
	FireDefense           int      `json:"fDef"`
	AirDefense            int      `json:"aDef"`
	Damage                string   `json:"nDam"`
	EarthDamage           string   `json:"eDamage"`
	ThunderDamage         string   `json:"tDamage"`
	WaterDamage           string   `json:"wDamage"`
	FireDamage            string   `json:"fDamage"`
	AirDamage             string   `json:"aDamage"`
	AttackSpeed           string   `json:"atkSpd"`
	Quest                 string   `json:"quest"`
	ClassRequirement      string   `json:"classReq"`
	StrengthPoints        int      `json:"strReqPoints"`
	DexterityPoints       int      `json:"dexReqPoints"`
	IntelligencePoints    int      `json:"intReqPoints"`
	DefensePoints         int      `json:"defReqPoints"`
	AgilityPoints         int      `json:"agiReqPoints"`
	MajorIds              []string `json:"majorIds,omitempty"`
	DamageBonus           int      `json:"nDamBonus"`
	DamageBonusRaw        int      `json:"nDamBonusRaw"`
	SpellDamage           int      `json:"sdPct"`
	SpellDamageRaw        int      `json:"sdPctRaw"`
	RainbowSpellDamageRaw int      `json:"rainbowRaw"`
	HealthRegen           int      `json:"hpRegen"`
	HealthRegenRaw        int      `json:"hpRegenRaw"`
	HealthBonus           int      `json:"hpBonus"`
	Poison                int      `json:"poison"`
	LifeSteal             int      `json:"ls"`
	ManaRegen             int      `json:"mr"`
	ManaSteal             int      `json:"ms"`
	SpellCostPct1         int      `json:"spPct1"`
	SpellCostRaw1         int      `json:"spRaw1"`
	SpellCostPct2         int      `json:"spPct2"`
	SpellCostRaw2         int      `json:"spRaw2"`
	SpellCostPct3         int      `json:"spPct3"`
	SpellCostRaw3         int      `json:"spRaw3"`
	SpellCostPct4         int      `json:"spPct4"`
	SpellCostRaw4         int      `json:"spRaw4"`
	Thorns                int      `json:"thorns"`
	Reflection            int      `json:"ref"`
	AttackSpeedBonus      int      `json:"atkSpdBonus"`
	Speed                 int      `json:"spd"`
	Exploding             int      `json:"expd"`
	SoulPoints            int      `json:"spRegen"`
	Sprint                int      `json:"sprint"`
	SprintRegen           int      `json:"sprintReg"`
	JumpHeight            int      `json:"jh"`
	XpBonus               int      `json:"xpb"`
	LootBonus             int      `json:"lb"`
	LootQuality           int      `json:"lq"`
	EmeraldStealing       int      `json:"eSteal"`
	GatherXpBonus         int      `json:"gXp"`
	GatherSpeed           int      `json:"gSpd"`
	BonusEarthDamage      int      `json:"eDamPct"`
	BonusThunderDamage    int      `json:"tDamPct"`
	BonusWaterDamage      int      `json:"wDamPct"`
	BonusFireDamage       int      `json:"fDamPct"`
	BonusAirDamage        int      `json:"aDamPct"`
	BonusEarthDefense     int      `json:"eDefPct"`
	BonusThunderDefense   int      `json:"tDefPct"`
	BonusWaterDefense     int      `json:"wDefPct"`
	BonusFireDefense      int      `json:"fDefPct"`
	BonusAirDefense       int      `json:"aDefPct"`
	DisplayName           string   `json:"displayName,omitempty"`
	Identified            bool     `json:"fixID,omitempty"`
	AllowCraftsman        bool     `json:"allowCraftsman,omitempty"`
}

type RequestStruct struct {
	Timestamp int64   `json:"timestamp"`
	Version   float64 `json:"version"`
}
