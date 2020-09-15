'use strict';

// const Boom = require('@hapi/boom');
const { v4 } = require('uuid');

const Joi = require('joi');
const { cartHandler, listenForResults } = require('./handler');

exports.plugin = {
  name: 'actions',
  version: '1.0.0',
  async register(server, options) {
    server.route({
      method: 'POST',
      path: '/addToCart',
      options: {
        description: 'Add Items to User Cart',
        notes: [
          'Add a specific item to user cart',
        ],
        tags: ['api'],
        validate: {
          payload: Joi.object({
            clientID: Joi.string().uuid().required(),
            productID: Joi.string().uuid().required(),
            quantity: Joi.number().required(),
          }),
        },
      },
      handler: 
      async (request, h) => {
        const { clientID, productID, quantity } = request.payload;
        const requestID = v4();

        const promise = listenForResults(requestID);

        cartHandler(clientID, productID, quantity, 'addToCart', requestID);
        return promise.then(({status, error}) => h.response({ status, error }).code(200));
      },
    });

    server.route({
      method: 'POST',
      path: '/checkOutCart',
      options: {
        description: 'Checkout Items in Client Cart',
        notes: [
          'Checking out action on an User Cart',
        ],
        tags: ['api'],
        validate: {
          payload: Joi.object({
            clientID: Joi.string().uuid().required(),
          }),
        },
      },
      handler: async (request, h) => {
        const { clientID } = request.payload;
        const requestID = v4();

        const promise = listenForResults(requestID);
        cartHandler(clientID, null, 0, 'checkOut', requestID);
        return promise.then(({ status, error }) => h.response({ status, error }).code(200));
      },
    });
  },
};
