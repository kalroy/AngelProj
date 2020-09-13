'use strict';

require('dotenv').config({ path: '.local.env', silent: true });

const Good = require('@hapi/good');
const Hapi = require('@hapi/hapi');
const Inert = require('@hapi/inert');
const Vision = require('@hapi/vision');
const HapiSwagger = require('hapi-swagger');

const OrderManagement = require('./orderManagement');

const pkg = require('../package.json');

const init = async () => {
  const server = Hapi.server({
    address: process.env.LISTEN_ADDR,
    port: process.env.PORT,
    routes: { cors: true },
  });

  await server.register([{
    plugin: Good,
    options: {
      ops: { interval: 1000 },
      reporters: {
        console: [{
          module: '@hapi/good-squeeze',
          name: 'Squeeze',
          args: [{ log: '*', response: { exclude: 'health' } }],
        }, {
          module: 'good-console',
        }, 'stdout'],
      },
    },
  }, {
    plugin: HapiSwagger,
    options: {
      info: {
        title: 'Order Management API Gateway',
        version: pkg.version,
      },
      reuseDefinitions: false,
      definitionPrefix: 'useLabel',
    },
  }, {
    plugin: Inert,
  }, {
    plugin: Vision,
  }]);

  await server.register([
    {
      plugin: OrderManagement,
      options: {
      },
    },
  ]);

  server.log(['OMApi'], 'Registered Plugins');

  await server.start();
  server.log(['OMApi'], `Server running on ${server.info.uri}`);
};

process.on('unhandledRejection', (err) => {
  console.log(err); // eslint-disable-line no-console
  process.exit(1);
});

init();
