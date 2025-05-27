package migrate

import (
	repositoryJira "team-meter/internal/repository/jira"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
	return db.AutoMigrate(
		&repositoryJira.Issue{},
		&repositoryJira.IssueChangelog{},
	)
}
