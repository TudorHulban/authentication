package helpers

import "bytes"

func ProcessFormURLEncoded(response []byte) map[string]string {
	result := make(map[string]string)

	var buffer []rune

	modeKey := true

	var key, value string

	runeEqual := []rune("=")
	runeAmpersand := []rune("&")

	for ix, digit := range bytes.Runes(response) {
		if digit == runeEqual[0] {
			if modeKey {
				key = string(buffer)

				modeKey = false
				buffer = []rune{}
			}

			continue
		}

		if digit == runeAmpersand[0] {
			if !modeKey {
				value = string(buffer)

				modeKey = true
				buffer = []rune{}

				if len(key) > 0 && len(value) > 0 {
					result[key] = value
				}

				key = ""
				value = ""
			}

			continue
		}

		if ix == len(bytes.Runes(response))-1 {
			if !modeKey {
				buffer = append(buffer, digit)

				value = string(buffer)

				if len(key) > 0 && len(value) > 0 {
					result[key] = value
				}
			}

			break
		}

		buffer = append(buffer, digit)
	}

	return result
}
