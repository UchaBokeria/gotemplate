package enums

type FilterType string

const (
	Range       FilterType = "range"       // Range slider
	Checkbox    FilterType = "checkbox"    // Checkbox input
	Select      FilterType = "select"      // Dropdown/select menu
	Radio       FilterType = "radio"       // Radio buttons
	Text        FilterType = "text"        // Text input
	Date        FilterType = "date"        // Date picker
	Number      FilterType = "number"      // Numeric input
	Color       FilterType = "color"       // Color picker
	Switch      FilterType = "switch"      // Toggle switch (checkbox-like)
	MultiSelect FilterType = "multiselect" // Multi-select dropdown
	Rating      FilterType = "rating"      // Star or other rating input
	Search      FilterType = "search"      // Search input
	// File          FilterType = "file"           // File upload
	// Password      FilterType = "password"       // Password input
	// Email         FilterType = "email"          // Email input
	// URL           FilterType = "url"            // URL input
	// Tel           FilterType = "tel"            // Telephone number input
	// Hidden        FilterType = "hidden"         // Hidden input
	// TextArea      FilterType = "textarea"       // Multi-line text input
	// Slider        FilterType = "slider"         // Custom slider input
	// Time          FilterType = "time"           // Time picker
	// Month         FilterType = "month"          // Month picker
	// Week          FilterType = "week"           // Week picker
	// DateTimeLocal FilterType = "datetime-local" // Local date and time picker
	// RangeDate     FilterType = "range-date"     // Date range picker
	// RangeNumber   FilterType = "range-number"   // Number range input
	// TagInput     FilterType = "tag-input"    // Input for tags or labels
	// Image        FilterType = "image"        // Image file upload
	// RangeTime    FilterType = "range-time"   // Time range picker
	// Wysiwyg      FilterType = "wysiwyg"      // Rich text editor
	// Markdown     FilterType = "markdown"     // Markdown editor
	// Autocomplete FilterType = "autocomplete" // Text input with suggestions
	// Currency     FilterType = "currency"     // Currency input
)
