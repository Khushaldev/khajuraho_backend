package repository

import (
	db "khajuraho/backend/database"
	"khajuraho/backend/model"

	"github.com/google/uuid"
)

//	var discoverItems = []model.DiscoverItem{
//		{
//			Name:         "Barbeque Nation",
//			Slug:         "barbeque-nation-dwarka",
//			CategoryID:   uuidPtr("232330a5-aa59-4a49-9289-efb531d7fc58"),
//			CityID:       uuidPtr("213b2410-2972-436a-be85-1303b9405bd1"),
//			Geopoint:     utils.ToGeoPoint(28.5843, 77.0151),
//			Description:  "Famous for grill-on-the-table buffet experience.",
//			Images:       pq.StringArray{"https://example.com/images/barbeque.jpg"},
//			Tags:         pq.StringArray{"buffet", "grill", "family"},
//			Rating:       4.5,
//			ReviewsCount: 2100,
//			Metadata:     utils.ToJSON(`{"cuisine": "North Indian", "price_range": "₹₹₹"}`),
//		},
//		{
//			Name:         "Biryani Blues",
//			Slug:         "biryani-blues-dwarka",
//			CategoryID:   uuidPtr("232330a5-aa59-4a49-9289-efb531d7fc58"),
//			CityID:       uuidPtr("213b2410-2972-436a-be85-1303b9405bd1"),
//			Geopoint:     utils.ToGeoPoint(28.5843, 77.0151),
//			Description:  "Authentic Hyderabadi biryani with great combo meals.",
//			Images:       pq.StringArray{"https://example.com/images/biryani-blues.jpg"},
//			Tags:         pq.StringArray{"biryani", "hyderabadi", "combo"},
//			Rating:       4.2,
//			ReviewsCount: 1200,
//			Metadata:     utils.ToJSON(`{"cuisine": "Hyderabadi", "price_range": "₹₹"}`),
//		},
//		{
//			Name:         "Pizza Hut",
//			Slug:         "pizza-hut-dwarka",
//			CategoryID:   uuidPtr("232330a5-aa59-4a49-9289-efb531d7fc58"),
//			CityID:       uuidPtr("213b2410-2972-436a-be85-1303b9405bd1"),
//			Geopoint:     utils.ToGeoPoint(28.5843, 77.0151),
//			Description:  "Global pizza chain with dine-in and delivery.",
//			Images:       pq.StringArray{"https://example.com/images/pizzahut.jpg"},
//			Tags:         pq.StringArray{"pizza", "cheese", "fast food"},
//			Rating:       4.0,
//			ReviewsCount: 1800,
//			Metadata:     utils.ToJSON(`{"cuisine": "Italian-American", "price_range": "₹₹"}`),
//		},
//		{
//			Name:         "McDonald's",
//			Slug:         "mcdonalds-dwarka",
//			CategoryID:   uuidPtr("232330a5-aa59-4a49-9289-efb531d7fc58"),
//			CityID:       uuidPtr("213b2410-2972-436a-be85-1303b9405bd1"),
//			Geopoint:     utils.ToGeoPoint(28.5843, 77.0151),
//			Description:  "Classic burgers, fries, and beverages.",
//			Images:       pq.StringArray{"https://example.com/images/mcdonalds.jpg"},
//			Tags:         pq.StringArray{"burgers", "fries", "kids"},
//			Rating:       4.1,
//			ReviewsCount: 2400,
//			Metadata:     utils.ToJSON(`{"cuisine": "American", "price_range": "₹₹"}`),
//		},
//		{
//			Name:         "Haldiram's",
//			Slug:         "haldirams-dwarka",
//			CategoryID:   uuidPtr("232330a5-aa59-4a49-9289-efb531d7fc58"),
//			CityID:       uuidPtr("213b2410-2972-436a-be85-1303b9405bd1"),
//			Geopoint:     utils.ToGeoPoint(28.5843, 77.0151),
//			Description:  "North Indian snacks, thalis and sweets.",
//			Images:       pq.StringArray{"https://example.com/images/haldirams.jpg"},
//			Tags:         pq.StringArray{"veg", "sweets", "thali"},
//			Rating:       4.3,
//			ReviewsCount: 3000,
//			Metadata:     utils.ToJSON(`{"cuisine": "North Indian", "price_range": "₹₹"}`),
//		},
//		{
//			Name:         "The Belgian Waffle Co.",
//			Slug:         "belgian-waffle-dwarka",
//			CategoryID:   uuidPtr("232330a5-aa59-4a49-9289-efb531d7fc58"),
//			CityID:       uuidPtr("213b2410-2972-436a-be85-1303b9405bd1"),
//			Geopoint:     utils.ToGeoPoint(28.5843, 77.0151),
//			Description:  "Hot waffles and dessert combos.",
//			Images:       pq.StringArray{"https://example.com/images/waffle.jpg"},
//			Tags:         pq.StringArray{"dessert", "waffle", "sweet"},
//			Rating:       4.4,
//			ReviewsCount: 900,
//			Metadata:     utils.ToJSON(`{"cuisine": "Dessert", "price_range": "₹₹"}`),
//		},
//		{
//			Name:         "Domino's Pizza",
//			Slug:         "dominos-dwarka",
//			CategoryID:   uuidPtr("232330a5-aa59-4a49-9289-efb531d7fc58"),
//			CityID:       uuidPtr("213b2410-2972-436a-be85-1303b9405bd1"),
//			Geopoint:     utils.ToGeoPoint(28.5843, 77.0151),
//			Description:  "Fast and fresh pizzas with cheese burst options.",
//			Images:       pq.StringArray{"https://example.com/images/dominos.jpg"},
//			Tags:         pq.StringArray{"pizza", "delivery", "fast food"},
//			Rating:       4.0,
//			ReviewsCount: 2200,
//			Metadata:     utils.ToJSON(`{"cuisine": "Pizza", "price_range": "₹₹"}`),
//		},
//		{
//			Name:         "Berco's",
//			Slug:         "bercos-dwarka",
//			CategoryID:   uuidPtr("232330a5-aa59-4a49-9289-efb531d7fc58"),
//			CityID:       uuidPtr("213b2410-2972-436a-be85-1303b9405bd1"),
//			Geopoint:     utils.ToGeoPoint(28.5843, 77.0151),
//			Description:  "Popular for Chinese and Thai food.",
//			Images:       pq.StringArray{"https://example.com/images/bercos.jpg"},
//			Tags:         pq.StringArray{"chinese", "thai", "dine-in"},
//			Rating:       4.3,
//			ReviewsCount: 1500,
//			Metadata:     utils.ToJSON(`{"cuisine": "Asian", "price_range": "₹₹₹"}`),
//		},
//		{
//			Name:         "Subway",
//			Slug:         "subway-dwarka",
//			CategoryID:   uuidPtr("232330a5-aa59-4a49-9289-efb531d7fc58"),
//			CityID:       uuidPtr("213b2410-2972-436a-be85-1303b9405bd1"),
//			Geopoint:     utils.ToGeoPoint(28.5843, 77.0151),
//			Description:  "Customizable sandwiches and healthy wraps.",
//			Images:       pq.StringArray{"https://example.com/images/subway.jpg"},
//			Tags:         pq.StringArray{"healthy", "sandwich", "wraps"},
//			Rating:       4.1,
//			ReviewsCount: 1100,
//			Metadata:     utils.ToJSON(`{"cuisine": "Fast Food", "price_range": "₹₹"}`),
//		},
//		{
//			Name:         "Wow! Momo",
//			Slug:         "wow-momo-dwarka",
//			CategoryID:   uuidPtr("232330a5-aa59-4a49-9289-efb531d7fc58"),
//			CityID:       uuidPtr("213b2410-2972-436a-be85-1303b9405bd1"),
//			Geopoint:     utils.ToGeoPoint(28.5843, 77.0151),
//			Description:  "Specialty momos and fusion snacks.",
//			Images:       pq.StringArray{"https://example.com/images/wowmomo.jpg"},
//			Tags:         pq.StringArray{"momo", "snack", "fusion"},
//			Rating:       4.2,
//			ReviewsCount: 950,
//			Metadata:     utils.ToJSON(`{"cuisine": "Tibetan", "price_range": "₹₹"}`),
//		},
//		{
//			Name:         "Cafe Coffee Day",
//			Slug:         "ccd-dwarka",
//			CategoryID:   uuidPtr("232330a5-aa59-4a49-9289-efb531d7fc58"),
//			CityID:       uuidPtr("213b2410-2972-436a-be85-1303b9405bd1"),
//			Geopoint:     utils.ToGeoPoint(28.5843, 77.0151),
//			Description:  "Coffee, snacks and chill vibes.",
//			Images:       pq.StringArray{"https://example.com/images/ccd.jpg"},
//			Tags:         pq.StringArray{"coffee", "snacks", "cafe"},
//			Rating:       4.0,
//			ReviewsCount: 1050,
//			Metadata:     utils.ToJSON(`{"cuisine": "Cafe", "price_range": "₹₹"}`),
//		},
//		{
//			Name:         "Burger King",
//			Slug:         "burger-king-dwarka",
//			CategoryID:   uuidPtr("232330a5-aa59-4a49-9289-efb531d7fc58"),
//			CityID:       uuidPtr("213b2410-2972-436a-be85-1303b9405bd1"),
//			Geopoint:     utils.ToGeoPoint(28.5843, 77.0151),
//			Description:  "Whoppers, fries and quick service.",
//			Images:       pq.StringArray{"https://example.com/images/burgerking.jpg"},
//			Tags:         pq.StringArray{"burger", "fast food", "combo"},
//			Rating:       4.1,
//			ReviewsCount: 1700,
//			Metadata:     utils.ToJSON(`{"cuisine": "American", "price_range": "₹₹"}`),
//		},
//		{
//			Name:         "Theobroma",
//			Slug:         "theobroma-dwarka",
//			CategoryID:   uuidPtr("232330a5-aa59-4a49-9289-efb531d7fc58"),
//			CityID:       uuidPtr("213b2410-2972-436a-be85-1303b9405bd1"),
//			Geopoint:     utils.ToGeoPoint(28.5843, 77.0151),
//			Description:  "Cakes, brownies, pastries and more.",
//			Images:       pq.StringArray{"https://example.com/images/theobroma.jpg"},
//			Tags:         pq.StringArray{"bakery", "dessert", "cake"},
//			Rating:       4.6,
//			ReviewsCount: 1300,
//			Metadata:     utils.ToJSON(`{"cuisine": "Bakery", "price_range": "₹₹₹"}`),
//		},
//		{
//			Name:         "KFC",
//			Slug:         "kfc-dwarka",
//			CategoryID:   uuidPtr("232330a5-aa59-4a49-9289-efb531d7fc58"),
//			CityID:       uuidPtr("213b2410-2972-436a-be85-1303b9405bd1"),
//			Geopoint:     utils.ToGeoPoint(28.5843, 77.0151),
//			Description:  "Crispy fried chicken and burgers.",
//			Images:       pq.StringArray{"https://example.com/images/kfc.jpg"},
//			Tags:         pq.StringArray{"chicken", "fast food", "combo"},
//			Rating:       4.0,
//			ReviewsCount: 2000,
//			Metadata:     utils.ToJSON(`{"cuisine": "American", "price_range": "₹₹"}`),
//		},
//		{
//			Name:         "Tandoori Nights",
//			Slug:         "tandoori-nights-dwarka",
//			CategoryID:   uuidPtr("232330a5-aa59-4a49-9289-efb531d7fc58"),
//			CityID:       uuidPtr("213b2410-2972-436a-be85-1303b9405bd1"),
//			Geopoint:     utils.ToGeoPoint(28.5843, 77.0151),
//			Description:  "Tandoori platters and North Indian food.",
//			Images:       pq.StringArray{"https://example.com/images/tandoori.jpg"},
//			Tags:         pq.StringArray{"tandoori", "north indian", "dine-in"},
//			Rating:       4.3,
//			ReviewsCount: 870,
//			Metadata:     utils.ToJSON(`{"cuisine": "North Indian", "price_range": "₹₹₹"}`),
//		},
//	}
func FindCategories() ([]model.Category, error) {
	var categories []model.Category

	// location := []model.Address{
	// 	{
	// 		Address:   "Dwarka Sector 12",
	// 		City:      "Delhi",
	// 		State:     "Delhi",
	// 		Country:   "India",
	// 		Latitude:  28.5843,
	// 		Longitude: 77.0151,
	// 	},
	// 	{
	// 		Address:   "Dwarka Sector 10",
	// 		City:      "Delhi",
	// 		State:     "Delhi",
	// 		Country:   "India",
	// 		Latitude:  28.5868,
	// 		Longitude: 77.0264,
	// 	},
	// 	{
	// 		Address:   "Dwarka Sector 6",
	// 		City:      "Delhi",
	// 		State:     "Delhi",
	// 		Country:   "India",
	// 		Latitude:  28.5895,
	// 		Longitude: 77.0462,
	// 	},
	// 	{
	// 		Address:   "Dwarka Sector 7",
	// 		City:      "Delhi",
	// 		State:     "Delhi",
	// 		Country:   "India",
	// 		Latitude:  28.5875,
	// 		Longitude: 77.0437,
	// 	},
	// 	{
	// 		Address:   "Dwarka Sector 5",
	// 		City:      "Delhi",
	// 		State:     "Delhi",
	// 		Country:   "India",
	// 		Latitude:  28.5908,
	// 		Longitude: 77.0479,
	// 	},
	// 	{
	// 		Address:   "Dwarka Sector 21",
	// 		City:      "Delhi",
	// 		State:     "Delhi",
	// 		Country:   "India",
	// 		Latitude:  28.5535,
	// 		Longitude: 77.0140,
	// 	},
	// 	{
	// 		Address:   "Dwarka Sector 3",
	// 		City:      "Delhi",
	// 		State:     "Delhi",
	// 		Country:   "India",
	// 		Latitude:  28.5972,
	// 		Longitude: 77.0362,
	// 	},
	// 	{
	// 		Address:   "Dwarka Sector 13",
	// 		City:      "Delhi",
	// 		State:     "Delhi",
	// 		Country:   "India",
	// 		Latitude:  28.5819,
	// 		Longitude: 77.0138,
	// 	},
	// 	{
	// 		Address:   "Dwarka Sector 4",
	// 		City:      "Delhi",
	// 		State:     "Delhi",
	// 		Country:   "India",
	// 		Latitude:  28.5932,
	// 		Longitude: 77.0493,
	// 	},
	// 	{
	// 		Address:   "Dwarka Sector 11",
	// 		City:      "Delhi",
	// 		State:     "Delhi",
	// 		Country:   "India",
	// 		Latitude:  28.5802,
	// 		Longitude: 77.0260,
	// 	},
	// 	{
	// 		Address:   "Dwarka Sector 14",
	// 		City:      "Delhi",
	// 		State:     "Delhi",
	// 		Country:   "India",
	// 		Latitude:  28.5781,
	// 		Longitude: 77.0154,
	// 	},
	// 	{
	// 		Address:   "Dwarka Sector 8",
	// 		City:      "Delhi",
	// 		State:     "Delhi",
	// 		Country:   "India",
	// 		Latitude:  28.5603,
	// 		Longitude: 77.0511,
	// 	},
	// 	{
	// 		Address:   "Dwarka Sector 9",
	// 		City:      "Delhi",
	// 		State:     "Delhi",
	// 		Country:   "India",
	// 		Latitude:  28.5652,
	// 		Longitude: 77.0409,
	// 	},
	// 	{
	// 		Address:   "Dwarka Sector 1",
	// 		City:      "Delhi",
	// 		State:     "Delhi",
	// 		Country:   "India",
	// 		Latitude:  28.6005,
	// 		Longitude: 77.0444,
	// 	},
	// 	{
	// 		Address:   "Dwarka Sector 18",
	// 		City:      "Delhi",
	// 		State:     "Delhi",
	// 		Country:   "India",
	// 		Latitude:  28.5667,
	// 		Longitude: 77.0204,
	// 	},
	// }

	// db.Instance.Create(&districts)
	err := db.Instance.Where("is_active = true").Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func uuidPtr(id string) uuid.UUID {
	u := uuid.MustParse(id)
	return u
}

// func utils.ToJSON(s string) datatypes.JSON {
// 	j, _ := datatypes.JSON([]byte(s), nil)
// 	return j
// }

//	func ptrUUID(id string) *uuid.UUID {
//		u := uuid.MustParse(id)
//		return &u
//	}
