import http from 'k6/http';
import {check} from 'k6';

export let options = {
    vus: 10,  // number of virtual users
    duration: '5s',  // duration of the test
};

export default function () {
    let apiKey = __ENV.DEFENDER_API_KEY;
    let URL = __ENV.URL + "/wall";

    let options = [
        {
            "body": {"prompt": "test", "pii_detection": true},
            "response_time": 500
        },
        {
            "body": {"prompt": "test", "pii_detection": false},
            "response_time": 400
        },
    ];

    let option = Math.floor(Math.random() * options.length) + 1;

    let res = http.post(URL,
        JSON.stringify(option[i].body),
        {headers: {'Content-Type': 'application/json', "x-api-key": apiKey}}
    );

    check(res, {
        'status was 200': (r) => r.status == 200,
        'transaction time OK': (r) => r.timings.duration < option[i].response_time,
    });
}