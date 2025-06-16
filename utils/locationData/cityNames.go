package locationData

// Cities contains a map of cities and their localities.
// Each key is a city name, and the value is a slice of localities within that city.
// This data can be used to populate dropdowns or other UI components in web application.

var Cities = map[string][]string{
	"Dibrugarh": {"Dibrugarh", "Chabua", "Duliajan", "Moran", "Naharkatia",
		"Namrup", "Jamirah Patra Gaon", "Barbari", "Boiragimoth", "Dinjan",
		"Duliajan Oil Town", "Lepetkata BCPL township", "Sarupathar Bengali", "Sepon",
	},
	// "Guwahati": {"Dispur", "Beltola", "Bhangagarh", "Maligaon",
	// 	"Paltan Bazar", "Fancy Bazar", "Uzan Bazar", "Pan Bazaar", "Athgaon",
	// 	"Kahilipara", "Rukminigaon", "Bamunimaidam", "Narengi", "Borjhar",
	// },

	// Add more cities as needed
}
