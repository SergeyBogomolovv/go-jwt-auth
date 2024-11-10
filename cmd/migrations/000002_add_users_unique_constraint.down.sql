ALTER TABLE users
ADD CONSTRAINT unique_username_and_email UNIQUE(username, email);