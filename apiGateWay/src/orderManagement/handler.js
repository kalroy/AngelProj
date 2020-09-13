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
  connection, channel, resultQ, requestID,
}) {
  return new Promise((resolve, reject) => {
    channel.consume(resultQ, async (msg) => {
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
      noAck: true,
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
  const resultQ = 'cart_result_queue';
  //await channel.assertQueue(resultQ, { durable: false });
  // start consuming messages
  const { status, error } = await consume({
    connection, channel, resultQ, requestID,
  });

  return { status, error };
}

const cartHandler = async (clientID, productID, quantity, jobType) => {
  const conn = await amqp.connect('amqp://guest:guest@rabbitmq:5672/');
  const pubChannel = await conn.createChannel();

  const pubQ = 'cart_manager_queue';
  await pubChannel.assertQueue(pubQ, { durable: false });

  const requestID = v4();

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

  const { status, error } = await listenForResults(requestID);
  return { status, error };
};

module.exports = cartHandler;
