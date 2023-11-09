package piscine

func PodiumPosition(podium [][]string) [][]string {
	if len(podium) == 0 {
		return [][]string{}
	}
	for i := 0; i < len(podium)-1; i++ {
		for j := 0; j < len(podium)-1-i; j++ {
			if len(podium[j]) > len(podium[j+1]) {
				podium[j], podium[j+1] = podium[j+1], podium[j]
			}
		}
	}
	return podium
}
