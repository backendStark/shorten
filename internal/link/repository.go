package link

import "shorten/pkg/db"

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(db *db.Db) *LinkRepository {
	return &LinkRepository{db}
}	

func (repo *LinkRepository) Create(link *Link) error {
	return repo.Database.Create(link).Error
}