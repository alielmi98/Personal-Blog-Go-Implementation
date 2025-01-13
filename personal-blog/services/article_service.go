package services

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/alielmi98/Personal-Blog-Go-Implementation/dto"
)

// ArticleService handles operations related to articles.
type ArticleService struct {
	Mu       sync.RWMutex     // Use RWMutex for concurrent access
	articles []dto.ArticleDTO // Slice to hold articles
	dir      string           // Directory for storing articles
	NextID   int              // Next available ID for new articles
}

// NewArticleService creates a new ArticleService.
func NewArticleService(dir string) *ArticleService {
	return &ArticleService{
		articles: []dto.ArticleDTO{}, // Initialize with an empty slice
		dir:      dir,
		NextID:   1,
	}
}

// LoadArticles loads articles from the JSON file into memory.
func (as *ArticleService) LoadArticles() error {
	as.Mu.Lock() // Lock for writing
	defer as.Mu.Unlock()

	// Check if the file exists and Create a new file if it does not exist
	if _, err := os.Stat(as.dir); os.IsNotExist(err) {
		file, err := os.Create(as.dir)
		if err != nil {
			return fmt.Errorf("error creating file: %w", err)
		}
		defer file.Close()

		return nil
	}

	// Read the file data
	data, err := os.ReadFile(as.dir)
	if err != nil {
		return fmt.Errorf("error loading file: %w", err)
	}

	// Check if the file is empty
	if len(data) == 0 {
		as.articles = []dto.ArticleDTO{} // Set articles to an empty slice
		as.NextID = 1                    // Initialize NextID to 1
		return nil
	}

	// Unmarshal the JSON data into the articles slice
	if err := json.Unmarshal(data, &as.articles); err != nil {
		return fmt.Errorf("error unmarshaling data: %w", err)
	}

	// Update NextID based on the highest article ID
	for _, article := range as.articles {
		if article.ID >= as.NextID {
			as.NextID = article.ID + 1
		}
	}

	return nil
}

// GetAllArticles retrieves all articles.
func (as *ArticleService) GetAllArticles() ([]dto.ArticleDTO, error) {
	// Load articles if the slice is empty
	if len(as.articles) == 0 {
		if err := as.LoadArticles(); err != nil {
			return nil, err
		}
	}

	return as.articles, nil
}

// SaveArticles saves all articles back to the JSON file.
func (as *ArticleService) SaveArticles() error {
	data, err := json.Marshal(as.articles)
	if err != nil {
		return fmt.Errorf("error marshalling articles: %w", err)
	}

	if err := os.WriteFile(as.dir, data, 0644); err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}

// CreateArticle saves a new article to the service.
func (as *ArticleService) CreateArticle(article dto.ArticleDTO) error {
	as.Mu.Lock()
	defer as.Mu.Unlock()

	article.CreatedAt = time.Now()
	article.ID = as.NextID
	as.articles = append(as.articles, article)
	as.NextID++

	return as.SaveArticles()
}

// DeleteArticle deletes an article from the slice and saves the remaining articles.
func (as *ArticleService) DeleteArticle(id int) error {
	as.Mu.Lock()
	defer as.Mu.Unlock()

	for i, article := range as.articles {
		if article.ID == id {
			as.articles = append(as.articles[:i], as.articles[i+1:]...)
			return as.SaveArticles() // Save remaining articles
		}
	}
	return fmt.Errorf("article with ID %d not found", id)
}

// GetArticle retrieves an article by its ID.
func (as *ArticleService) GetArticle(id int) (dto.ArticleDTO, bool) {
	as.Mu.RLock() // Lock for reading
	defer as.Mu.RUnlock()

	for _, article := range as.articles {
		if article.ID == id {
			return article, true
		}
	}
	return dto.ArticleDTO{}, false
}

// UpdateArticle updates an existing article by ID.
func (as *ArticleService) UpdateArticle(updatedArticle dto.ArticleDTO) error {
	as.Mu.Lock()
	defer as.Mu.Unlock()

	for i, article := range as.articles {
		if article.ID == updatedArticle.ID {
			updatedArticle.CreatedAt = article.CreatedAt
			updatedArticle.ModifiedAt = time.Now()
			as.articles[i] = updatedArticle
			return as.SaveArticles()
		}
	}
	return fmt.Errorf("article with ID %d not found", updatedArticle.ID)
}
