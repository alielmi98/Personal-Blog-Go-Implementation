# Personal Blog

A personal blog built with Go, allowing users to write and publish articles on various topics. This project is part of the backend projects for beginners on [roadmap.sh](https://roadmap.sh/projects/personal-blog) and aims to help Go developers practice their skills.

## Features

- **Guest Section**:
- Home Page: List of published articles.
- Article Page: Display content of individual articles.

- **Admin Section**:
- Dashboard: Manage articles (create, edit, delete).
- Authentication: Secure admin area with simple login.

## Technologies Used

- **Backend**: Go (Golang)
- **Storage**: Filesystem (articles stored as JSON or Markdown)
- **Frontend**: HTML, CSS
- **Design Patterns**: MVC (Model-View-Controller), Repository Pattern

## Installation

### Prerequisites

- Go installed on your machine.

### Steps

1. Clone the repository:
```bash
git clone https://github.com/alielmi98/Personal-Blog-Go-Implementation.git
cd personal-blog
```

2. Run the application:
```bash
go run main.go
```

3. Access the application at `http://localhost:PORT`, where `PORT` is defined in your configuration.

## Usage

- **Guest Users**: Access the home page to view articles and click on an article title to read its content.
- **Admin Users**: Log in to manage articles via the dashboard.

## Project Structure

```
personal-blog/
├── main.go                 # Main entry point of the application
├── go.mod                  # Go module file
├── articles/               # Directory for article storage
│   └── article.json        # JSON file for articles
├── config/                 # Configuration files
│   ├── config.go
│   └── config.json
├── dto/                    # Data Transfer Objects
│   └── data_transfer_objects.go
├── handlers/               # HTTP Handlers
│   ├── admin.go
│   ├── guest.go
│   ├── login_logout.go
│   └── render.go
├── middlewares/            # Middleware for authentication
│   └── auth.go
├── services/               # Business logic for article management
│   └── article_service.go
├── static/                 # Static files (CSS)
│   └── style.css
└── templates/              # HTML templates
    ├── base.tmpl
    ├── dashboard.tmpl
    ├── home.tmpl
    ├── article.tmpl
    ├── login.tmpl
    ├── create_article.tmpl
    └── update_article.tmpl
```

## Acknowledgments

This project is based on guidelines from [roadmap.sh](https://roadmap.sh/projects/personal-blog) under the backend projects for beginners section, aimed at helping and educating other developers.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.