package query

const AccountIPCreate = `INSERT INTO account_ip (
accid
ip
count
lastused
) VALUES (
:accid
:ip
:count
:lastused
)`

const AccountIP = "SELECT * FROM account_ip WHERE accid = ?"

const AccountIPList = "SELECT * FROM account_ip ORDER BY accid ASC"

const AccountIPUpdate = `UPDATE account_ip SET
accid=:accid,
ip=:ip,
count=:count,
lastused=:lastused
WHERE accid = :accid`

const AccountIPDelete = "DELETE FROM account_ip WHERE accid = ?"
