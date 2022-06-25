package controller

import (
	"Persona-Mysql_Rest/model"
	"context"
	"database/sql"
	"log"
	// "strconv"
	// "github.com/gorilla/mux"
)

func GetPersonas1(ctx context.Context, db *sql.DB) (model.AllPersonas, error) {
	query := "Select * From Persona"
	ListaPersonas1 := model.AllPersonas{}

	datos, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for datos.Next() {
		per := model.Persona{}
		err = datos.Scan(&per.IdPersona, &per.Nombre, &per.Apellido, &per.Email, &per.Genero)
		if err != nil {
			return nil, err
		}
		ListaPersonas1 = append(ListaPersonas1, per)
	}

	// log.Println("%+v",ListaPersonas1)
	return ListaPersonas1, err

}

func CreatePerson(ctx context.Context, db *sql.DB, p *model.Persona) error {
	query := `insert into Persona values (?,?,?,?,?)`

	datos, err := db.ExecContext(ctx, query, p.IdPersona, p.Nombre, p.Apellido, p.Email, p.Genero)

	if err != nil {
		return err
	}
	id, err := datos.LastInsertId()
	defer db.Close()

	if err != nil {
		return err
	}
	log.Println(id)
	return nil
}

func EliminarPersona1(ctx context.Context, db *sql.DB, id int64) (bool, error) {
	query := "Delete From Persona Where IdPersona=?"

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func BuscarPersona1(ctx context.Context, db *sql.DB, id int64) model.Persona {
	var per model.Persona
	query := "Select * From Persona where IdPersona=?"

	//err := db.QueryRow("SELECT * FROM users WHERE name = $1", name).Scan(&u.Id, &u.Name, &u.Score)

	db.QueryRowContext(ctx, query, id).Scan(&per.IdPersona, &per.Nombre, &per.Apellido, &per.Email, &per.Genero)

	return per
}

func UpdatePersona1(ctx context.Context, db *sql.DB, p *model.Persona) error {
	//	var p model.Persona

	query := "Update Persona SET Nombre=?, Apellido=?, Email=?, Genero=? where IdPersona=?"

	data, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	_, err = data.ExecContext(
		ctx,
		p.Nombre,
		p.Apellido,
		p.Email,
		p.Genero,
		p.IdPersona,
	)

	if err != nil {
		return err
	}
	defer data.Close()

	return nil

}
