package creature

import "database/sql"

// Repository defines the interface for creature data operations
type Repository interface {
	GetAll() ([]Creature, error)
	GetByID(id int64) (*Creature, error)
	Create(creature *Creature) error
	Update(id int64, creature *Creature) error
	Delete(id int64) error
}

// SQLiteRepository represents the SQLite implementation of the creature repository
type SQLiteRepository struct {
	db *sql.DB
}

// NewSQLiteRepository creates a new SQLiteRepository instance
func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

//
// Implement the Repository interface methods for SQLiteRepository
//

// GetAll retrieves all creatures from the database
func (r *SQLiteRepository) GetAll() ([]Creature, error) {
	query := "SELECT * FROM creatures"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	creatures := make([]Creature, 0)
	for rows.Next() {
		var creature Creature
		err := rows.Scan(
			&creature.ID,
			&creature.Picture,
			&creature.Name,
			&creature.Type,
			&creature.HP,
		)
		if err != nil {
			return nil, err
		}
		creatures = append(creatures, creature)
	}

	return creatures, nil
}

// GetByID retrieves a creature by its ID from the database
func (r *SQLiteRepository) GetByID(id int64) (*Creature, error) {
	query := "SELECT * FROM creatures WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var creature Creature
	err := row.Scan(
		&creature.ID,
		&creature.Picture,
		&creature.Name,
		&creature.Type,
		&creature.HP,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No rows found
		}
		return nil, err
	}

	return &creature, nil
}

// Create inserts a new creature into the database
func (r *SQLiteRepository) Create(creature *Creature) error {
	query := `
		INSERT INTO creatures (
			picture,
			name,
			type,
			hp,
		) VALUES (?, ?, ?, ?)
	`

	result, err := r.db.Exec(query,
		creature.Picture,
		creature.Name,
		creature.Type,
		creature.HP,
	)
	if err != nil {
		return err
	}

	creatureID, _ := result.LastInsertId()
	creature.ID = creatureID

	return nil
}

// Update updates an existing creature in the database
func (r *SQLiteRepository) Update(id int64, creature *Creature) error {
	query := `
		UPDATE creatures SET
			picture = ?,
			name = ?,
			type = ?,
			hp = ?,
		WHERE id = ?
	`

	_, err := r.db.Exec(query,
		creature.Picture,
		creature.Name,
		creature.Type,
		creature.HP,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes a creature from the database
func (r *SQLiteRepository) Delete(id int64) error {
	query := "DELETE FROM creatures WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
