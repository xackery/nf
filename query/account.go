package query

const AccountCreate = `INSERT INTO account (
id
name
charname
sharedplat
password
status
ls_id
lsaccount_id
gmspeed
invulnerable
flymode
ignore_tells
revoked
karma
minilogin_ip
hideme
rulesflag
suspendeduntil
time_creation
ban_reason
suspend_reason
crc_eqgame
crc_skillcaps
crc_basedata
) VALUES (
:id
:name
:charname
:sharedplat
:password
:status
:ls_id
:lsaccount_id
:gmspeed
:invulnerable
:flymode
:ignore_tells
:revoked
:karma
:minilogin_ip
:hideme
:rulesflag
:suspendeduntil
:time_creation
:ban_reason
:suspend_reason
:crc_eqgame
:crc_skillcaps
:crc_basedata
)`

const Account = "SELECT * FROM account WHERE id = ?"

const AccountList = "SELECT * FROM account ORDER BY id ASC"

const AccountUpdate = `UPDATE account SET
id=:id,
name=:name,
charname=:charname,
sharedplat=:sharedplat,
password=:password,
status=:status,
ls_id=:ls_id,
lsaccount_id=:lsaccount_id,
gmspeed=:gmspeed,
invulnerable=:invulnerable,
flymode=:flymode,
ignore_tells=:ignore_tells,
revoked=:revoked,
karma=:karma,
minilogin_ip=:minilogin_ip,
hideme=:hideme,
rulesflag=:rulesflag,
suspendeduntil=:suspendeduntil,
time_creation=:time_creation,
ban_reason=:ban_reason,
suspend_reason=:suspend_reason,
crc_eqgame=:crc_eqgame,
crc_skillcaps=:crc_skillcaps,
crc_basedata=:crc_basedata
WHERE id = :id`

const AccountDelete = "DELETE FROM account WHERE id = ?"
