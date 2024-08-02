import http from 'k6/http';
import { check, sleep } from 'k6';
import { randomString } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

export const options = {
    vus: 500,
    duration: '5s',
};

export default function () {
    const url = 'http://localhost:8080/api/orders';

    const payload = JSON.stringify({
        name: randomString(10),
        ticket_ids: [
            {
                "ticket_id": 2,
                "quantity":1
            }
        ]
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
            'X-API-KEY': '8aba5f15-dfc0-4fbd-b85f-a471ba752203'
        },
    };

    const res = http.post(url, payload, params);

    check(res, {
        'status is 201': (r) => r.status === 201,
        'status error': (r) => r.status >= 400,
    });

    sleep(1);
}