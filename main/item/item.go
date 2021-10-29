package item

type Item struct {
	Name                  string `json:"name"`
	DisplayName           string `json:"displayName"`
	Tier                  string `json:"tier"`
	Set                   string `json:"set"`
	Sockets               int    `json:"sockets"`
	Type                  string `json:"type"`
	ArmorType             string `json:"armorType"`
	ArmorColor            string `json:"armorColor"`
	AddedLore             string `json:"addedLore"`
	Material              string `json:"material"`
	DropType              string `json:"dropType"`
	Restrictions          string `json:"restrictions"`
	Damage                string `json:"damage"`
	FireDamage            string `json:"fireDamage"`
	WaterDamage           string `json:"waterDamage"`
	AirDamage             string `json:"airDamage"`
	ThunderDamage         string `json:"thunderDamage"`
	EarthDamage           string `json:"earthDamage"`
	AttackSpeed           string `json:"attackSpeed"`
	Health                int    `json:"health"`
	FireDefense           int    `json:"fireDefense"`
	WaterDefense          int    `json:"waterDefense"`
	AirDefense            int    `json:"airDefense"`
	ThunderDefense        int    `json:"thunderDefense"`
	EarthDefense          int    `json:"earthDefense"`
	Level                 int    `json:"level"`
	Quest                 string `json:"quest"`
	ClassRequirement      string `json:"classRequirement"`
	Strength              int    `json:"strength"`
	Dexterity             int    `json:"dexterity"`
	Intelligence          int    `json:"intelligence"`
	Agility               int    `json:"agility"`
	Defense               int    `json:"defense"`
	HealthRegen           int    `json:"healthRegen"`
	ManaRegen             int    `json:"manaRegen"`
	SpellDamage           int    `json:"spellDamage"`
	RainbowSpellDamageRaw int    `json:"rainbowSpellDamageRaw"`
	DamageBonus           int    `json:"damageBonus"`
	LifeSteal             int    `json:"lifeSteal"`
	ManaSteal             int    `json:"manaSteal"`
	SpellCostPct1         int    `json:"spellCostPct1"`
	SpellCostRaw1         int    `json:"spellCostRaw1"`
	SpellCostPct2         int    `json:"spellCostPct2"`
	SpellCostRaw2         int    `json:"spellCostRaw2"`
	SpellCostPct3         int    `json:"spellCostPct3"`
	SpellCostRaw3         int    `json:"spellCostRaw3"`
	SpellCostPct4         int    `json:"spellCostPct4"`
	SpellCostRaw4         int    `json:"spellCostRaw4"`
	XpBonus               int    `json:"xpBonus"`
	LootBonus             int    `json:"lootBonus"`
	LootQuality           int    `json:"lootQuality"`
	GatherXpBonus         int    `json:"gatherXpBonus"`
	GatherSpeed           int    `json:"gatherSpeed"`
	Reflection            int    `json:"reflection"`
	StrengthPoints        int    `json:"strengthPos"`
	DexterityPoints       int    `json:"dexterityPos"`
	IntelligencePoints    int    `json:"intelligencePos"`
	AgilityPoints         int    `json:"agilityPos"`
	DefensePoints         int    `json:"defensePos"`
	Thorns                int    `json:"thorns"`
	Exploding             int    `json:"exploding"`
	Speed                 int    `json:"speed"`
	AttackSpeedBonus      int    `json:"attackSpeedBonus"`
	Poison                int    `json:"poison"`
	HealthBonus           int    `json:"healthBonus"`
	SoulPoints            int    `json:"soulPos"`
	Sprint                int    `json:"spr"`
	SprintRegen           int    `json:"sprRegen"`
	JumpHeight            int    `json:"jumpHeight"`
	EmeraldStealing       int    `json:"emeraldStealing"`
	HealthRegenRaw        int    `json:"healthRegenRaw"`
	SpellDamageRaw        int    `json:"spellDamageRaw"`
	DamageBonusRaw        int    `json:"damageBonusRaw"`
	BonusFireDamage       int    `json:"bonusFireDamage"`
	BonusWaterDamage      int    `json:"bonusWaterDamage"`
	BonusAirDamage        int    `json:"bonusAirDamage"`
	BonusThunderDamage    int    `json:"bonusThunderDamage"`
	BonusEarthDamage      int    `json:"bonusEarthDamage"`
	BonusFireDefense      int    `json:"bonusFireDefense"`
	BonusWaterDefense     int    `json:"bonusWaterDefense"`
	BonusAirDefense       int    `json:"bonusAirDefense"`
	BonusThunderDefense   int    `json:"bonusThunderDefense"`
	BonusEarthDefense     int    `json:"bonusEarthDefense"`
	AccessoryType         string `json:"accessoryType"`
	Identified            bool   `json:"identified"`
	Skin                  string `json:"skin"`
	Category              string `json:"category"`
}

type List []Item

type Request struct{
	Timestamp int64     `json:"timestamp"`
	Version   float32 `json:"version"`
}

type DictWithItemList struct {
	Items   List    `json:"items"`
	Request Request `json:"request"`
}
