package migrations

import (
	transporter "github.com/wawandco/transporter/core"
)

func init() {
	migration := transporter.Migration{
		Identifier: 1464969180108881647,
		Up: func(tx *transporter.Tx) {
			tx.Exec(`
        CREATE TABLE  IF NOT EXISTS baq_gophers (
          id integer NOT NULL,
          name character varying(255),
          company character varying(255),
        );

        CREATE SEQUENCE baq_gophers_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;
        ALTER SEQUENCE baq_gophers_id_seq OWNED BY baq_gophers.id;
        ALTER TABLE ONLY baq_gophers ALTER COLUMN id SET DEFAULT nextval('baq_gophers_id_seq'::regclass);
        ALTER TABLE ONLY baq_gophers ADD CONSTRAINT baq_gophers_pkey PRIMARY KEY (id);
        `)
		},
		Down: func(tx *transporter.Tx) {
			tx.Exec("")
		},
	}

	//Register the migration to run up or down acordingly.
	transporter.Add(migration)
}
