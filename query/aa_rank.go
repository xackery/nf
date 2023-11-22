package query

const AARankCreate = `INSERT INTO aa_ranks (
id,
upper_hotkey_sid,
lower_hotkey_sid,
title_sid,
desc_sid,
cost,
level_req,
spell,
spell_type,
recast_time,
expansion,
prev_id,
next_id
) VALUES (
:id,
:upper_hotkey_sid,
:lower_hotkey_sid,
:title_sid,
:desc_sid,
:cost,
:level_req,
:spell,
:spell_type,
:recast_time,
:expansion,
:prev_id,
:next_id
)
`

const AARank = "SELECT * FROM aa_ranks WHERE id = ?"

const AARankList = "SELECT * FROM aa_ranks WHERE aa_id = ? ORDER BY id ASC"

const AARankUpdate = `UPDATE aa_ranks SET
upper_hotkey_sid = :upper_hotkey_sid,
lower_hotkey_sid = :lower_hotkey_sid,
title_sid = :title_sid,
desc_sid = :desc_sid,
cost = :cost,
level_req = :level_req,
spell = :spell,
spell_type = :spell_type,
recast_time = :recast_time,
expansion = :expansion,
prev_id = :prev_id,
next_id = :next_id
WHERE id = :id
`

const AARankDelete = "DELETE FROM aa_ranks WHERE id = ?"
