package query

const AAAbilityCreate = `INSERT INTO aa_ability (
id,
name,
category,
classes,
races,
drakkin_heritage,
deities,
status,
type,
charges,
grant_only,
first_rank_id,
enabled,
reset_on_death
) VALUES (
:id,
:name,
:category,
:classes,
:races,
:drakkin_heritage,
:deities,
:status,
:type,
:charges,
:grant_only,
:first_rank_id,
:enabled,
:reset_on_death
)`

const AAAbility = "SELECT * FROM aa_ability WHERE id = ?"

const AAAbilityUpdate = `UPDATE aa_ability SET
id=:id, 
name=:name, 
category=:category, 
classes=:classes, 
races=:races, 
drakkin_heritage=:drakkin_heritage, 
deities=:deities, 
status=:status, 
type=:type, 
charges=:charges, 
grant_only=:grant_only, 
first_rank_id=:first_rank_id, 
enabled=:enabled, 
reset_on_death=:reset_on_death
WHERE id = :id`

const AAAbilityDelete = "DELETE FROM aa_ability WHERE id = ?"
