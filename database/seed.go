package database

import (
	"github.com/lionelritchie29/recipitor-be/models"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	db.Where("name = ?", "Apple").FirstOrCreate(&models.Item{
		Name: "Apple",
		Description: "Apple fruit",
		Image: "https://i.ibb.co/bg7wK41/apple-158989157.jpg",
	})

	db.Where("name = ?", "Chicken Breast").FirstOrCreate(&models.Item{
		Name: "Chicken Breast",
		Description: "Breast of chicken",
		Image: "https://i.ibb.co/bRnWYMm/Chicken-Breast-Boneless-3-4-Pieces-Hero-Shot-1.jpg",
	})

	db.Where("name = ?", "Banana").FirstOrCreate(&models.Item{
		Name: "Banana",
		Description: "Banana fruit",
		Image: "https://i.ibb.co/mz8n7Lw/Banana-Single-1.jpg",
	})

	db.Where("name = ?", "Peanut Butter").FirstOrCreate(&models.Item{
		Name: "Peanut Butter",
		Description: "Peanut Butter",
		Image: "https://i.ibb.co/58M8mFT/noccio-noccio-peanut-butter-creamy-500-g-full01-otbz2ln6-1.jpg",
	})

	db.Where("name = ?", "Egg").FirstOrCreate(&models.Item{
		Name: "Egg",
		Description: "an oval or round object laid by a female bird, reptile, fish, or invertebrate, usually containing a developing embryo. The eggs of birds are enclosed in a chalky shell, while those of reptiles are in a leathery membrane.",
		Image: "https://i.ibb.co/dKW7sTz/15xp-egg-promo-super-Jumbo-v2.jpg",
	})

	db.Where("name = ?", "Milk").FirstOrCreate(&models.Item{
		Name: "Milk",
		Description: "an opaque white or bluish-white liquid secreted by the mammary glands of female mammals, serving for the nourishment of their young",
		Image: "https://i.ibb.co/xXLjCtS/original-almond-beverage-189l-en.png",
	})

	db.Where("name = ?", "Watermelon").FirstOrCreate(&models.Item{
		Name: "Watermelon",
		Description: "Watermelon is a flowering plant species of the Cucurbitaceae family and the name of its edible fruit.",
		Image: "https://i.ibb.co/d0ynx4b/61r5l3-Ml7-L-SL1500.jpg",
	})

	db.Where("name = ?", "Almond").FirstOrCreate(&models.Item{
		Name: "Almond",
		Description: "The almond is a species of tree native to Iran and surrounding countries, including the Levant.",
		Image: "https://i.ibb.co/s3Rf4c4/content-Image.jpg",
	})

	db.Where("name = ?", "Chocolate Chip").FirstOrCreate(&models.Item{
		Name: "Chocolate Chip",
		Description: "Chocolate chips or chocolate morsels are small chunks of sweetened chocolate, used as an ingredient in a number of desserts, in trail mix and less commonly in some breakfast foods such as pancakes.",
		Image: "https://i.ibb.co/ZWTF5yS/download.jpg",
	})

	db.Where("name = ?", "Plain Loaf").FirstOrCreate(&models.Item{
		Name: "Plain Loaf",
		Description: "A plain loaf, slices of which are known in Scots as plain breid, is a traditional style of loaf made chiefly in Scotland and Ireland. It has a dark, well-fired crust on the top and bottom of the bread.",
		Image: "https://i.ibb.co/4KsyjW4/k-Photo-Recipe-Ramp-Up-2021-11-Potato-Bread-potato-bread-01.jpg",
	})

	db.Where("name = ?", "Orange").FirstOrCreate(&models.Item{
		Name: "Orange",
		Description: "An orange is a fruit of various citrus species in the family Rutaceae; it primarily refers to Citrus × sinensis, which is also called sweet orange, to distinguish it from the related Citrus × aurantium, referred to as bitter orange.",
		Image: "https://i.ibb.co/xSkzJsL/6000191272346-R.jpg",
	})

	db.Where("name = ?", "Flour").FirstOrCreate(&models.Item{
		Name: "Flour",
		Description: "Flour is a powder made by grinding raw grains, roots, beans, nuts, or seeds. Flours are used to make many different foods.",
		Image: "https://i.ibb.co/MnNs4M7/flour-in-bowl-bread-flour-vs-all-purpose-flour-1611940896.jpg",
	})

	db.Where("name = ?", "Green Tea").FirstOrCreate(&models.Item{
		Name: "Green Tea",
		Description: "Green tea is a type of tea that is made from Camellia sinensis leaves and buds that have not undergone the same withering and oxidation process which is used to make oolong teas and black teas.",
		Image: "https://i.ibb.co/cxkYZM6/green-tea-wight-loss-1643990040.jpg",
	})

	db.Where("name = ?", "White Rice").FirstOrCreate(&models.Item{
		Name: "White Rice",
		Description: "White rice is milled rice that has had its husk, bran, and germ removed.",
		Image: "https://i.ibb.co/6F81fMK/How-to-boil-rice-square-FS-6126.jpg",
	})

	db.Where("name = ?", "Chicken Tight").FirstOrCreate(&models.Item{
		Name: "Chicken Tight",
		Description: "Part of chicken's legs.",
		Image: "https://i.ibb.co/2spLX43/images-1.jpg",
	})

	db.Where("name = ?", "Sugar").FirstOrCreate(&models.Item{
		Name: "Sugar",
		Description: "Sugar is the generic name for sweet-tasting, soluble carbohydrates, many of which are used in food.",
		Image: "https://i.ibb.co/mC5Mnx2/images-2.jpg",
	})

	db.Where("name = ?", "Cheese").FirstOrCreate(&models.Item{
		Name: "Cheese",
		Description: "Cheese is a dairy product produced in wide ranges of flavors, textures, and forms by coagulation of the milk protein casein. It comprises proteins and fat from milk.",
		Image: "https://i.ibb.co/grw4sF3/images-3.jpg",
	})

	db.Where("name = ?", "Lemon").FirstOrCreate(&models.Item{
		Name: "Lemon",
		Description: "The lemon is a species of small evergreen trees in the flowering plant family Rutaceae, native to Asia, primarily Northeast India, Northern Myanmar or China.",
		Image: "https://i.ibb.co/7XMV5Wr/Picture-of-lemons-cut-in-halves.jpg",
	})

	db.Where("name = ?", "Strawberry").FirstOrCreate(&models.Item{
		Name: "Strawberry",
		Description: "The garden strawberry is a widely grown hybrid species of the genus Fragaria, collectively known as the strawberries, which are cultivated worldwide for their fruit.",
		Image: "https://i.ibb.co/cw6W5h5/square-1432664914-strawberry-facts1.jpg",
	})

	db.Where("name = ?", "Red Chilli Pepper").FirstOrCreate(&models.Item{
		Name: "Red Chilli Pepper",
		Description: "Chili peppers are widely used in many cuisines as a spice to add 'heat' to dishes.",
		Image: "https://i.ibb.co/q935sxS/download-1.jpg",
	})

}