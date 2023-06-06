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

-- Creo la tabla de Chats
CREATE TABLE IF NOT EXISTS `chats` (
    id INT PRIMARY KEY AUTO_INCREMENT,
    channel_id INT NOT NULL,
    user_id INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    KEY `channel_id` (`channel_id`),
    KEY `user_id` (`user_id`),
    CONSTRAINT `chats_fk_1` FOREIGN KEY (`channel_id`) REFERENCES `channels` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `chats_fk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Creo tabla de relaciones entre usuarios y chat
CREATE TABLE IF NOT EXISTS `user_chat` (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    chat_id INT NOT NULL,
    UNIQUE KEY `user_chat_unique` (`user_id`, `chat_id`),
    KEY `user_id` (`user_id`),
    KEY `chat_id` (`chat_id`),
    CONSTRAINT `user_chat_fk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `user_chat_fk_2` FOREIGN KEY (`chat_id`) REFERENCES `chats` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Creo la tabla de Messages
CREATE TABLE IF NOT EXISTS `messages` (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    chat_id INT NOT NULL,
    message VARCHAR(256) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    KEY `chat_id` (`chat_id`),
    KEY `user_id` (`user_id`),
    CONSTRAINT `messages_fk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `messages_fk_2` FOREIGN KEY (`chat_id`) REFERENCES `chats` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

