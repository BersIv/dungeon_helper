package character

import (
	"bytes"
	"context"
	"dungeons_helper/internal/alignment"
	"dungeons_helper/internal/class"
	"dungeons_helper/internal/races"
	"dungeons_helper/internal/skills"
	"dungeons_helper/internal/stats"
	"dungeons_helper/internal/subraces"
	"dungeons_helper/utilMocks"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

type MockRepo struct{}

func (m *MockRepo) GetAllCharactersByAccId(ctx context.Context, idAcc int64) ([]GetAllCharactesRes, error) {
	return []GetAllCharactesRes{
		{IdChar: 1, CharName: "First", Avatar: "avatar_data"},
		{IdChar: 2, CharName: "Second", Avatar: "avatar_data"}}, nil
}

func (m *MockRepo) GetCharacterById(ctx context.Context, id int64) (*Character, error) {
	return &Character{
		Id:            1,
		Hp:            100,
		Lvl:           1,
		Exp:           0,
		CharName:      "charName",
		Sex:           true,
		Weight:        100,
		Height:        100,
		AddLanguage:   "addLanguage",
		Ideals:        "ideals",
		Weaknesses:    "weaknesses",
		Traits:        "traits",
		Allies:        "allies",
		Organizations: "organizations",
		Enemies:       "enemies",
		Story:         "story",
		Goals:         "goals",
		Treasures:     "treasures",
		Notes:         "notes",
		Class:         "className",
		Race:          "raceName",
		Subrace:       "subraceName",
		Stats: stats.GetStatsRes{
			Strength:     1,
			Dexterity:    2,
			Constitution: 3,
			Intelligence: 4,
			Wisdom:       5,
			Charisma:     6,
		},
		CharacterSkills: "characterSkills",
		Alignment:       "alignmentName",
		Avatar:          "image",
	}, nil
}

func (m *MockRepo) CreateCharacter(ctx context.Context, character *CreateCharacterReq) error {
	return nil
}

func (m *MockRepo) UpdateCharacterHpById(ctx context.Context, id int64, hp int64) error {
	return nil
}

func (m *MockRepo) UpdateCharacterExpById(ctx context.Context, id int64, exp int64) error {
	return nil
}

func (m *MockRepo) SetActiveCharacterById(ctx context.Context, req *SetActiveCharReq) error {
	return nil
}

type MockRepoError struct{}

func (m *MockRepoError) GetAllCharactersByAccId(ctx context.Context, idAcc int64) ([]GetAllCharactesRes, error) {
	return []GetAllCharactesRes{}, errors.New("fake repo error")
}

func (m *MockRepoError) GetCharacterById(ctx context.Context, id int64) (*Character, error) {
	return &Character{}, errors.New("fake repo error")
}

func (m *MockRepoError) CreateCharacter(ctx context.Context, character *CreateCharacterReq) error {
	return errors.New("fake repo error")
}

func (m *MockRepoError) UpdateCharacterHpById(ctx context.Context, id int64, hp int64) error {
	return errors.New("fake repo error")
}

func (m *MockRepoError) UpdateCharacterExpById(ctx context.Context, id int64, exp int64) error {
	return errors.New("fake repo error")
}

func (m *MockRepoError) SetActiveCharacterById(ctx context.Context, req *SetActiveCharReq) error {
	return errors.New("fake repo error")
}

type MockService struct{}

func (m *MockService) GetAllCharactersByAccId(c context.Context, idAcc int64) ([]GetAllCharactesRes, error) {
	return []GetAllCharactesRes{
		{IdChar: 1, CharName: "First", Avatar: "avatar_data"},
		{IdChar: 2, CharName: "Second", Avatar: "avatar_data"}}, nil
}

func (m *MockService) GetCharacterById(c context.Context, id int64) (*Character, error) {
	return &Character{
		Id:            1,
		Hp:            100,
		Lvl:           1,
		Exp:           0,
		CharName:      "charName",
		Sex:           true,
		Weight:        100,
		Height:        100,
		AddLanguage:   "addLanguage",
		Ideals:        "ideals",
		Weaknesses:    "weaknesses",
		Traits:        "traits",
		Allies:        "allies",
		Organizations: "organizations",
		Enemies:       "enemies",
		Story:         "story",
		Goals:         "goals",
		Treasures:     "treasures",
		Notes:         "notes",
		Class:         "className",
		Race:          "raceName",
		Subrace:       "subraceName",
		Stats: stats.GetStatsRes{
			Strength:     1,
			Dexterity:    2,
			Constitution: 3,
			Intelligence: 4,
			Wisdom:       5,
			Charisma:     6,
		},
		CharacterSkills: "characterSkills",
		Alignment:       "alignmentName",
		Avatar:          "image",
	}, nil
}

func (m *MockService) CreateCharacter(c context.Context, character *CreateCharacterReq) error {
	return nil
}

func (m *MockService) SetActiveCharacterById(c context.Context, req *SetActiveCharReq) error {
	return nil
}

type MockServiceError struct{}

func (m *MockServiceError) GetAllCharactersByAccId(c context.Context, idAcc int64) ([]GetAllCharactesRes, error) {
	return []GetAllCharactesRes{}, errors.New("fake service error")
}

func (m *MockServiceError) GetCharacterById(c context.Context, id int64) (*Character, error) {
	return &Character{}, errors.New("fake service error")
}

func (m *MockServiceError) CreateCharacter(c context.Context, character *CreateCharacterReq) error {
	return errors.New("fake service error")
}

func (m *MockServiceError) SetActiveCharacterById(c context.Context, req *SetActiveCharReq) error {
	return errors.New("fake service error")
}

func TestAllCharactersByAccIdRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockRows := sqlmock.NewRows([]string{"id", "charName", "image"}).
		AddRow(1, "First", "image_data").
		AddRow(2, "Second", "image_data")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT c.id, c.charName, i.image FROM characters c
						JOIN accChar ac ON c.id = ac.idChar JOIN class cl ON c.idClass = cl.id
						JOIN image i ON c.idAvatar = i.id WHERE ac.idAccount = ? GROUP BY c.id`)).WillReturnRows(mockRows)

	repo := NewRepository(db)

	ctx := context.Background()
	characters, err := repo.GetAllCharactersByAccId(ctx, 1)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if characters == nil {
		t.Errorf("expected account data, got %+v", characters)
	}

	expected := []GetAllCharactesRes{
		{IdChar: 1, CharName: "First", Avatar: "image_data"},
		{IdChar: 2, CharName: "Second", Avatar: "image_data"},
	}

	if !reflect.DeepEqual(characters, expected) {
		t.Errorf("expected %+v, got %+v", expected, characters)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAllCharactersByAccIdRepository_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT c.id, c.charName, i.image FROM characters c
						JOIN accChar ac ON c.id = ac.idChar JOIN class cl ON c.idClass = cl.id
						JOIN image i ON c.idAvatar = i.id WHERE ac.idAccount = ? GROUP BY c.id`)).WillReturnError(errors.New("fake error"))

	repo := NewRepository(db)

	ctx := context.Background()
	_, err = repo.GetAllCharactersByAccId(ctx, 1)
	if err.Error() != "fake error" {
		t.Errorf("unexpected error: %s", err)
	}

}

func TestCharacterByIdRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockRows := sqlmock.NewRows([]string{"id", "hp", "lvl", "exp", "charName", "sex", "weight", "height", "addLanguage", "ideals", "weaknesses", "traits",
		"allies", "organzations", "enemies", "story", "goals", "treasures", "notes", "className", "raceName", "subraceName", "sttrength", "dexterity", "constitution",
		"intelligence", "wisdom", "charisma", "characterSkills", "alignmentName", "image"}).
		AddRow(1, 100, 1, 0, "charName", true, 100, 100, "addLanguage", "ideals", "weaknesses", "traits",
			"allies", "organizations", "enemies", "story", "goals", "treasures", "notes", "className", "raceName", "subraceName", 1, 2, 3,
			4, 5, 6, "characterSkills", "alignmentName", "image")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT c.id, c.hp, c.lvl, c.exp, c.charName, c.sex, c.weight, c.height, c.addLanguage,
						c.ideals, c.weaknesses, c.traits, c.allies, c.organizations, c.enemies, c.story,
						c.goals, c.treasures, c.notes, cl.className, r.raceName, s.subraceName, st.strength,
						st.dexterity, st.constitution, st.intelligence, st.wisdom, st.charisma,
						GROUP_CONCAT(sk.skillName SEPARATOR ', ')
						AS characterSkills, a.alignmentName, i.image FROM characters c
						JOIN accChar ac ON c.id = ac.idChar JOIN class cl ON c.idClass = cl.id
						JOIN races r on c.idRace = r.id JOIN subrace s on c.idSubrace = s.id
						JOIN stats st on c.idStats = st.id JOIN charSkills cs on c.id = cs.idChar
						JOIN skills sk on cs.idSkill = sk.id JOIN alignment a on c.idAlignment = a.id
						JOIN image i ON c.idAvatar = i.id WHERE c.id = ? GROUP BY c.id`)).WillReturnRows(mockRows)

	repo := NewRepository(db)

	ctx := context.Background()
	character, err := repo.GetCharacterById(ctx, 1)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if character == nil {
		t.Errorf("expected account data, got %+v", character)
	}

	expected := &Character{
		Id:            1,
		Hp:            100,
		Lvl:           1,
		Exp:           0,
		CharName:      "charName",
		Sex:           true,
		Weight:        100,
		Height:        100,
		AddLanguage:   "addLanguage",
		Ideals:        "ideals",
		Weaknesses:    "weaknesses",
		Traits:        "traits",
		Allies:        "allies",
		Organizations: "organizations",
		Enemies:       "enemies",
		Story:         "story",
		Goals:         "goals",
		Treasures:     "treasures",
		Notes:         "notes",
		Class:         "className",
		Race:          "raceName",
		Subrace:       "subraceName",
		Stats: stats.GetStatsRes{
			Strength:     1,
			Dexterity:    2,
			Constitution: 3,
			Intelligence: 4,
			Wisdom:       5,
			Charisma:     6,
		},
		CharacterSkills: "characterSkills",
		Alignment:       "alignmentName",
		Avatar:          "image",
	}

	if !reflect.DeepEqual(character, expected) {
		t.Errorf("expected %+v, got %+v", expected, character)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCharacterByIdRepository_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT c.id, c.hp, c.lvl, c.exp, c.charName, c.sex, c.weight, c.height, c.addLanguage,
						c.ideals, c.weaknesses, c.traits, c.allies, c.organizations, c.enemies, c.story,
						c.goals, c.treasures, c.notes, cl.className, r.raceName, s.subraceName, st.strength,
						st.dexterity, st.constitution, st.intelligence, st.wisdom, st.charisma,
						GROUP_CONCAT(sk.skillName SEPARATOR ', ')
						AS characterSkills, a.alignmentName, i.image FROM characters c
						JOIN accChar ac ON c.id = ac.idChar JOIN class cl ON c.idClass = cl.id
						JOIN races r on c.idRace = r.id JOIN subrace s on c.idSubrace = s.id
						JOIN stats st on c.idStats = st.id JOIN charSkills cs on c.id = cs.idChar
						JOIN skills sk on cs.idSkill = sk.id JOIN alignment a on c.idAlignment = a.id
						JOIN image i ON c.idAvatar = i.id WHERE c.id = ? GROUP BY c.id`)).WillReturnError(errors.New("fake error"))

	repo := NewRepository(db)

	ctx := context.Background()
	_, err = repo.GetCharacterById(ctx, 1)
	if err.Error() != "fake error" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestCreateCharacterRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRepository(db)

	ctx := context.Background()
	char := &CreateCharacterReq{
		Hp:              100,
		Exp:             0,
		CharName:        "Test Character",
		Sex:             true,
		Weight:          100,
		Height:          100,
		Class:           class.Class{Id: 1},
		Race:            races.Races{Id: 1},
		Subrace:         subraces.CreateCharReq{Id: 1},
		Stats:           stats.GetStatsRes{Strength: 10, Dexterity: 10, Constitution: 10, Intelligence: 10, Wisdom: 10, Charisma: 10},
		AddLanguage:     "Common",
		Alignment:       alignment.Alignment{Id: 1},
		Ideals:          "Ideal",
		Weaknesses:      "Weakness",
		Traits:          "Trait",
		Allies:          "Allies",
		Organizations:   "Org",
		Enemies:         "Enemies",
		Story:           "Story",
		Goals:           "Goals",
		Treasures:       "Treasures",
		Notes:           "Notes",
		CharacterSkills: []skills.Skills{{Id: 1, SkillName: "None"}},
		IdAcc:           1,
		Avatar:          "avatar_data",
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO image(image) VALUE (?)`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO stats`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO characters`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO charSkills(idSkill, idChar) VALUES (?, ?)`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE accChar SET act = false WHERE idAccount = ?`)).WithArgs(char.IdAcc).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO accChar(act, idAccount, idChar) VALUES (?, ?, ?)`)).WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateCharacter(ctx, char)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateCharacterRepository_TxError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRepository(db)

	ctx := context.Background()
	char := &CreateCharacterReq{
		Hp:              100,
		Exp:             0,
		CharName:        "Test Character",
		Sex:             true,
		Weight:          100,
		Height:          100,
		Class:           class.Class{Id: 1},
		Race:            races.Races{Id: 1},
		Subrace:         subraces.CreateCharReq{Id: 1},
		Stats:           stats.GetStatsRes{Strength: 10, Dexterity: 10, Constitution: 10, Intelligence: 10, Wisdom: 10, Charisma: 10},
		AddLanguage:     "Common",
		Alignment:       alignment.Alignment{Id: 1},
		Ideals:          "Ideal",
		Weaknesses:      "Weakness",
		Traits:          "Trait",
		Allies:          "Allies",
		Organizations:   "Org",
		Enemies:         "Enemies",
		Story:           "Story",
		Goals:           "Goals",
		Treasures:       "Treasures",
		Notes:           "Notes",
		CharacterSkills: []skills.Skills{{Id: 1, SkillName: "None"}},
		IdAcc:           1,
		Avatar:          "avatar_data",
	}

	mock.ExpectBegin().WillReturnError(errors.New("fake begin tx error"))

	err = repo.CreateCharacter(ctx, char)
	if err.Error() != "fake begin tx error" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCreateCharacterRepository_ImageError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRepository(db)

	ctx := context.Background()
	char := &CreateCharacterReq{
		Hp:              100,
		Exp:             0,
		CharName:        "Test Character",
		Sex:             true,
		Weight:          100,
		Height:          100,
		Class:           class.Class{Id: 1},
		Race:            races.Races{Id: 1},
		Subrace:         subraces.CreateCharReq{Id: 1},
		Stats:           stats.GetStatsRes{Strength: 10, Dexterity: 10, Constitution: 10, Intelligence: 10, Wisdom: 10, Charisma: 10},
		AddLanguage:     "Common",
		Alignment:       alignment.Alignment{Id: 1},
		Ideals:          "Ideal",
		Weaknesses:      "Weakness",
		Traits:          "Trait",
		Allies:          "Allies",
		Organizations:   "Org",
		Enemies:         "Enemies",
		Story:           "Story",
		Goals:           "Goals",
		Treasures:       "Treasures",
		Notes:           "Notes",
		CharacterSkills: []skills.Skills{{Id: 1, SkillName: "None"}},
		IdAcc:           1,
		Avatar:          "avatar_data",
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO image(image) VALUE (?)`)).WillReturnError(errors.New("fake image error"))

	err = repo.CreateCharacter(ctx, char)
	if err.Error() != "fake image error" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCreateCharacterRepository_StatsError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRepository(db)

	ctx := context.Background()
	char := &CreateCharacterReq{
		Hp:              100,
		Exp:             0,
		CharName:        "Test Character",
		Sex:             true,
		Weight:          100,
		Height:          100,
		Class:           class.Class{Id: 1},
		Race:            races.Races{Id: 1},
		Subrace:         subraces.CreateCharReq{Id: 1},
		Stats:           stats.GetStatsRes{Strength: 10, Dexterity: 10, Constitution: 10, Intelligence: 10, Wisdom: 10, Charisma: 10},
		AddLanguage:     "Common",
		Alignment:       alignment.Alignment{Id: 1},
		Ideals:          "Ideal",
		Weaknesses:      "Weakness",
		Traits:          "Trait",
		Allies:          "Allies",
		Organizations:   "Org",
		Enemies:         "Enemies",
		Story:           "Story",
		Goals:           "Goals",
		Treasures:       "Treasures",
		Notes:           "Notes",
		CharacterSkills: []skills.Skills{{Id: 1, SkillName: "None"}},
		IdAcc:           1,
		Avatar:          "avatar_data",
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO image(image) VALUE (?)`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO stats`).WillReturnError(errors.New("fake stats error"))
	mock.ExpectExec(`INSERT INTO characters`).WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateCharacter(ctx, char)
	if err.Error() != "fake stats error" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCreateCharacterRepository_CharacterError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRepository(db)

	ctx := context.Background()
	char := &CreateCharacterReq{
		Hp:              100,
		Exp:             0,
		CharName:        "Test Character",
		Sex:             true,
		Weight:          100,
		Height:          100,
		Class:           class.Class{Id: 1},
		Race:            races.Races{Id: 1},
		Subrace:         subraces.CreateCharReq{Id: 1},
		Stats:           stats.GetStatsRes{Strength: 10, Dexterity: 10, Constitution: 10, Intelligence: 10, Wisdom: 10, Charisma: 10},
		AddLanguage:     "Common",
		Alignment:       alignment.Alignment{Id: 1},
		Ideals:          "Ideal",
		Weaknesses:      "Weakness",
		Traits:          "Trait",
		Allies:          "Allies",
		Organizations:   "Org",
		Enemies:         "Enemies",
		Story:           "Story",
		Goals:           "Goals",
		Treasures:       "Treasures",
		Notes:           "Notes",
		CharacterSkills: []skills.Skills{{Id: 1, SkillName: "None"}},
		IdAcc:           1,
		Avatar:          "avatar_data",
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO image(image) VALUE (?)`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO stats`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO characters`).WillReturnError(errors.New("fake character error"))

	err = repo.CreateCharacter(ctx, char)
	if err.Error() != "fake character error" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCreateCharacterRepository_SkillsError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRepository(db)

	ctx := context.Background()
	char := &CreateCharacterReq{
		Hp:              100,
		Exp:             0,
		CharName:        "Test Character",
		Sex:             true,
		Weight:          100,
		Height:          100,
		Class:           class.Class{Id: 1},
		Race:            races.Races{Id: 1},
		Subrace:         subraces.CreateCharReq{Id: 1},
		Stats:           stats.GetStatsRes{Strength: 10, Dexterity: 10, Constitution: 10, Intelligence: 10, Wisdom: 10, Charisma: 10},
		AddLanguage:     "Common",
		Alignment:       alignment.Alignment{Id: 1},
		Ideals:          "Ideal",
		Weaknesses:      "Weakness",
		Traits:          "Trait",
		Allies:          "Allies",
		Organizations:   "Org",
		Enemies:         "Enemies",
		Story:           "Story",
		Goals:           "Goals",
		Treasures:       "Treasures",
		Notes:           "Notes",
		CharacterSkills: []skills.Skills{{Id: 1, SkillName: "None"}},
		IdAcc:           1,
		Avatar:          "avatar_data",
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO image(image) VALUE (?)`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO stats`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO characters`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO charSkills(idSkill, idChar) VALUES (?, ?)`)).WillReturnError(errors.New("fake skills error"))

	err = repo.CreateCharacter(ctx, char)
	if err.Error() != "fake skills error" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCreateCharacterRepository_AccCharUpdateError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRepository(db)

	ctx := context.Background()
	char := &CreateCharacterReq{
		Hp:              100,
		Exp:             0,
		CharName:        "Test Character",
		Sex:             true,
		Weight:          100,
		Height:          100,
		Class:           class.Class{Id: 1},
		Race:            races.Races{Id: 1},
		Subrace:         subraces.CreateCharReq{Id: 1},
		Stats:           stats.GetStatsRes{Strength: 10, Dexterity: 10, Constitution: 10, Intelligence: 10, Wisdom: 10, Charisma: 10},
		AddLanguage:     "Common",
		Alignment:       alignment.Alignment{Id: 1},
		Ideals:          "Ideal",
		Weaknesses:      "Weakness",
		Traits:          "Trait",
		Allies:          "Allies",
		Organizations:   "Org",
		Enemies:         "Enemies",
		Story:           "Story",
		Goals:           "Goals",
		Treasures:       "Treasures",
		Notes:           "Notes",
		CharacterSkills: []skills.Skills{{Id: 1, SkillName: "None"}},
		IdAcc:           1,
		Avatar:          "avatar_data",
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO image(image) VALUE (?)`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO stats`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO characters`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO charSkills(idSkill, idChar) VALUES (?, ?)`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE accChar SET act = false WHERE idAccount = ?`)).WillReturnError(errors.New("fake update accChar error"))

	err = repo.CreateCharacter(ctx, char)
	if err.Error() != "fake update accChar error" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCreateCharacterRepository_AccCharInsertError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRepository(db)

	ctx := context.Background()
	char := &CreateCharacterReq{
		Hp:              100,
		Exp:             0,
		CharName:        "Test Character",
		Sex:             true,
		Weight:          100,
		Height:          100,
		Class:           class.Class{Id: 1},
		Race:            races.Races{Id: 1},
		Subrace:         subraces.CreateCharReq{Id: 1},
		Stats:           stats.GetStatsRes{Strength: 10, Dexterity: 10, Constitution: 10, Intelligence: 10, Wisdom: 10, Charisma: 10},
		AddLanguage:     "Common",
		Alignment:       alignment.Alignment{Id: 1},
		Ideals:          "Ideal",
		Weaknesses:      "Weakness",
		Traits:          "Trait",
		Allies:          "Allies",
		Organizations:   "Org",
		Enemies:         "Enemies",
		Story:           "Story",
		Goals:           "Goals",
		Treasures:       "Treasures",
		Notes:           "Notes",
		CharacterSkills: []skills.Skills{{Id: 1, SkillName: "None"}},
		IdAcc:           1,
		Avatar:          "avatar_data",
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO image(image) VALUE (?)`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO stats`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO characters`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO charSkills(idSkill, idChar) VALUES (?, ?)`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE accChar SET act = false WHERE idAccount = ?`)).WithArgs(char.IdAcc).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO accChar(act, idAccount, idChar) VALUES (?, ?, ?)`)).WillReturnError(errors.New("fake insert accChar error"))

	err = repo.CreateCharacter(ctx, char)
	if err.Error() != "fake insert accChar error" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestUpdateCharacterHpByIdRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE characters SET hp = ? WHERE id = ?`)).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepository(db)

	ctx := context.Background()
	err = repo.UpdateCharacterHpById(ctx, 1, 1)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateCharacterHpByIdRepository_TxError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin().WillReturnError(errors.New("fake begin tx error"))

	repo := NewRepository(db)

	ctx := context.Background()
	err = repo.UpdateCharacterHpById(ctx, 1, 1)
	if err.Error() != "fake begin tx error" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestUpdateCharacterHpByIdRepository_HpUpdateError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE characters SET hp = ? WHERE id = ?`)).WillReturnError(errors.New("fake hp update error"))

	repo := NewRepository(db)

	ctx := context.Background()
	err = repo.UpdateCharacterHpById(ctx, 1, 1)
	if err.Error() != "fake hp update error" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestUpdateCharacterExpByIdRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE characters SET exp = ? WHERE id = ?`)).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepository(db)

	ctx := context.Background()
	err = repo.UpdateCharacterExpById(ctx, 1, 1)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateCharacterExpByIdRepository_TxError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin().WillReturnError(errors.New("fake begin tx error"))

	repo := NewRepository(db)

	ctx := context.Background()
	err = repo.UpdateCharacterExpById(ctx, 1, 1)
	if err.Error() != "fake begin tx error" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestUpdateCharacterExpByIdRepository_ExpUpdateError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE characters SET exp = ? WHERE id = ?`)).WillReturnError(errors.New("fake exp update error"))

	repo := NewRepository(db)

	ctx := context.Background()
	err = repo.UpdateCharacterExpById(ctx, 1, 1)
	if err.Error() != "fake exp update error" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestSetActiveCharacterByIdRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE accChar SET act = false WHERE idAccount = ?`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE accChar SET act = true WHERE idChar = ?`)).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepository(db)

	ctx := context.Background()
	err = repo.SetActiveCharacterById(ctx, &SetActiveCharReq{IdAcc: 1, IdChar: 1})
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSetActiveCharacterByIdRepository_TxError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin().WillReturnError(errors.New("fake begin tx error"))

	repo := NewRepository(db)

	ctx := context.Background()
	err = repo.SetActiveCharacterById(ctx, &SetActiveCharReq{IdAcc: 1, IdChar: 1})
	if err.Error() != "fake begin tx error" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestSetActiveCharacterByIdRepository_SetFalseError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE accChar SET act = false WHERE idAccount = ?`)).WillReturnError(errors.New("fake set false error"))

	repo := NewRepository(db)

	ctx := context.Background()
	err = repo.SetActiveCharacterById(ctx, &SetActiveCharReq{IdAcc: 1, IdChar: 1})
	if err.Error() != "fake set false error" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestSetActiveCharacterByIdRepository_SetTrueError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE accChar SET act = false WHERE idAccount = ?`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE accChar SET act = true WHERE idChar = ?`)).WillReturnError(errors.New("fake set true error"))

	repo := NewRepository(db)

	ctx := context.Background()
	err = repo.SetActiveCharacterById(ctx, &SetActiveCharReq{IdAcc: 1, IdChar: 1})
	if err.Error() != "fake set true error" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestGetAllCharactersByAccIdService(t *testing.T) {
	mockRepo := &MockRepo{}
	service := NewService(mockRepo)

	ctx := context.Background()
	characters, err := service.GetAllCharactersByAccId(ctx, 1)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	expected := []GetAllCharactesRes{
		{IdChar: 1, CharName: "First", Avatar: "avatar_data"},
		{IdChar: 2, CharName: "Second", Avatar: "avatar_data"}}

	if !reflect.DeepEqual(characters, expected) {
		t.Errorf("expected %+v, got %+v", expected, characters)
	}
}

