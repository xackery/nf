package query

const AARankPrereqCreate = `INSERT INTO aa_rank_prereqs (
rank_id,
aa_id,
points
) VALUES (
:rank_id,
:aa_id,
:points
)`

const AARankPrereq = "SELECT * FROM aa_rank_prereqs WHERE rank_id = ?"

const AARankPrereqList = "SELECT * FROM aa_rank_prereqs WHERE rank_id = ? ORDER BY aa_id ASC"

const AARankPrereqUpdate = `UPDATE aa_rank_prereqs SET
points = :points
WHERE rank_id = :rank_id
`

const AARankPrereqDelete = "DELETE FROM aa_rank_prereqs WHERE rank_id = ?"
