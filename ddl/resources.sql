DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS resources;
DROP TABLE IF EXISTS types;
DROP TABLE IF EXISTS menu_items;
DROP TABLE IF EXISTS templates;

CREATE TABLE users (
  id VARCHAR(64) NOT NULL PRIMARY KEY,
  email VARCHAR(255) NOT NULL,
  password_hash VARCHAR(256),
  password_salt VARCHAR(256),
  password_reset_token VARCHAR(256),
  first_name VARCHAR(50),
  last_name VARCHAR(50),
  is_blocked BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  password_last_changed_at TIMESTAMP WITH TIME ZONE,
  last_login_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE resources (
	path VARCHAR(1024) NOT NULL PRIMARY KEY,
  redirect_path VARCHAR(1024),  -- Only used for redirects
	version INT NOT NULL DEFAULT 1,
	is_current_version BOOLEAN DEFAULT FALSE,
	type VARCHAR(64),
	template VARCHAR(64),
	title VARCHAR(255),
	short_title VARCHAR(255),
	long_title VARCHAR(255),
	description VARCHAR(255),
	content TEXT,
	data JSON,
	cache_ttl INT,
	is_published BOOLEAN,
	published_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	unpublished_at TIMESTAMP WITH TIME ZONE,
	created_at TIMESTAMP WITH TIME ZONE,
	created_by_user VARCHAR(64),
	last_modified_at TIMESTAMP WITH TIME ZONE,
	last_modified_by_user VARCHAR(64),
	deleted_at TIMESTAMP WITH TIME ZONE,
	deleted_by_user VARCHAR(64)
);

CREATE TABLE types (
  id VARCHAR(64),
  name VARCHAR(64),
  fields JSON,
  title_pattern VARCHAR(255),
  path_pattern VARCHAR(1024),
  default_template VARCHAR(64),
  default_is_published BOOLEAN,
  default_cache_ttl INT
);


-- SELECT * FROM menu_items;
--
--                   id                  |              parent_id               | position |        query         | label_pattern |  path  | id_attribute_pattern | class_attribute_pattern | title_attribute_pattern
-- --------------------------------------+--------------------------------------+----------+----------------------+---------------+--------+----------------------+-------------------------+-------------------------
--  1f7c37da-b5a8-4a6d-881b-ae91630b2157 |                                      |    32768 |                      | Horses        | horses |                      |                         |
--  11804e84-98ac-4f18-b516-28e89a6d1160 | 1f7c37da-b5a8-4a6d-881b-ae91630b2157 |    32768 | type=horse&year=2015 | [[name]]      |        |                      |                         |
--

CREATE TABLE menu_items (
  id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
  parent_id uuid,
  position INT NOT NULL DEFAULT 32768,
  query VARCHAR(1024),         -- Instead of providing a resource, you can provide a query from which to build dynamic menu items
  label_pattern VARCHAR(255),  -- What goes inside the <a>
  path VARCHAR(1024),          -- Which resource does this menu point at?
  id_attribute_pattern VARCHAR(255),    -- Example: 'menu-[menu]-[position]'
  class_attribute_pattern VARCHAR(255), -- Example: '[menu]'
  title_attribute_pattern VARCHAR(255)  -- What goes on the title="" attribute
);

CREATE TABLE templates (
  id VARCHAR(64) NOT NULL PRIMARY KEY,
  content TEXT,
  content_type VARCHAR(50),
  is_root BOOLEAN NOT NULL DEFAULT FALSE -- Templates that start with <html> are considered "root" templates
);

