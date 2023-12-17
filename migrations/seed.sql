INSERT INTO "users" (id, name, created_at)
VALUES ('1', 'Alice', '2023-01-01 10:00:00');
INSERT INTO "users" (id, name, created_at)
VALUES ('2', 'Bob', '2023-01-02 11:00:00');
INSERT INTO "users" (id, name, created_at)
VALUES ('3', 'Charlie', '2023-01-03 12:00:00');
INSERT INTO "foods" (id, name, description, category, image)
VALUES (
    1,
    'Pizza',
    'Cheese and tomato',
    'bakso2',
    'pizza.jpg'
  );
INSERT INTO "foods" (id, name, description, category, image)
VALUES (
    2,
    'Burger',
    'Beef patty with lettuce and tomato',
    'bakso',
    'burger.jpg'
  );
INSERT INTO "foods" (id, name, description, category, image)
VALUES (
    3,
    'Sushi',
    'Fresh salmon and rice',
    'bakso3',
    'sushi.jpg'
  );
INSERT INTO "detection_history" (id, result, created_at, user_id)
VALUES (
    1,
    '{"data":"bakso"}',
    '2023-01-01 13:00:00',
    '1'
  );
INSERT INTO "detection_history" (id, result, created_at, user_id)
VALUES (
    2,
    '{"data":"bakso"}',
    '2023-01-02 14:00:00',
    '2'
  );
INSERT INTO "detection_history" (id, result, created_at, user_id)
VALUES (
    3,
    '{"data":"bakso"}',
    '2023-01-03 15:00:00',
    '3'
  );
INSERT INTO "shops" (
    id,
    name,
    location,
    gmaps_link,
    latitude,
    longitude,
    image
  )
VALUES (
    1,
    'Best Bites',
    '123 Main St',
    'http://maps.google.com/?q=123 Main St',
    40.7128,
    -74.006,
    'shop1.jpg'
  );
INSERT INTO "shops" (
    id,
    name,
    location,
    gmaps_link,
    latitude,
    longitude,
    image
  )
VALUES (
    2,
    'Tasty Treats',
    '456 Elm St',
    'http://maps.google.com/?q=456 Elm St',
    34.0522,
    -118.2437,
    'shop2.jpg'
  );
INSERT INTO "shops" (
    id,
    name,
    location,
    gmaps_link,
    latitude,
    longitude,
    image
  )
VALUES (
    3,
    'Yummy Corner',
    '789 Oak St',
    'http://maps.google.com/?q=789 Oak St',
    41.8781,
    -87.6298,
    'shop3.jpg'
  );
INSERT INTO "shop_foods" (id, shop_id, food_id, price)
VALUES (1, 1, 1, 10000);
INSERT INTO "shop_foods" (id, shop_id, food_id, price)
VALUES (2, 1, 2, 20000);
INSERT INTO "shop_foods" (id, shop_id, food_id, price)
VALUES (3, 2, 3, 30000);