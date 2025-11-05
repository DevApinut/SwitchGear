package model

import (
	"time"

	"gorm.io/gorm"
)

// Upload represents a file upload record
type Upload struct {
	gorm.Model
	// Basic file information
	FileName     string `json:"file_name" gorm:"column:file_name;not null"`
	OriginalName string `json:"original_name" gorm:"column:original_name;not null"`
	FilePath     string `json:"file_path" gorm:"column:file_path;not null"`
	FileSize     int64  `json:"file_size" gorm:"column:file_size;not null"`

	// File metadata
	MimeType  string `json:"mime_type" gorm:"column:mime_type"`
	Extension string `json:"extension" gorm:"column:extension"`
	FileHash  string `json:"file_hash" gorm:"column:file_hash;unique"` // For duplicate detection

	// Upload metadata
	UploadedBy uint      `json:"uploaded_by" gorm:"column:uploaded_by"`
	UploadDate time.Time `json:"upload_date" gorm:"column:upload_date;default:CURRENT_TIMESTAMP"`
	Status     string    `json:"status" gorm:"column:status;default:'uploaded'"` // uploaded, processing, completed, failed

	// Additional metadata
	Description string `json:"description" gorm:"column:description"`
	Tags        string `json:"tags" gorm:"column:tags"`   // JSON string for tags
	Metadata    string `json:"metadata" gorm:"type:json"` // Additional metadata as JSON

	// File processing
	IsProcessed   bool   `json:"is_processed" gorm:"column:is_processed;default:false"`
	ProcessedPath string `json:"processed_path" gorm:"column:processed_path"`
	ThumbnailPath string `json:"thumbnail_path" gorm:"column:thumbnail_path"`
}

// FileUploadRequest represents the request structure for file upload
type FileUploadRequest struct {
	Description string   `json:"description" form:"description"`
	Tags        []string `json:"tags" form:"tags"`
	Category    string   `json:"category" form:"category"`
}

// FileUploadResponse represents the response structure for file upload
type FileUploadResponse struct {
	ID           uint      `json:"id"`
	FileName     string    `json:"file_name"`
	OriginalName string    `json:"original_name"`
	FileSize     int64     `json:"file_size"`
	MimeType     string    `json:"mime_type"`
	Extension    string    `json:"extension"`
	UploadDate   time.Time `json:"upload_date"`
	Status       string    `json:"status"`
	URL          string    `json:"url"` // Public access URL
	ThumbnailURL string    `json:"thumbnail_url,omitempty"`
}

// FileListResponse represents the response for listing files
type FileListResponse struct {
	Files      []FileUploadResponse `json:"files"`
	Total      int64                `json:"total"`
	Page       int                  `json:"page"`
	PageSize   int                  `json:"page_size"`
	TotalPages int                  `json:"total_pages"`
}

// TableName returns the table name for Upload model
func (Upload) TableName() string {
	return "uploads"
}
