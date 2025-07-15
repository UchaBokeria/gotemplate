# ShopHub - Modern E-commerce Platform

A comprehensive, modern e-commerce web application built with the **Gotth stack** (Golang, Tailwind CSS, Templ, htmx) featuring a beautiful user interface inspired by modern shopping platforms like Temu.

## Features

### 🛍️ E-commerce Functionality
- **Modern Landing Page**: Hero section with animated elements and product carousels
- **Product Catalog**: Advanced filtering, search, and multiple view modes
- **Product Grid**: Responsive grid with card, list, and compact view options
- **Shopping Categories**: Electronics, Fashion, Home & Garden, Sports & Outdoors
- **User Authentication**: Login/Register with tabbed interface and social login options

### 🎨 Design & User Experience  
- **Responsive Design**: Mobile-first approach with seamless desktop experience
- **Modern UI Components**: Using custom components with hover effects and animations
- **Navigation**: Multi-language support (EN, RU, GE) with dropdown switcher
- **Interactive Elements**: Smooth animations, hover effects, and transitions
- **Professional Styling**: Beautiful gradients, shadows, and visual hierarchy

### 🌐 Multi-language Support
- English (EN)
- Russian (RU) 
- Georgian (GE)

### 📱 Pages Included
- **Landing Page**: Hero section, product carousels, feature highlights
- **Products Page**: Advanced filtering, search, category filters, grid view modes
- **About Us**: Company story, mission & values, team section, statistics
- **Authentication**: Login/Register forms with social login options
- **Responsive Navigation**: Mobile menu, language switcher, user actions

## Technologies and Libraries Used

- **Golang**: Backend development with robust server-side logic
- **Tailwind CSS**: Utility-first CSS framework for rapid UI development
- **Templ**: Type-safe Go templating engine for dynamic HTML generation
- **htmx**: Modern web applications with server-side logic
- **GORM**: Object-Relational Mapping library for database interactions
- **Echo**: Lightweight and fast Go web framework
- **Air**: Live-reloading tool for enhanced development productivity
- **PostgreSQL**: Powerful, open-source relational database system

## Project Structure

```
├── web/
│   ├── app/
│   │   ├── controllers/     # HTTP route handlers
│   │   ├── view/
│   │   │   ├── components/  # Reusable UI components
│   │   │   │   ├── navigation.templ
│   │   │   │   └── footer.templ
│   │   │   ├── pages/       # Page templates
│   │   │   │   ├── page.templ (Landing)
│   │   │   │   ├── products.templ
│   │   │   │   ├── about.templ
│   │   │   │   └── auth.templ
│   │   │   ├── layouts/     # Layout templates
│   │   │   └── index.templ  # Main layout
│   │   └── app.go          # Route definitions
│   └── web.go              # Web server setup
├── internal/               # Internal application code
├── cmd/                   # Application entry points
├── public/               # Static assets
├── Makefile             # Build and development commands
└── README.md
```

## Getting Started

### Prerequisites
- Go 1.24.4+
- Node.js (for Tailwind CSS)
- PostgreSQL (optional, for database features)

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd shophub
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   npm install  # or bun install
   ```

3. **Set up environment**
   ```bash
   make env  # Creates .env file from .example.env
   ```

4. **Start development server**
   ```bash
   make dev
   ```

This will start:
- Templ file watching and generation
- Tailwind CSS compilation and watching  
- Air for live reloading
- Web server on port 3000

### Development Commands

```bash
# Start development environment with live reloading
make dev

# Build for production
make prod

# Generate Templ files
make templ

# Compile Tailwind CSS
make tailwind

# Run tests
make test

# Database migrations (when database is set up)
make migrate

# Seed database with sample data
make seed
```

## Features Walkthrough

### Landing Page
- **Hero Section**: Gradient background with animated elements and call-to-action buttons
- **Product Carousels**: Categorized product sections with "View All" buttons
- **Feature Highlights**: Free shipping, easy returns, secure payment
- **Responsive Design**: Optimized for all screen sizes

### Products Page
- **Advanced Filtering**: Search, categories, price ranges, brands, ratings
- **View Modes**: Grid (default), list, and compact views
- **Sorting Options**: Featured, price, rating, newest
- **Pagination**: Navigate through product results
- **Product Cards**: Detailed cards with ratings, prices, and add-to-cart buttons

### About Us Page
- **Company Story**: History and mission narrative
- **Values Section**: Core company values with icons
- **Team Section**: Leadership team profiles
- **Statistics**: Key company metrics and achievements

### Authentication
- **Tabbed Interface**: Switch between login and register
- **Form Validation**: Client-side and server-side validation
- **Social Login**: Google and Facebook integration options
- **Security Features**: Password confirmation, terms acceptance

### Navigation & UX
- **Responsive Navigation**: Mobile-friendly hamburger menu
- **Language Switcher**: Dropdown with flag icons
- **Smooth Animations**: Fade-in effects and hover transitions
- **Loading States**: Visual feedback for user actions

## Customization

The application is built with modularity in mind:

1. **Components**: Reusable UI components in `web/app/view/components/`
2. **Styling**: Tailwind CSS classes with custom animations in `imports.templ`
3. **Layout**: Main layout structure in `web/app/view/index.templ`
4. **Routing**: Route definitions in `web/app/app.go`

## Browser Support

- Chrome 90+
- Firefox 88+  
- Safari 14+
- Edge 90+
- Mobile browsers with modern CSS support

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes and add tests
4. Ensure all tests pass (`make test`)
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [a-h/templ](https://github.com/a-h/templ) - Type-safe Go templating
- [Tailwind CSS](https://tailwindcss.com/) - Utility-first CSS framework
- [Echo](https://echo.labstack.com/) - High performance Go web framework
- [HTMX](https://htmx.org/) - High power tools for HTML
- Design inspiration from modern e-commerce platforms

---

**Built with ❤️ using modern Go web technologies** 