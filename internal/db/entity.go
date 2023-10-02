package db

import (
	"context"
	"errors"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pessoa struct {
	Id         uuid.UUID `json:"id"`
	Apelido    string    `json:"Apelido"`
	Nome       string    `json:"Nome"`
	Nascimento string    `json:"Nascimento"`
	Stack      []string  `json:"Stack"`
}

func (p *Pessoa) CreatePerson(db *pgxpool.Pool) error {
	query := "INSERT INTO pessoas(Apelido, Nome, Nascimento, Stack) VALUES($1, $2, $3, $4) RETURNING id"

	err := db.QueryRow(context.Background(), query, p.Apelido, p.Nome, p.Nascimento, strings.Join(p.Stack, " | ")).Scan(&p.Id)
	if err != nil {
		err = errors.New("apelido já cadastrado")
	}
	return err
}

func (p *Pessoa) getPerson(db *pgxpool.Pool, id string) error {
	return db.QueryRow(context.Background(), "SELECT ID, Apelido, Nome, Nascimento, string_to_array(Stack, ' | ') as Stack FROM pessoas WHERE id=$1", id).Scan(&p.Id, &p.Apelido, &p.Nome, &p.Nascimento, &p.Stack)
}

func (p *Pessoa) searchPerson(db *pgxpool.Pool, term string) ([]Pessoa, error) {
	pessoas := make([]Pessoa, 50)
	rows, err := db.Query(context.Background(), "SELECT ID, Apelido, Nome, Nascimento, string_to_array(Stack, ' | ') as stack FROM pessoas WHERE busca ilike '%' || $1 || '%' limit 50", term)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var currentValue int
	for rows.Next() {
		var p Pessoa
		err = rows.Scan(&p.Id, &p.Apelido, &p.Nome, &p.Nascimento, &p.Stack)
		if err != nil {
			return nil, err
		}
		pessoas[currentValue] = p
		currentValue++
	}
	return pessoas, nil

}

func (p *Pessoa) countPerson(db *pgxpool.Pool) (int, error) {
	var total int
	err := db.QueryRow(context.Background(), "SELECT COUNT(id) FROM pessoas").Scan(&total)
	return total, err
}

func (p *Pessoa) Validate() bool {
	var val bool
	val = p.validateNome()
	if !val {
		return !val
	}
	val = p.validateApelido()
	if !val {
		return !val
	}
	val = p.validateNascimento()
	if !val {
		return !val
	}
	val = p.validateStack()
	if !val {
		return !val
	}
	return true
}

func (p *Pessoa) validateApelido() bool {
	return len(p.Apelido) > 0 && len(p.Apelido) <= 32
}

func (p *Pessoa) validateNome() bool {
	return len(p.Nome) > 0 && len(p.Nome) <= 100
}

func (p *Pessoa) validateNascimento() bool {
	rgx := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	return rgx.MatchString(p.Nascimento)
}

func (p *Pessoa) validateStack() bool {
	if len(p.Stack) == 0 {
		return true
	}

	for i := 0; i < len(p.Stack); i += 1 {
		if len(p.Stack[i]) > 32 {
			return false
		}
	}

	return true
}
