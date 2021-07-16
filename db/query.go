package db

import (
	"app/structures"
	"fmt"

	"github.com/jmoiron/sqlx"
)

var (
	buildingAttributesGetSQL string = `
	SELECT 
		a.fid, 
		b.ddf, 
		a.ffh, 
		a.bldg_rep_cost, 
		a.cont_rep_cost, 
		b.event_type,
		b.epoch,
		b.dg + b.wv as depth 
	FROM properties.buildings_attributed a
	INNER JOIN properties.buildings_depth as b
	ON a.fid = b.fid
	WHERE 
		a.ffh IS NOT NULL AND
		a.fid >= 0 AND
		b.dg IS NOT NULL AND
		b.dg != 'NaN' AND
		b.wv != 'NaN';`

	buildingLossUpsertSQL string = `	
	INSERT into properties.buildings_loss(fid, 
										epoch, 
										event_type, 
										structure_damage_percent, 
										structure_damage_value, 
										content_damage_percent,
										content_damage_value)
							VALUES ($1, $2, $3, $4, $5, $6, $7) 
	ON CONFLICT (fid, epoch, event_type) 
	DO
	UPDATE SET 
		structure_damage_percent = EXCLUDED.structure_damage_percent,
		structure_damage_value = EXCLUDED.structure_damage_value,
		content_damage_percent = EXCLUDED.content_damage_percent,
		content_damage_value = EXCLUDED.content_damage_value;`

	buildingLossUpsertBatchSQL string = `	
	INSERT into properties.buildings_loss(fid, 
										epoch, 
										event_type, 
										structure_damage_percent, 
										structure_damage_value, 
										content_damage_percent,
										content_damage_value)
							VALUES (:fid, 
									:epoch, 
									:event_type, 
									:structure_damage_percent, 
									:structure_damage_value, 
									:content_damage_percent,
									:content_damage_value) 
	ON CONFLICT (fid, epoch, event_type) 
	DO
	UPDATE SET 
		structure_damage_percent = EXCLUDED.structure_damage_percent,
		structure_damage_value = EXCLUDED.structure_damage_value,
		content_damage_percent = EXCLUDED.content_damage_percent,
		content_damage_value = EXCLUDED.content_damage_value;`
)

func QueryBuildingAttributes(db *sqlx.DB) ([]structures.StructureSimpleDeterministic, error) {

	rows, err := db.Queryx(buildingAttributesGetSQL)

	if err != nil {
		return []structures.StructureSimpleDeterministic{}, err
	}

	defer rows.Close()
	result := make([]structures.StructureSimpleDeterministic, 0)
	for rows.Next() {
		var s structures.StructureSimpleDeterministic
		err = rows.StructScan(&s)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, s)
	}
	return result, nil
}

func UpsertBuildingLoss(ssd structures.StructureSimpleDeterministic, db *sqlx.DB) error {

	_, err := db.Exec(buildingLossUpsertSQL,
		ssd.FID,
		ssd.Epoch,
		ssd.Event,
		ssd.StructureDamagePercent,
		ssd.StructureDamageValue,
		ssd.ContentDamagePercent,
		ssd.ContentDamageValue)

	if err != nil {
		return err
	}

	return nil
}

func UpsertBuildingBatchLoss(ssd []structures.StructureSimpleDeterministicResult, db *sqlx.DB) error {

	_, err := db.NamedExec(buildingLossUpsertBatchSQL, ssd)

	if err != nil {
		return err
	}

	return nil
}
