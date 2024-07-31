DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_views WHERE viewname = 'ticket_event_view') THEN
        EXECUTE '
        CREATE VIEW ticket_event_view AS
        SELECT
            e.id AS event_id,
            e.name AS event_name,
            e.location AS event_location,
            e.created_at AS event_created_at,
            e.updated_at AS event_updated_at,
            json_agg(json_build_object(
                    ''id'', t.id,
                    ''name'', t.name,
                    ''stock'', t.stock,
                    ''price'', t.price,
                    ''created_at'', t.created_at,
                    ''updated_at'', t.updated_at
                     )) AS tickets
        FROM
            events e
                JOIN
            ticket_events te ON e.id = te.event_id
                JOIN
            tickets t ON te.ticket_id = t.id
        GROUP BY
            e.id, e.name, e.location, e.created_at, e.updated_at;
        ';
END IF;
END $$;