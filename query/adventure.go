package query

const AdventureCreate = `INSERT INTO adventure_template (
zone,
zone_version,
is_hard,
is_raid,
min_level,
max_level,
type,
type_data,
type_count,
assa_x,
assa_y,
assa_z,
assa_h,
text,
duration,
zone_in_time,
win_points,
lose_points,
theme,
zone_in_zone_id,
zone_in_x,
zone_in_y,
zone_in_object_id,
dest_x,
dest_y,
dest_z,
dest_h,
graveyard_zone_id,
graveyard_x,
graveyard_y,
graveyard_z,
graveyard_radius
) VALUES (
:zone, 
:zone_version, 
:is_hard, 
:is_raid, 
:min_level, 
:max_level, 
:type, 
:type_data, 
:type_count, 
:assa_x, 
:assa_y, 
:assa_z, 
:assa_h, 
:text, 
:duration, 
:zone_in_time, 
:win_points, 
:lose_points, 
:theme, 
:zone_in_zone_id, 
:zone_in_x, 
:zone_in_y, 
:zone_in_object_id, 
:dest_x, 
:dest_y, 
:dest_z, 
:dest_h, 
:graveyard_zone_id, 
:graveyard_x, 
:graveyard_y, 
:graveyard_z, 
:graveyard_radius
)`

const Adventure = "SELECT * FROM adventure_template WHERE id = ?"

const AdventureList = "SELECT * FROM adventure_template ORDER BY id ASC"

const AdventureUpdate = `UPDATE adventure_template SET
zone = :zone,
zone_version = :zone_version,
is_hard = :is_hard,
is_raid = :is_raid,
min_level = :min_level,
max_level = :max_level,
type = :type,
type_data = :type_data,
type_count = :type_count,
assa_x = :assa_x,
assa_y = :assa_y,
assa_z = :assa_z,
assa_h = :assa_h,
text = :text,
duration = :duration,
zone_in_time = :zone_in_time,
win_points = :win_points,
lose_points = :lose_points,
theme = :theme,
zone_in_zone_id = :zone_in_zone_id,
zone_in_x = :zone_in_x,
zone_in_y = :zone_in_y,
zone_in_object_id = :zone_in_object_id,
dest_x = :dest_x,
dest_y = :dest_y,
dest_z = :dest_z,
dest_h = :dest_h,
graveyard_zone_id = :graveyard_zone_id,
graveyard_x = :graveyard_x,
graveyard_y = :graveyard_y,
graveyard_z = :graveyard_z,
graveyard_radius = :graveyard_radius
WHERE id = :id`

const AdventureDelete = "DELETE FROM adventure_template WHERE id = ?"
