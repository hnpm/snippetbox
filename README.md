# snippetbox


#### Database script
    CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

    CREATE USER 'web'@'localhost';
    GRANT SELECT, INSERT ON snippetbox.* TO 'web'@'localhost';
    ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';

    USE snippetbox;

    CREATE TABLE snippets (
        id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
        title VARCHAR(100) NOT NULL,
        content TEXT NOT NULL,
        created DATETIME NOT NULL,
        expires DATETIME NOT NULL
    );

    CREATE INDEX idx_snippets_created ON snippets(created);

    CREATE TABLE users (
        id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL,
        hashed_password CHAR(60) NOT NULL,
        created DATETIME NOT NULL
    );

    ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);

#### Database for testing
    CREATE DATABASE test_snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

    CREATE USER 'test_web'@'localhost';
    GRANT CREATE, DROP, ALTER, INDEX, SELECT, INSERT, UPDATE, DELETE ON test_snippetbox.* TO 'test_web'@'localhost';
    ALTER USER 'test_web'@'localhost' IDENTIFIED BY 'pass';

#### Generating a Self-Signed TLS Certificate (macOS)
    go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost --ca
