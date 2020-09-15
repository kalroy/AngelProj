'use strict';

const { v4 } = require('uuid');
const amqp = require('amqplib');

const getAddToCartJob = (clientID, productID, quantity, requestID) => ({
  requestID,
  clientID,
  productID,
  quantity,
  jobType: 'addToCart',
});

const getCheckOutJob = (clientID, requestID) => ({
  requestID,
  clientID,
  jobType: 'checkOut',
});

function consume({
  connection, channel, queue, requestID,
}) {
  return new Promise((resolve, reject) => {
    channel.consume(queue, async (msg) => {
      if (msg.properties.correlationId === requestID) {
        try {
          const r = JSON.parse(msg.content.toString());
          const status = r.isSuccess ? 'Success' : 'Failed';
          const error = r.error ? r.error : '';
          await connection.close();
          resolve({ status, error });
        } catch (e) {
          reject(e);
        }
      }
    }, {
      noAck: false,
    });

    // handle connection closed
    // connection.on('close', (err) => reject(err));

    // // handle errors
    // connection.on('error', (err) => reject(err));
  });
}

async function listenForResults(requestID) {
  const connection = await amqp.connect('amqp://guest:guest@rabbitmq:5672/');
  const channel = await connection.createChannel();
  const exchange = 'cart_result_exchange';
  channel.assertExchange(exchange, 'fanout', { durable: true });
  const queue = `queue-${v4()}`;
  channel.assertQueue(queue, { exclusive: true, durable: false });
  channel.bindQueue(queue, exchange, '');
  // start consuming messages
  return consume({
    connection, channel, queue, requestID,
  });
}

const cartHandler = async (clientID, productID, quantity, jobType, requestID) => {
  const conn = await amqp.connect('amqp://guest:guest@rabbitmq:5672/');
  const pubChannel = await conn.createChannel();

  const pubQ = 'cart_manager_queue';
  await pubChannel.assertQueue(pubQ, { durable: false });

  let job;
  switch (jobType) {
    case 'addToCart':
      job = getAddToCartJob(clientID, productID, quantity, requestID);
      break;
    case 'checkOut':
      job = getCheckOutJob(clientID, requestID);
      break;
    default:
      return false;
  }

  pubChannel.sendToQueue(pubQ, Buffer.from(JSON.stringify(job)), {
    correlationId: requestID,
  });

  console.log('Sent to queue')

  return; 
};

module.exports = {
  cartHandler,
  listenForResults,
};
