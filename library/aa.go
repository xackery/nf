package library

func AACategory(val int) string {
	switch val {
	case 1:
		return "Passive"
	case 2:
		return "Progression"
	case 3:
		return "Shroud Passive"
	case 4:
		return "Shroud Active"
	case 5:
		return "Veteran Reward"
	case 6:
		return "Tradeskill"
	case 7:
		return "Expendable"
	case 8:
		return "Racial Innate"
	case 9:
		return "Everquest"
	default:
		return "None"
	}
}

func AAType(val int) string {
	switch val {
	case 1:
		return "General"
	case 2:
		return "Archetype"
	case 3:
		return "Class"
	case 4:
		return "PoP Advanced"
	case 5:
		return "PoP Abilities"
	case 6:
		return "Gates of Discord"
	case 7:
		return "Omens of War"
	case 8:
		return "Veteran"
	case 9:
		return "Dragons of Norrath"
	case 10:
		return "Depths of Darkhollow"
	default:
		return "Not Applicable"
	}
}
