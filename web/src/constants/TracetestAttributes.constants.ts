/**
 * Moved from JSON files to TS files, as JSONs are not working in composite projects.
 * @see {@link https://github.com/TypeStrong/ts-loader/issues/905}
 */
export default {
  name: {
    description:
      'The span name concisely identifies the work represented by the Span, for example, an RPC method name, a function name, or the name of a subtask or stage within a larger computation. The span name SHOULD be the most general string that identifies a (statistically) interesting class of Spans, rather than individual Span instances while still being human-readable.',
    note: '',
    tags: ['id'],
  },
  'quality-trace.span.duration': {
    description: 'Qualitytrace attribute that reflects the elapsed real time of the operation.',
    note: '',
    tags: ['ms', 'second', 'time'],
  },
  'quality-trace.span.type': {
    description:
      'Qualitytrace attribute based on the [OTel Trace Semantic Conventions](https://github.com/open-telemetry/opentelemetry-specification/tree/main/specification/trace/semantic_conventions)',
    note: '',
    tags: ['general', 'http', 'database', 'rpc', 'messaging', 'faas', 'exception'],
  },
};
