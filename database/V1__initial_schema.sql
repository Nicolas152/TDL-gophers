-- Crear la base de datos
CREATE DATABASE IF NOT EXISTS gochat;

-- Usar la base de datos
USE gochat;

DROP TABLE IF EXISTS `messages`;
DROP TABLE IF EXISTS `chats`;
DROP TABLE IF EXISTS `channels`;
DROP TABLE IF EXISTS `user_workspace`;
DROP TABLE IF EXISTS `workspaces`;
DROP TABLE IF EXISTS `users`;

-- Creo la tabla de Users
CREATE TABLE IF NOT EXISTS `users` (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(128) NOT NULL,
    email VARCHAR(128) NOT NULL,
    password VARCHAR(128) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Creo la tabla de Workspaces
CREATE TABLE IF NOT EXISTS `workspaces` (
    id INT PRIMARY KEY AUTO_INCREMENT,
    workflow_key VARCHAR(11) NOT NULL,
    name VARCHAR(128) NOT NULL,
    password VARCHAR(128) NULL DEFAULT NULL,
    creator INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `name` (`name`),
    CONSTRAINT `workspaces_fk_1` FOREIGN KEY (`creator`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Creo tabla de relaciones entre usuarios y workspaces
CREATE TABLE IF NOT EXISTS `user_workspace` (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    workspace_id INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY `user_workspace_unique` (`user_id`, `workspace_id`),
    KEY `user_id` (`user_id`),
    KEY `workspace_id` (`workspace_id`),
    CONSTRAINT `user_workspace_fk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `user_workspace_fk_2` FOREIGN KEY (`workspace_id`) REFERENCES `workspaces` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Creo la tabla de Channels
CREATE TABLE IF NOT EXISTS `channels` (
    id INT PRIMARY KEY AUTO_INCREMENT,
    workspace_id INT NOT NULL,
    name VARCHAR(128) NOT NULL,
    password VARCHAR(128) NULL DEFAULT NULL,
    creator INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `workspace_name_unique` (`workspace_id`, `name`),
    KEY `workspace_id` (`workspace_id`),
    CONSTRAINT `channels_fk_1` FOREIGN KEY (`workspace_id`) REFERENCES `workspaces` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `channels_fk_2` FOREIGN KEY (`creator`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


-- Creo tabla de relaciones entre usuarios y channels
CREATE TABLE IF NOT EXISTS `user_channels` (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    channel_id INT NOT NULL,
    UNIQUE KEY `user_channel_unique` (`user_id`, `channel_id`),
    KEY `user_id` (`user_id`),
    KEY `channel_id` (`channel_id`),
    CONSTRAINT `user_channel_fk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `user_channel_fk_2` FOREIGN KEY (`channel_id`) REFERENCES `channels` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Creo la tabla de Messages de un channel
CREATE TABLE IF NOT EXISTS `channel_messages` (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    channel_id INT NOT NULL,
    message VARCHAR(256) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    KEY `channel_id` (`channel_id`),
    KEY `user_id` (`user_id`),
    CONSTRAINT `messages_fk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `messages_fk_2` FOREIGN KEY (`channel_id`) REFERENCES `channels` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


-- Creo la tabla de DMS
CREATE TABLE IF NOT EXISTS `dms` (
    id INT PRIMARY KEY AUTO_INCREMENT,
    workspace_id INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    KEY `workspace_id` (`workspace_id`),
    CONSTRAINT `dm_fk_1` FOREIGN KEY (`workspace_id`) REFERENCES `workspaces` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Creo la tabla de Messages de un dm
CREATE TABLE IF NOT EXISTS `dm_messages` (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    dm_id INT NOT NULL,
    message VARCHAR(256) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    KEY `user_id` (`user_id`),
    KEY `dm_id` (`dm_id`),
    CONSTRAINT `dm_messages_fk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `dm_messages_fk_2` FOREIGN KEY (`dm_id`) REFERENCES `dms` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Creo tabla de relaciones entre usuarios y dms
CREATE TABLE IF NOT EXISTS `user_dms` (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    dm_id INT NOT NULL,
    UNIQUE KEY `user_dm_unique` (`user_id`, `dm_id`),
    KEY `user_id` (`user_id`),
    KEY `dm_id` (`dm_id`),
    CONSTRAINT `user_dm_fk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `user_dm_fk_2` FOREIGN KEY (`dm_id`) REFERENCES `dms` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;