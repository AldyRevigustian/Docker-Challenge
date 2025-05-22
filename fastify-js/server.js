const fastify = require('fastify')({ logger: true });

fastify.get('/', async (request, reply) => {
  return { msg: 'Hello, from Fastify!' };
});

const start = async () => {
  try {
    await fastify.listen({ port: 8080, host: '0.0.0.0' });
    console.log('Server running on http://localhost:8080');
  } catch (err) {
    fastify.log.error(err);
    process.exit(1);
  }
};

start();
