import http from 'k6/http';
import {check} from 'k6';

// export let options = {
//     vus: 10,  // number of virtual users
//     duration: '30s',  // duration of the test
// };

export default function () {
    let apiKey = __ENV.API_KEY;
    let res = http.post('https://prompt.safetorun.com/keep',
        JSON.stringify({"prompt": "test"}),
        {headers: {'Content-Type': 'application/json', "x-api-key": apiKey}}
    );

    check(res, {
        'status was 200': (r) => r.status == 200,

        'transaction time OK': (r) => r.timings.duration < 200,
    });
}