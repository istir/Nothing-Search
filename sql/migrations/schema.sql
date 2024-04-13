-- Create "files" table
CREATE TABLE files (id integer NULL, name text NULL, path integer NULL, thumbnail_exists integer NULL DEFAULT 0, date_modified integer NULL DEFAULT 0, date_created integer NULL DEFAULT 0, PRIMARY KEY (id));
-- Create index "files_id" to table: "files"
CREATE UNIQUE INDEX files_id ON files (id);
-- Create index "files_path" to table: "files"
CREATE UNIQUE INDEX files_path ON files (path);
