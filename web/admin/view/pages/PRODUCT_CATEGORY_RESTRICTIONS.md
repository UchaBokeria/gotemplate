# Product-Category Restrictions Implementation

This document outlines the implementation of product-category restrictions that ensure products can only be associated with level 4 (parts) categories.

## Requirements Implemented

### A. Level 4 Category Restriction
Products can now only be associated with level 4 categories (parts level). Any attempt to associate products with categories of other levels will result in an error.

### B. Enhanced Admin Category Selector
The admin product form now shows only level 4 categories with full path information and enhanced search capabilities.

## Implementation Details

### 1. Database Model Validation

**File**: `internal/models/categories.model.go`

Added GORM hooks to the `ProductCategory` model:
- `BeforeCreate()`: Validates category level before creating associations
- `BeforeUpdate()`: Validates category level before updating associations

```go
func (pc *ProductCategory) BeforeCreate(tx *gorm.DB) error {
    // Validates that category.Level == 4
}
```

### 2. Controller-Level Validation

**File**: `web/admin/controllers/products/products.controller.go`

Added pre-validation in both `create()` and `update()` functions:
- Checks category levels before attempting database association
- Returns clear error messages for invalid categories
- Prevents transaction from starting if validation fails

### 3. New API Endpoint for Product Forms

**File**: `web/admin/controllers/categories/categories.controller.go`

Added new endpoint: `/admin/categories/list-for-products`
- Returns only level 4 (parts) categories
- Includes full path information
- Enhanced search across name, label, and path
- Optimized response structure for product forms

**Route**: `web/admin/controllers/categories/categories.go`
```go
categoriesGroup.GET("/list-for-products", controller.Set[dtos.CategoryListDto](listForProducts))
```

### 4. Enhanced Frontend Category Selector

**File**: `web/admin/view/pages/products.templ`

Updated JavaScript category selector:
- Uses new `/list-for-products` endpoint
- Enhanced search functionality (name, label, path, fullPath)
- Displays full category paths in both dropdown and selected categories
- Better user experience with hierarchical information

## Category Hierarchy Levels

```
Level 0: Vehicle (e.g., "cars")
Level 1: Maker/Brand (e.g., "Abarth", "BMW")
Level 2: Model (e.g., "124 Spider", "X5")
Level 3: Type (e.g., "1.4 Turbo Multi Air (2016-2017)")
Level 4: Parts (e.g., "Air Filter", "Engine Oil") ← ONLY THIS LEVEL FOR PRODUCTS
```

## Error Handling

### Database Level Errors
```
"products can only be associated with level 4 categories (parts), category 'BMW' is level 1"
```

### Controller Level Errors
```
"Products can only be associated with level 4 categories (parts). Category 'BMW' is level 1"
```

## API Response Format

### New `/list-for-products` Endpoint Response
```json
{
  "categories": [
    {
      "id": 123,
      "name": "Air Filter",
      "label": "Air Filter",
      "path": "/cars/bmw/x5/x5-2.0d-xdrive",
      "fullPath": "/cars/bmw/x5/x5-2.0d-xdrive/Air Filter",
      "description": "Air Filter for X5 2.0d xDrive",
      "type": "parts"
    }
  ],
  "total": 1000
}
```

## Search Capabilities

The enhanced search now works across:
- Category name
- Category label  
- Category path
- Full category path (path + name)

Users can search for:
- "air filter" → finds all air filter categories
- "bmw" → finds all BMW-related parts
- "x5" → finds all X5 model parts
- "/cars/bmw" → finds parts in BMW hierarchy

## Benefits

1. **Data Integrity**: Ensures products are only linked to actual parts
2. **Better UX**: Users see full category hierarchy when selecting
3. **Enhanced Search**: Find categories by any part of their path
4. **Clear Errors**: Detailed error messages for invalid associations
5. **Performance**: Optimized queries for product forms (level 4 only)

## Backward Compatibility

- Existing valid product-category associations (level 4) remain unchanged
- Invalid associations (if any exist) will be blocked from updates
- Frontend gracefully handles categories without path information 