package main

import (
	"github.com/astaxie/beego/migration"
	//"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Users_20221024_133803 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Users_20221024_133803{}
	m.Created = "20221024_133803"

	migration.Register("Users_20221024_133803", m)
}

// Run the migrations
func (m *Users_20221024_133803) Up() {
	m.SQL(`
	CREATE TABLE IF NOT EXISTS public.users (
		"id" serial NOT NULL,
		"email" varchar NULL,
		"mobile" varchar NULL,
		"user_name" varchar NULL,
		"password" varchar NULL,
		"status" bool NULL DEFAULT true,
		"created_at" timestamp NOT NULL,
		"updated_at" timestamp NOT NULL
	)
	`)
	m.SQL(`
	CREATE UNIQUE INDEX IF NOT EXISTS user_email_idx ON public.users (email)
	`)
	m.SQL(`
	CREATE UNIQUE INDEX IF NOT EXISTS user_user_nameidx ON public.users (user_name)
	`)
	m.SQL(`
	CREATE UNIQUE INDEX IF NOT EXISTS user_user_mobile_idx ON public.users (mobile)
	`)

}

// Reverse the migrations
func (m *Users_20221024_133803) Down() {
	 //use m.SQL("DROP TABLE users") //to reverse schema update

}
