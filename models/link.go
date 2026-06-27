// Package models
package models

import (
	"time"
)

/*
* ID
* Title
* Url
* description
* created_at
* updated_at
*
 */

type Link struct {
	ID        int       `json:"-"`
	Title     string    `json:"title" binding:"required"`
	URL       string    `json:"url" binding:"required,url"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type LinkUpdate struct {
	ID        int       `json:"-"`
	Title     string    `json:"title,omitempty"`
	URL       string    `json:"url" binding:"omitempty,url"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
