-- Create "files" table
CREATE TABLE files (id integer , name text , path text , thumbnail_exists integer  DEFAULT 0, date_modified integer  DEFAULT 0, date_created integer  DEFAULT 0, PRIMARY KEY (id));
-- Create index "files_id" to table: "files"
CREATE UNIQUE INDEX files_id ON files (id);
-- Create index "files_path" to table: "files"
CREATE UNIQUE INDEX files_path ON files (path);
