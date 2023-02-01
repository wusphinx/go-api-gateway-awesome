((
  config = pipy.solve('config.js'),

) => pipy()

  .export('main', {
    __route: undefined,
    __isTLS: false,
  })

  .listen(config.listen)
  .link('inbound-http')

  .pipeline('inbound-http')
  .demuxHTTP().to(
    $ => $.chain(config.plugins)
  )
)()