func TestGetAllCharactersByAccIdService_RepoError(t *testing.T) {
	mockRepo := &MockRepoError{}
	service := NewService(mockRepo)

	ctx := context.Background()
	_, err := service.GetAllCharactersByAccId(ctx, 1)
	if err.Error() != "fake repo error" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestGetCharacterByIdService(t *testing.T) {
	mockRepo := &MockRepo{}
	service := NewService(mockRepo)

	ctx := context.Background()
	character, err := service.GetCharacterById(ctx, 1)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	expected := &Character{
		Id:            1,
		Hp:            100,
		Lvl:           1,
		Exp:           0,
		CharName:      "charName",
		Sex:           true,
		Weight:        100,
		Height:        100,
		AddLanguage:   "addLanguage",
		Ideals:        "ideals",
		Weaknesses:    "weaknesses",
		Traits:        "traits",
		Allies:        "allies",
		Organizations: "organizations",
		Enemies:       "enemies",
		Story:         "story",
		Goals:         "goals",
		Treasures:     "treasures",
		Notes:         "notes",
		Class:         "className",
		Race:          "raceName",
		Subrace:       "subraceName",
		Stats: stats.GetStatsRes{
			Strength:     1,
			Dexterity:    2,
			Constitution: 3,
			Intelligence: 4,
			Wisdom:       5,
			Charisma:     6,
		},
		CharacterSkills: "characterSkills",
		Alignment:       "alignmentName",
		Avatar:          "image",
	}

	if !reflect.DeepEqual(character, expected) {
		t.Errorf("expected %+v, got %+v", expected, character)
	}
}

func TestGetCharacterByIdService_RepoError(t *testing.T) {
	mockRepo := &MockRepoError{}
	service := NewService(mockRepo)

	ctx := context.Background()
	_, err := service.GetCharacterById(ctx, 1)
	if err.Error() != "fake repo error" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestCreateCharacterService(t *testing.T) {
	mockRepo := &MockRepo{}
	service := NewService(mockRepo)
	character := &CreateCharacterReq{
		Hp:            100,
		Exp:           0,
		CharName:      "charName",
		Sex:           true,
		Weight:        100,
		Height:        100,
		AddLanguage:   "addLanguage",
		Ideals:        "ideals",
		Weaknesses:    "weaknesses",
		Traits:        "traits",
		Allies:        "allies",
		Organizations: "organizations",
		Enemies:       "enemies",
		Story:         "story",
		Goals:         "goals",
		Treasures:     "treasures",
		Notes:         "notes",
		Class:         class.Class{Id: 1},
		Race:          races.Races{Id: 1},
		Subrace:       subraces.CreateCharReq{Id: 1},
		Stats: stats.GetStatsRes{
			Strength:     1,
			Dexterity:    2,
			Constitution: 3,
			Intelligence: 4,
			Wisdom:       5,
			Charisma:     6,
		},
		CharacterSkills: []skills.Skills{{Id: 1}},
		Alignment:       alignment.Alignment{Id: 1},
		Avatar:          "image",
	}
	ctx := context.Background()
	err := service.CreateCharacter(ctx, character)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestCreateCharacterService_RepoError(t *testing.T) {
	mockRepo := &MockRepoError{}
	service := NewService(mockRepo)
	character := &CreateCharacterReq{
		Hp:            100,
		Exp:           0,
		CharName:      "charName",
		Sex:           true,
		Weight:        100,
		Height:        100,
		AddLanguage:   "addLanguage",
		Ideals:        "ideals",
		Weaknesses:    "weaknesses",
		Traits:        "traits",
		Allies:        "allies",
		Organizations: "organizations",
		Enemies:       "enemies",
		Story:         "story",
		Goals:         "goals",
		Treasures:     "treasures",
		Notes:         "notes",
		Class:         class.Class{Id: 1},
		Race:          races.Races{Id: 1},
		Subrace:       subraces.CreateCharReq{Id: 1},
		Stats: stats.GetStatsRes{
			Strength:     1,
			Dexterity:    2,
			Constitution: 3,
			Intelligence: 4,
			Wisdom:       5,
			Charisma:     6,
		},
		CharacterSkills: []skills.Skills{{Id: 1}},
		Alignment:       alignment.Alignment{Id: 1},
		Avatar:          "image",
	}
	ctx := context.Background()
	err := service.CreateCharacter(ctx, character)
	if err.Error() != "fake repo error" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestSetActiveCharacterByIdService(t *testing.T) {
	mockRepo := &MockRepo{}
	service := NewService(mockRepo)
	character := &SetActiveCharReq{
		IdAcc:  1,
		IdChar: 1,
	}
	ctx := context.Background()
	err := service.SetActiveCharacterById(ctx, character)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestSetActiveCharacterByIdService_RepoError(t *testing.T) {
	mockRepo := &MockRepoError{}
	service := NewService(mockRepo)
	character := &SetActiveCharReq{
		IdAcc:  1,
		IdChar: 1,
	}
	ctx := context.Background()
	err := service.SetActiveCharacterById(ctx, character)
	if err.Error() != "fake repo error" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestGetAllCharactersByAccIdHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetAllCharactersByAccId(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content-type header: got %v want %v", contentType, expectedContentType)
	}

	var characters []GetAllCharactesRes
	if err := json.Unmarshal(rr.Body.Bytes(), &characters); err != nil {
		t.Errorf("error unmarshalling response body: %v", err)
	}
}

func TestGetAllCharactersByAccIdHandler_TokenError(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: errors.New("token error")}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetAllCharactersByAccId(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}
}

func TestGetAllCharactersByAccIdHandler_ServiceError(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockRepoError{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetAllCharactersByAccId(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}

func TestGetCharacterByIdHandler(t *testing.T) {
	requestBody := []byte(`{"id": 1}`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetCharacterById(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content-type header: got %v want %v", contentType, expectedContentType)
	}

	var character *Character
	if err := json.Unmarshal(rr.Body.Bytes(), &character); err != nil {
		t.Errorf("error unmarshalling response body: %v", err)
	}
}

func TestGetCharacterByIdHandler_TokenError(t *testing.T) {
	requestBody := []byte(`{"id": 1}`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: errors.New("fake token error")}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetCharacterById(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}
}

func TestGetCharacterByIdHandler_BadRequest(t *testing.T) {
	requestBody := []byte(`{"i`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetCharacterById(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestGetCharacterByIdHandler_ServiceError(t *testing.T) {
	requestBody := []byte(`{"id": 1}`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockServiceError{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetCharacterById(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}

func TestCreateCharacterHandler(t *testing.T) {
	requestBody := []byte(`{
		"hp": 100,
		"exp": 0,
		"avatar": "",
		"charName": "name",
		"sex": false,
		"weight": 12,
		"height": 21,
		"charClass": {
			"id": 1,
			"className": "Бард"
		},
		"race": {
			"id": 2,
			"raceName": "Дворф"
		},
		"subrace": {
			"id": 2,
			"raceName": "Фул",
			"stats": {
				"strength": 2,
				"dexterity": 2,
				"constitution": 2,
				"intelligence": 2,
				"wisdom": 2,
				"charisma": 2
			}
		},
		"stats": {
			"strength": 13,
			"dexterity": 0,
			"constitution": 0,
			"intelligence": 0,
			"wisdom": 0,
			"charisma": 0
		},
		"addLanguage": "ggg",
		"characterSkills": [
		{
			"id": 2,
			"skillName": "Акробатика"
		}
		],
		"alignment": {
			"id": 1,
			"alignmentName": ""
		},
		"ideals": "",
		"weaknesses": "",
		"traits": "",
		"allies": "",
		"organizations": "",
		"enemies": "",
		"story": "",
		"goals": "",
		"treasures": "",
		"notes": ""
	}`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.CreateCharacter(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content-type header: got %v want %v", contentType, expectedContentType)
	}
}

func TestCreateCharacterHandler_BadRequest(t *testing.T) {
	requestBody := []byte(`{
		"hp": 100
		"notes": ""
	}`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.CreateCharacter(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestCreateCharacterHandler_TokenError(t *testing.T) {
	requestBody := []byte(`{
		"hp": 100,
		"exp": 0,
		"avatar": "",
		"charName": "name",
		"sex": false,
		"weight": 12,
		"height": 21,
		"charClass": {
			"id": 1,
			"className": "Бард"
		},
		"race": {
			"id": 2,
			"raceName": "Дворф"
		},
		"subrace": {
			"id": 2,
			"raceName": "Фул",
			"stats": {
				"strength": 2,
				"dexterity": 2,
				"constitution": 2,
				"intelligence": 2,
				"wisdom": 2,
				"charisma": 2
			}
		},
		"stats": {
			"strength": 13,
			"dexterity": 0,
			"constitution": 0,
			"intelligence": 0,
			"wisdom": 0,
			"charisma": 0
		},
		"addLanguage": "ggg",
		"characterSkills": [
		{
			"id": 2,
			"skillName": "Акробатика"
		}
		],
		"alignment": {
			"id": 1,
			"alignmentName": ""
		},
		"ideals": "",
		"weaknesses": "",
		"traits": "",
		"allies": "",
		"organizations": "",
		"enemies": "",
		"story": "",
		"goals": "",
		"treasures": "",
		"notes": ""
	}`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: errors.New("fake token error")}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.CreateCharacter(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}
}

func TestCreateCharacterHandler_ServiceError(t *testing.T) {
	requestBody := []byte(`{
		"hp": 100,
		"exp": 0,
		"avatar": "",
		"charName": "name",
		"sex": false,
		"weight": 12,
		"height": 21,
		"charClass": {
			"id": 1,
			"className": "Бард"
		},
		"race": {
			"id": 2,
			"raceName": "Дворф"
		},
		"subrace": {
			"id": 2,
			"raceName": "Фул",
			"stats": {
				"strength": 2,
				"dexterity": 2,
				"constitution": 2,
				"intelligence": 2,
				"wisdom": 2,
				"charisma": 2
			}
		},
		"stats": {
			"strength": 13,
			"dexterity": 0,
			"constitution": 0,
			"intelligence": 0,
			"wisdom": 0,
			"charisma": 0
		},
		"addLanguage": "ggg",
		"characterSkills": [
		{
			"id": 2,
			"skillName": "Акробатика"
		}
		],
		"alignment": {
			"id": 1,
			"alignmentName": ""
		},
		"ideals": "",
		"weaknesses": "",
		"traits": "",
		"allies": "",
		"organizations": "",
		"enemies": "",
		"story": "",
		"goals": "",
		"treasures": "",
		"notes": ""
	}`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockServiceError{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.CreateCharacter(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}

func TestSetActiveCharacterByIdHandler(t *testing.T) {
	requestBody := []byte(`{"idAcc": 1, "idChar": 1}`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.SetActiveCharacterById(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content-type header: got %v want %v", contentType, expectedContentType)
	}
}

func TestSetActiveCharacterByIdHandler_BadRequest(t *testing.T) {
	requestBody := []byte(`{"idAcc": 1, "idChar: 1}`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.SetActiveCharacterById(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestSetActiveCharacterByIdHandler_TokenError(t *testing.T) {
	requestBody := []byte(`{"idAcc": 1, "idChar": 1}`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: errors.New("fake token error")}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.SetActiveCharacterById(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}
}

func TestSetActiveCharacterByIdHandler_ServiceError(t *testing.T) {
	requestBody := []byte(`{"idAcc": 1, "idChar": 1}`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockRepoError{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.SetActiveCharacterById(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}
