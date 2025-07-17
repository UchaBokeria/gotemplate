package products

import (
	"main/internal/models"
	"main/internal/storage"
)

var Seed = []models.Product{
	// Air Filters
	{
		Name:        "Premium Engine Air Filter",
		Category:    "air-filters",
		Price:       24.99,
		Stock:       150,
		Status:      "active",
		Description: "High-performance engine air filter designed for optimal airflow and superior engine protection. Compatible with most vehicle makes and models.",
	},
	{
		Name:        "Cabin Air Filter",
		Category:    "air-filters",
		Price:       19.99,
		Stock:       200,
		Status:      "active",
		Description: "Premium cabin air filter that ensures clean air circulation inside your vehicle. Removes dust, pollen, and odors effectively.",
	},
	{
		Name:        "Performance Air Filter",
		Category:    "air-filters",
		Price:       39.99,
		Stock:       75,
		Status:      "active",
		Description: "High-flow performance air filter for enhanced engine performance and improved fuel efficiency. Perfect for performance enthusiasts.",
	},
	{
		Name:        "Heavy Duty Air Filter",
		Category:    "air-filters",
		Price:       34.99,
		Stock:       120,
		Status:      "active",
		Description: "Industrial-grade air filter designed for heavy-duty applications and extreme driving conditions.",
	},

	// Oil Filters
	{
		Name:        "Standard Oil Filter",
		Category:    "oil-filters",
		Price:       12.99,
		Stock:       300,
		Status:      "active",
		Description: "Reliable standard oil filter that provides excellent filtration and engine protection for everyday driving.",
	},
	{
		Name:        "Synthetic Oil Filter",
		Category:    "oil-filters",
		Price:       18.99,
		Stock:       250,
		Status:      "active",
		Description: "Advanced synthetic oil filter designed for use with synthetic oils. Extended service life and superior protection.",
	},
	{
		Name:        "High-Flow Oil Filter",
		Category:    "oil-filters",
		Price:       22.99,
		Stock:       180,
		Status:      "active",
		Description: "High-flow oil filter that maintains optimal oil pressure while providing excellent filtration efficiency.",
	},
	{
		Name:        "Extended Life Oil Filter",
		Category:    "oil-filters",
		Price:       16.99,
		Stock:       220,
		Status:      "active",
		Description: "Long-lasting oil filter designed for extended service intervals. Perfect for modern engines with extended oil change periods.",
	},

	// Fuel Filters
	{
		Name:        "Inline Fuel Filter",
		Category:    "fuel-filters",
		Price:       15.99,
		Stock:       160,
		Status:      "active",
		Description: "High-quality inline fuel filter that ensures clean fuel delivery to your engine. Easy installation and reliable performance.",
	},
	{
		Name:        "Fuel Pump Filter",
		Category:    "fuel-filters",
		Price:       28.99,
		Stock:       90,
		Status:      "active",
		Description: "Specialized fuel pump filter that protects your fuel pump and injection system from contamination.",
	},
	{
		Name:        "Diesel Fuel Filter",
		Category:    "fuel-filters",
		Price:       32.99,
		Stock:       110,
		Status:      "active",
		Description: "Heavy-duty diesel fuel filter designed for diesel engines. Removes water and contaminants effectively.",
	},
	{
		Name:        "High-Pressure Fuel Filter",
		Category:    "fuel-filters",
		Price:       24.99,
		Stock:       85,
		Status:      "active",
		Description: "High-pressure fuel filter suitable for modern fuel injection systems. Maintains optimal fuel pressure and cleanliness.",
	},

	// Brake Parts
	{
		Name:        "Brake Pads Set - Front",
		Category:    "brake-parts",
		Price:       89.99,
		Stock:       80,
		Status:      "active",
		Description: "Complete front brake pads set with excellent stopping power and long-lasting performance. Includes shims and hardware.",
	},
	{
		Name:        "Brake Pads Set - Rear",
		Category:    "brake-parts",
		Price:       79.99,
		Stock:       90,
		Status:      "active",
		Description: "High-quality rear brake pads designed for optimal braking performance and minimal noise.",
	},
	{
		Name:        "Brake Rotors - Front Pair",
		Category:    "brake-parts",
		Price:       129.99,
		Stock:       45,
		Status:      "active",
		Description: "Premium front brake rotors with superior heat dissipation and durability. Smooth and consistent braking performance.",
	},
	{
		Name:        "Brake Rotors - Rear Pair",
		Category:    "brake-parts",
		Price:       119.99,
		Stock:       50,
		Status:      "active",
		Description: "High-performance rear brake rotors engineered for maximum stopping power and longevity.",
	},
	{
		Name:        "Brake Fluid DOT 4",
		Category:    "brake-parts",
		Price:       14.99,
		Stock:       200,
		Status:      "active",
		Description: "Premium DOT 4 brake fluid that ensures reliable braking performance in all weather conditions.",
	},
	{
		Name:        "Brake Caliper Kit",
		Category:    "brake-parts",
		Price:       159.99,
		Stock:       30,
		Status:      "active",
		Description: "Complete brake caliper kit including all necessary seals and hardware for professional installation.",
	},

	// Engine Parts
	{
		Name:        "Spark Plugs Set (4pc)",
		Category:    "engine-parts",
		Price:       34.99,
		Stock:       120,
		Status:      "active",
		Description: "Premium spark plugs set for improved ignition and engine performance. Set of 4 plugs suitable for most 4-cylinder engines.",
	},
	{
		Name:        "Engine Oil 5W-30 Synthetic",
		Category:    "engine-parts",
		Price:       24.99,
		Stock:       150,
		Status:      "active",
		Description: "High-quality synthetic engine oil that provides superior protection and performance in all driving conditions.",
	},
	{
		Name:        "Timing Belt Kit",
		Category:    "engine-parts",
		Price:       89.99,
		Stock:       40,
		Status:      "active",
		Description: "Complete timing belt kit including belt, tensioners, and pulleys. Essential for proper engine timing maintenance.",
	},
	{
		Name:        "Engine Coolant",
		Category:    "engine-parts",
		Price:       19.99,
		Stock:       180,
		Status:      "active",
		Description: "Premium engine coolant that provides excellent heat transfer and corrosion protection for your cooling system.",
	},

	// Some inactive/draft products for testing filters
	{
		Name:        "Test Product - Inactive",
		Category:    "air-filters",
		Price:       99.99,
		Stock:       0,
		Status:      "inactive",
		Description: "This is an inactive product for testing purposes.",
	},
	{
		Name:        "Draft Product - Coming Soon",
		Category:    "oil-filters",
		Price:       45.99,
		Stock:       0,
		Status:      "draft",
		Description: "This product is in draft status and will be available soon.",
	},
}

func Populate() {
	for _, row := range Seed {
		storage.DB.Create(&row)
	}
}
