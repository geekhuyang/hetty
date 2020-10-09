package proj

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dstotijn/hetty/pkg/db/sqlite"
	"github.com/dstotijn/hetty/pkg/scope"
)

// Service is used for managing projects.
type Service struct {
	dbPath string
	db     *sqlite.Client
	name   string

	Scope *scope.Scope
}

type Project struct {
	Name         string
	DatabasePath string
}

var ErrNoProject = errors.New("proj: no open project")

// NewService returns a new Service.
func NewService(dbPath string) (*Service, error) {
	// Create directory for DBs if it doesn't exist yet.
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dbPath, 0755); err != nil {
			return nil, fmt.Errorf("proj: could not create project directory: %v", err)
		}
	}

	return &Service{
		dbPath: dbPath,
		db:     &sqlite.Client{},
		Scope:  scope.New(nil),
	}, nil
}

// Close closes the currently open project database (if there is one).
func (svc *Service) Close() error {
	return svc.db.Close()
}

// Database returns the currently open database. If no database is open, it will
// return `nil`.
func (svc *Service) Database() *sqlite.Client {
	return svc.db
}

// Open opens a database identified with `name`. If a database with this
// identifier doesn't exist yet, it will be automatically created.
func (svc *Service) Open(name string) (Project, error) {
	if err := svc.db.Close(); err != nil {
		return Project{}, fmt.Errorf("proj: could not close previously open database: %v", err)
	}

	dbPath := filepath.Join(svc.dbPath, name+".db")

	err := svc.db.Open(dbPath)
	if err != nil {
		return Project{}, fmt.Errorf("proj: could not open database: %v", err)
	}

	svc.name = name

	return Project{
		Name:         name,
		DatabasePath: dbPath,
	}, nil
}

func (svc *Service) CurrentProject() (Project, error) {
	if !svc.db.IsOpen() {
		return Project{}, ErrNoProject
	}
	dbPath := filepath.Join(svc.dbPath, svc.name+".db")

	return Project{
		DatabasePath: dbPath,
		Name:         svc.name,
	}, nil
}
