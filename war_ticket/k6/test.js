import http from 'k6/http';
import { check, sleep } from 'k6';
import { randomString } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

export const options = {
    vus: 100,
    duration: '5s',
};

export default function () {
    const url = 'http://localhost:8080/api/orders';

    const payload = JSON.stringify({
        name: randomString(10),
        ticket_ids: [
            {
                "ticket_id": 3,
                "quantity":1
            },
            {
                "ticket_id": 4,
                "quantity":2
            }
        ]
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
            'X-API-KEY': '818e1618-db13-4a74-a99b-32bfc55a5c49'
        },
    };

    const res = http.post(url, payload, params);

    check(res, {
        'status is 201': (r) => r.status === 201,
        'status error': (r) => r.status >= 400,
    });

    sleep(1);
}