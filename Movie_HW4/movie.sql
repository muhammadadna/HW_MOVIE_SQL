/* Jika diperlukan */
CREATE TABLE movie(
    id int unsigned NOT NULL AUTO_INCREMENT,
    title char(255) NOT NULL,
    slug char(255) UNIQUE NOT NULL,
    description text NOT NULL,
    duration int(5) NOT NULL,
    image char(255) NOT NULL,
    PRIMARY KEY(id)
);