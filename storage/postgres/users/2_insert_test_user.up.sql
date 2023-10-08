INSERT INTO users VALUES(
  gen_random_uuid(),
  gen_random_uuid(),
  0,
  'john',
  'doe',
  'john.doe@example.com',
  'tenancy/test',
  NOW() 
);