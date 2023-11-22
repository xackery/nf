package query

const AARankEffectCreate = `INSERT INTO aa_rank_effects (
rank_id
slot
effect_id
base1
base2
) VALUES (
:rank_id
:slot
:effect_id
:base1
:base2
)`

const AARankEffect = "SELECT * FROM aa_rank_effects WHERE rank_id = ? AND slot = ?"

const AARankEffectList = "SELECT * FROM aa_rank_effects WHERE rank_id = ? ORDER BY slot ASC"

const AARankEffectUpdate = `UPDATE aa_rank_effects SET
slot = :slot
effect_id = :effect_id
base1 = :base1
base2 = :base2
WHERE rank_id = :rank_id AND slot = :slot
`

const AARankEffectDelete = "DELETE FROM aa_rank_effects WHERE rank_id = ? AND slot = ?"
