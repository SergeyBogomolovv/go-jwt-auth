CREATE OR REPLACE FUNCTION is_user_exists(u_username VARCHAR DEFAULT '', u_email VARCHAR DEFAULT '')
RETURNS BOOLEAN AS $$
BEGIN
  IF u_username = '' AND u_email != '' THEN
  	RETURN EXISTS(SELECT 1 FROM users WHERE email = u_email);
  ELSIF u_email = '' AND u_username != '' THEN
  	RETURN EXISTS(SELECT 1 FROM users WHERE username = u_username);
  ELSIF u_email != '' AND u_username != '' THEN
  	RETURN EXISTS(SELECT 1 FROM users WHERE username = u_username OR email = u_email);
  END IF;
  RETURN false;
END;
$$ LANGUAGE plpgsql;