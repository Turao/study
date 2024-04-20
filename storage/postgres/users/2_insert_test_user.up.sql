INSERT INTO users VALUES(
  gen_random_uuid(),
  gen_random_uuid(),
  0,
  'john.doe@example.com',
  'john',
  'doe',
  'tenancy/test',
  NOW() 
);