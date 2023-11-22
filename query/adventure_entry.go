package query

const AdventureEntryCreate = `INSERT INTO adventure_template_entry (
	id,
	template_id
) VALUES (
	:id,
	:template_id
)`

const AdventureEntry = "SELECT * FROM adventure_template_entry WHERE id = ?"

const AdventureEntryList = "SELECT * FROM adventure_template_entry WHERE template_id = ? ORDER BY id ASC"

const AdventureEntryUpdate = `UPDATE adventure_template_entry SET
	id=:id,
	template_id=:template_id
WHERE id = :id`

const AdventureEntryDelete = "DELETE FROM adventure_template_entry WHERE id = ?"
