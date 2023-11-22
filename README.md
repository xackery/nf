# nf
Neriak Forge


# Dummy sqlite
For doing edits and general testing, it's handy to dump a sqlite copy of the database.  This is done with the following commands:
- `pip install mysql-to-sqlite3` refers to [this](https://github.com/techouse/mysql-to-sqlite3)
- ` mysql2sqlite -f peq.db -d peq -h 10.0.0.4 -u eqemu -p -K` change settings accordingly, but basically dumps peq.db, the -K is to ensure indexes are unique (without it you'll get errors on a eqemu dump)
- peq.db will be about 280MB, I suggest making a zip copy (~76MB) and using it to restore when testing destructive calls
- `make run` already sets EQ_SQLITE=bin/peq.db, so just copy your peq.db to the bin subfolder.

# Example usage
- `curl localhost:11000/api/v1/aas/999` get aa (no results)
- `curl -X POST localhost:11000/api/v1/aas -d "{\"aa\":{\"id\":999,\"name\":\"Innate Run Speed\",\"category\":-1,\"classes\":65535,\"races\":65535,\"drakkinHeritage\":127,\"deities\":131071,\"status\":0,\"type\":1,\"charges\":0,\"grantOnly\":false,\"firstRankId\":1,\"enabled\":true,\"resetOnDeath\":false}}"` create aa
- `curl localhost:11000/api/v1/aas/999` get aa (should result proper)
- `curl -X PUT localhost:11000/api/v1/aas/999 -d "{\"aa\":{\"name\":\"Innate Run Speed Test\",\"category\":-1,\"classes\":65535,\"races\":65535,\"drakkinHeritage\":127,\"deities\":131071,\"status\":0,\"type\":1,\"charges\":0,\"grantOnly\":false,\"firstRankId\":1,\"enabled\":true,\"resetOnDeath\":false}}"` update aa
- `curl localhost:11000/api/v1/aas/999` get aa (should have name "Innate Run Speed Test")
- `curl -X DELETE localhost:11000/api/v1/aas/999` delete aa
- `curl -X DELETE localhost:11000/api/v1/aas/999` delete aa (no rows in result set)
- `curl localhost:11000/api/v1/aas/999` get aa (no results)