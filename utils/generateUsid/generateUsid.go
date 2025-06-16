package generateUsid

import (
	"YesFix/utils/locationData"

	"github.com/jaevor/go-nanoid"
)

func GenerateUsid(city string) string {
	cityCode := locationData.CityCodes[city]

	decenaryID, err := nanoid.CustomUnicode("0123456789abcdefghijklmnopqrstuvwxyz", 10)
	if err != nil {
		panic(err)
	}

	id := decenaryID()

	usid := cityCode + "" + id

	return usid
}
