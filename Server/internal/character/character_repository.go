package character

import (
	"context"
	"database/sql"
	"dungeons_helper/db"
)

type repository struct {
	db db.DatabaseTX
}

func NewRepository(db db.DatabaseTX) Repository {
	return &repository{db: db}
}

func (r *repository) GetAllCharactersByAccId(ctx context.Context, idAcc int64) ([]GetAllCharactesRes, error) {
	var chars []GetAllCharactesRes

	query := `SELECT c.id, c.charName, i.image FROM characters c
		JOIN accChar ac ON c.id = ac.idChar JOIN class cl ON c.idClass = cl.id
		JOIN image i ON c.idAvatar = i.id WHERE ac.idAccount = ? GROUP BY c.id`

	rows, err := r.db.QueryContext(ctx, query, idAcc)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var char GetAllCharactesRes
		err := rows.Scan(&char.IdChar, &char.CharName, &char.Avatar)
		if err != nil {
			return nil, err
		}
		chars = append(chars, char)
	}

	return chars, nil
}

func (r *repository) GetCharacterById(ctx context.Context, id int64) (*Character, error) {
	char := Character{}
	//var imageBytes []byte

	query := `SELECT c.id, c.hp, c.lvl, c.exp, c.charName, c.sex, c.weight, c.height, c.addLanguage,
		c.ideals, c.weaknesses, c.traits, c.allies, c.organizations, c.enemies, c.story,
		c.goals, c.treasures, c.notes, cl.className, r.raceName, s.subraceName, st.strength,
		st.dexterity, st.constitution, st.intelligence, st.wisdom, st.charisma,
		GROUP_CONCAT(sk.skillName SEPARATOR ', ')
		AS characterSkills, a.alignmentName, i.image FROM characters c
		JOIN accChar ac ON c.id = ac.idChar JOIN class cl ON c.idClass = cl.id
		JOIN races r on c.idRace = r.id JOIN subrace s on c.idSubrace = s.id
		JOIN stats st on c.idStats = st.id JOIN charSkills cs on c.id = cs.idChar
		JOIN skills sk on cs.idSkill = sk.id JOIN alignment a on c.idAlignment = a.id
		JOIN image i ON c.idAvatar = i.id WHERE c.id = ? GROUP BY c.id`

	err := r.db.QueryRowContext(ctx, query, id).Scan(&char.Id, &char.Hp, &char.Lvl, &char.Exp, &char.CharName, &char.Sex, &char.Weight, &char.Height, &char.AddLanguage,
		&char.Ideals, &char.Weaknesses, &char.Traits, &char.Allies, &char.Organizations, &char.Enemies,
		&char.Story, &char.Goals, &char.Treasures, &char.Notes, &char.Class, &char.Race,
		&char.Subrace, &char.Stats.Strength, &char.Stats.Dexterity, &char.Stats.Constitution,
		&char.Stats.Intelligence, &char.Stats.Wisdom, &char.Stats.Charisma, &char.CharacterSkills,
		&char.Alignment, &char.Avatar)
	if err != nil {
		return nil, err
	}

	return &char, nil
}

func (r *repository) CreateCharacter(ctx context.Context, char *CreateCharacterReq) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
			return
		}
		err = tx.Commit()
	}()

	query := `INSERT INTO image(image) VALUE (?)`
	result, err := tx.ExecContext(ctx, query, char.Avatar)
	if err != nil {
		return err
	}
	idImage, err := result.LastInsertId()
	if err != nil {
		return err
	}

	query = `INSERT INTO stats(strength, dexterity, constitution, intelligence, wisdom, charisma) VALUES (?, ?, ?, ?, ?, ?)`
	result, err = tx.ExecContext(ctx, query, char.Stats.Strength, char.Stats.Dexterity, char.Stats.Constitution, char.Stats.Intelligence, char.Stats.Wisdom, char.Stats.Charisma)
	if err != nil {
		return err
	}
	statsId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	query = `INSERT INTO characters(hp, lvl, exp, idAvatar, charName, sex, weight, height, idClass, 
                       idRace, idSubrace, idStats, addLanguage, idAlignment, ideals, weaknesses, 
                       traits, allies, organizations, enemies, story, goals, treasures, notes)  
					   VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err = tx.ExecContext(ctx, query, char.Hp, 1, char.Exp, idImage, char.CharName, char.Sex, char.Weight, char.Height,
		char.Class.Id, char.Race.Id, char.Subrace.Id, statsId, char.AddLanguage, char.Alignment.Id, char.Ideals,
		char.Weaknesses, char.Traits, char.Allies, char.Organizations, char.Enemies, char.Story, char.Goals, char.Treasures, char.Notes)
	if err != nil {
		return err
	}
	charId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	query = `INSERT INTO charSkills(idSkill, idChar) VALUES (?, ?)`
	for _, skill := range char.CharacterSkills {
		_, err = tx.ExecContext(ctx, query, skill.Id, charId)
		if err != nil {
			return err
		}
	}

	query = `UPDATE accChar SET act = false WHERE idAccount = ?`
	_, err = tx.ExecContext(ctx, query, char.IdAcc)
	if err != nil {
		return err
	}

	query = `INSERT INTO accChar(act, idAccount, idChar) VALUES (?, ?, ?)`
	_, err = tx.ExecContext(ctx, query, true, char.IdAcc, charId)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateCharacterHpById(ctx context.Context, id int64, hp int64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
			return
		}
		err = tx.Commit()
	}()

	query := `UPDATE characters SET hp = ? WHERE id = ?`
	_, err = tx.ExecContext(ctx, query, hp, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateCharacterExpById(ctx context.Context, id int64, hp int64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
			return
		}
		err = tx.Commit()
	}()

	query := `UPDATE characters SET exp = ? WHERE id = ?`
	_, err = tx.ExecContext(ctx, query, hp, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) SetActiveCharacterById(ctx context.Context, req *SetActiveCharReq) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
			return
		}
		err = tx.Commit()
	}()

	query := "UPDATE accChar SET act = false WHERE idAccount = ?"
	_, err = tx.ExecContext(ctx, query, req.IdAcc)
	if err != nil {
		return err
	}

	query = "UPDATE accChar SET act = true WHERE idChar = ?"
	_, err = tx.ExecContext(ctx, query, req.IdChar)
	if err != nil {
		return err
	}

	return nil
}
