import http from 'k6/http';
import {check} from 'k6';

export let options = {
    vus: 10,  // number of virtual users
    duration: '5s',  // duration of the test
};

export default function () {
    let apiKey = __ENV.DEFENDER_API_KEY;
    let URL = __ENV.URL + "/moat";
    let res = http.post(URL,
        JSON.stringify({"prompt": "test"}),
        {headers: {'Content-Type': 'application/json', "x-api-key": apiKey}}
    );

    check(res, {
        'status was 200': (r) => r.status == 200,

        'transaction time OK': (r) => r.timings.duration < 200,
    });
}