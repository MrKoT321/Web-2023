CREATE TABLE post (
    `post_id` INT NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(255) NOT NULL,
    `subtitle` VARCHAR(255) NOT NULL,
    `img_modifier` VARCHAR(255) NOT NULL,
    `author` VARCHAR(255) NOT NULL,
    `author_modifier` VARCHAR(255) NOT NULL,
    `publish_date` VARCHAR(255) NOT NULL,
<<<<<<< HEAD
    `content` TEXT NOT NULL,
=======
>>>>>>> 2933629f94622529b12de65018a845ea35fee356
    `featured` TINYINT(1),
    PRIMARY KEY (`post_id`)
) ENGINE = InnoDB 
CHARACTER SET = utf8mb4 
COLLATE utf8mb4_unicode_ci
;