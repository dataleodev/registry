package sql

const (
	UsersSelectAll   = "SELECT * FROM users;"
	UserSelectById   = "SELECT * FROM users WHERE id=$1;"
	UserDelete       = "DELETE FROM users WHERE id=$1;"
	UserInsertNew    = "INSERT INTO users (id,name,email,region,password,created) VALUES($1,$2,$3,$4,$5,$6);"
	UserUpdateGroup  = "UPDATE users SET ugroup = $2 WHERE id = $1;"
	UserUpdateRegion = "UPDATE users SET region = $2 WHERE id = $1;"
	UserUpdateRandG  = "UPDATE users SET ugroup = $2, region = $3 WHERE id = $1;"
	RegionAddNew     = "INSERT INTO regions (id, name,description) VALUES ($1,$2,$3);"
	RegionsSelectAll = "SELECT * FROM regions;"
	NodeDelete       = "DELETE FROM nodes WHERE id=$1;"
	NodeGetById      = "SELECT * FROM nodes WHERE id=$1 or addr=$1;"
	NodeGetAll       = "SELECT * FROM nodes;"
	NodeAddNew       = "INSERT INTO nodes (id, addr, name, type, region,lat,long,created, master)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9);"
)
