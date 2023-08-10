package Measurement

import (
	"healthtracker/Database"
	"math"
	"strings"
	"time"
)

type Measurement struct {
	Id            string
	Steps         int
	SleepDuration int
	SleepQuality  int
	Mood          int
	Energy        int
}

func GetOne(Id string) (*Measurement, error) {
	row := Database.DB.QueryRow("select * from measurements where id = ? limit 1", Id)

	temp := new(Measurement)

	err := row.Scan(&temp.Id, &temp.Steps, &temp.SleepDuration, &temp.SleepQuality, &temp.Mood, &temp.Energy)
	if err != nil {
		return nil, err
	}

	s := strings.Split(temp.Id, "T")
	temp.Id = s[0]
	return temp, nil
}

func GetAll(page int) ([]*Measurement, error) {
	rows, _ := Database.DB.Query("select * from measurements order by id desc limit ?, 10", 10*(page-1))

	var all []*Measurement

	for rows.Next() {
		temp := new(Measurement)
		err := rows.Scan(&temp.Id, &temp.Steps, &temp.SleepDuration, &temp.SleepQuality, &temp.Mood, &temp.Energy)
		if err != nil {
			return nil, err
		}

		s := strings.Split(temp.Id, "T")
		temp.Id = s[0]
		all = append(all, temp)
	}

	return all, nil
}

func Delete(Id string) error {
	begin, err := Database.DB.Begin()
	if err != nil {
		return err
	}

	_, err = begin.Exec("delete from measurements where id = ?", Id)
	if err != nil {
		err := begin.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	err = begin.Commit()
	if err != nil {
		err := begin.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	return nil
}

func Create(m *Measurement) error {
	begin, err := Database.DB.Begin()
	if err != nil {
		return err
	}

	_, err = begin.Exec(`insert into measurements(
                         id,
                         steps, 
                         "sleepDuration", 
                         "sleepQuality", 
                         mood, 
                         energy
                         ) values (?,?,?,?,?,?)`,
		time.Now().Format("2006-01-02"),
		m.Steps,
		m.SleepDuration,
		m.SleepQuality,
		m.Mood,
		m.Energy,
	)
	if err != nil {
		err := begin.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	err = begin.Commit()
	if err != nil {
		err := begin.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	return nil
}

func Edit(m *Measurement, id string) error {
	begin, err := Database.DB.Begin()
	if err != nil {
		return err
	}

	_, err = begin.Exec(`update measurements set 
                         steps = ?, 
                         "sleepDuration" = ?, 
                         "sleepQuality" = ?, 
                         mood = ?, 
                         energy = ?
                         where id = ?`,
		m.Steps,
		m.SleepDuration,
		m.SleepQuality,
		m.Mood,
		m.Energy,
		id,
	)
	if err != nil {
		err := begin.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	err = begin.Commit()
	if err != nil {
		err := begin.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	return nil
}

func DoneToday() bool {
	m, err := GetOne(time.Now().Format("2006-01-02"))
	if err != nil {
		return false
	}

	return m != nil
}

func Count() int {
	row := Database.DB.QueryRow("select count(*) from measurements")

	temp := 0

	err := row.Scan(&temp)
	if err != nil {
		return 0
	}

	temp = int(math.Ceil(float64(temp) / 10))

	return temp
}
