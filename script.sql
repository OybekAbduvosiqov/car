SELECT
    o.id,
    o.car_id,

    COALESCE(c.id::varchar, ''),
    COALESCE(c.state_number, ''),
    COALESCE(c.model, ''),

    COALESCE(cl.first_name || ' ' || cl.last_name, ''),

    o.client_id,
    o.total_price,
    o.paid_price,
    o.day_count,
    o.give_km,
    o.receive_km,
    o.status,
    o.created_at,
    o.updated_at
FROM "order" AS o
JOIN car AS c ON c.id = o.car_id
JOIN client AS cl ON cl.id = o.client_id
WHERE o.id = '0abbe136-a04b-48db-98cb-3fa5cd11722f'


SELECT
    SUM(d.price),
    cl.first_name || ' ' || cl.last_name AS full_name
FROM debtors AS d
JOIN client AS cl ON cl.id = d.client_id
GROUP BY full_name


SELECT
    i.id,
    i.name,
    SUM(o.paid_price / 100 * c.investor_percentage)
FROM "order" AS o
JOIN car AS c ON c.id = o.car_id
JOIN investor AS i ON i.id = c.investor_id
GROUP BY i.id, i.name
;
