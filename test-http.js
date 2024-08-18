import http from "k6/http";
import { check } from 'k6';

export const options = {
    iterations: 10000,
    VUs: 10,
};

export default () => {
    let data = { 
        Name: "test-name",
        Value: 1.23,
        Stock: 10,
        CreatedAt: "2024-08-18 14:02:56"
    };

    const response = http.post("http://localhost:7007/product", JSON.stringify(data), {
        headers: { 'Content-Type': 'application/json' },
      });

    check(response, {
      'status is OK': (r) => r.status === 201,
    });
};