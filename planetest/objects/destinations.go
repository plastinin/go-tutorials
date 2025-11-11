package objects

func getDestinationsMap() map[string]uint {
	resultMap := make(map[string]uint)
	resultMap["Москва"] = 0
	resultMap["Череповец"] = 800
	resultMap["Екатеринбург"] = 1_200
	resultMap["Санкт-Петербург"] = 700
	resultMap["Анталия"] = 5_000

	return resultMap
}

func GetDistanceFromTo(From string, To string) uint {
	destinationMap := getDestinationsMap()
	fromDistance := destinationMap[From]
	toDistance := destinationMap[To]

	return fromDistance + toDistance
}
