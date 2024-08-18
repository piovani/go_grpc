import grpc from 'k6/net/grpc';
import { check } from 'k6';

export const options = {
    iterations: 10000,
    VUs: 10,
};

const client = new grpc.Client();
client.load(['ui/grpc'], 'product.proto');

export default () => {
    client.connect('localhost:7008', {
        plaintext: true
    });
  
    const data = { 
        Name: "test-name",
        Value: 1.23,
        Stock: 10,
        CreatedAt: "2024-08-18 14:02:56"
    };
    const response = client.invoke('Product/CreateProduct', data);

    check(response, {
      'status is OK': (r) => {
        if (r && r.status === grpc.StatusOK) {
            return true
        }

        console.error("Error: ", r.error.message)
            return false 
        }       
    });
  

    client.close();
  };