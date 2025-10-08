package spaces

import (
	"fmt"
	"time"

	"github.com/fancyinnovations/fancyspaces/src/internal/auth"
	"github.com/google/uuid"
)

type DB interface {
	GetByID(id string) (*Space, error)
	GetBySlug(slug string) (*Space, error)
	Create(s *Space) (*Space, error)
	Update(id string, s *Space) error
	Delete(id string) error
}

type Store struct {
	db DB
}

type Configuration struct {
	DB DB
}

func New(cfg Configuration) *Store {
	return &Store{
		db: cfg.DB,
	}
}

func (s *Store) GetByID(id string) (*Space, error) {
	return s.db.GetByID(id)
}

func (s *Store) GetBySlug(slug string) (*Space, error) {
	return s.db.GetBySlug(slug)
}

func (s *Store) Create(creator *auth.User, req *CreateOrUpdateSpaceReq) (*Space, error) {
	if !creator.IsActive {
		return nil, ErrUserNotActive
	}
	if !creator.Verified {
		return nil, ErrUserNotVerified
	}

	space := &Space{
		ID:          uuid.New().String(),
		Slug:        req.Slug,
		Title:       req.Title,
		Description: req.Description,
		Categories:  req.Categories,
		IconURL:     req.IconURL,
		Status:      StatusDraft,
		CreatedAt:   time.Now(),
		Members: []Member{
			{
				UserID: creator.ID,
				Role:   RoleOwner,
			},
		},
	}

	if err := space.Validate(); err != nil {
		return nil, fmt.Errorf("invalid space: %w", err)
	}

	return s.db.Create(space)
}

func (s *Store) Update(id string, req *CreateOrUpdateSpaceReq) error {
	space, err := s.db.GetByID(id)
	if err != nil {
		return err
	}

	space.Slug = req.Slug
	space.Title = req.Title
	space.Description = req.Description
	space.Categories = req.Categories
	space.IconURL = req.IconURL

	if err := space.Validate(); err != nil {
		return fmt.Errorf("invalid space: %w", err)
	}

	return s.db.Update(id, space)
}

func (s *Store) Delete(id string) error {
	return s.db.Delete(id)
}

func (s *Store) ChangeStatus(space *Space, to Status) error {
	if to == space.Status {
		return nil // no change
	}

	space.Status = to

	if err := space.Validate(); err != nil {
		return fmt.Errorf("invalid space: %w", err)
	}

	return s.db.Update(space.ID, space)
}
