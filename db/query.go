package db

import (
	"app/structures"
	"fmt"

	"github.com/jmoiron/sqlx"
)

var (
	buildingAttributesGetSQL string = `
	SELECT 
		a.uid, 
		a.ddf, 
		a.ffh, 
		a.structure_value, 
		a.content_value, 
		b.event_type,
		b.epoch,
		b.dg + b.wv as depth 
	FROM geo_context.buildings_loss_attributed a
	INNER JOIN summary.buildings_depth as b
	ON a.uid = b.uid
	WHERE 
		a.ffh IS NOT NULL AND
		a.uid < 50000 AND
		b.dg IS NOT NULL AND
		b.dg != 'NaN' AND
		b.wv != 'NaN';`

	buildingLossUpsertSQL string = `	
	INSERT into summary.buildings_loss(uid, 
										epoch, 
										event_type, 
										structure_damage_percent, 
										structure_damage_value, 
										content_damage_percent,
										content_damage_value)
							VALUES ($1, $2, $3, $4, $5, $6, $7) 
	ON CONFLICT (uid, epoch, event_type) 
	DO
	UPDATE SET 
		structure_damage_percent = EXCLUDED.structure_damage_percent,
		structure_damage_value = EXCLUDED.structure_damage_value,
		content_damage_percent = EXCLUDED.content_damage_percent,
		content_damage_value = EXCLUDED.content_damage_value;`

	buildingLossUpsertBatchSQL string = `	
	INSERT into summary.buildings_loss(uid, 
										epoch, 
										event_type, 
										structure_damage_percent, 
										structure_damage_value, 
										content_damage_percent,
										content_damage_value)
							VALUES (:uid, 
									:epoch, 
									:event_type, 
									:structure_damage_percent, 
									:structure_damage_value, 
									:content_damage_percent,
									:content_damage_value) 
	ON CONFLICT (uid, epoch, event_type) 
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
