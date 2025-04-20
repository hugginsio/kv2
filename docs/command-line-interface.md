# Command Line Interface

The `kv2` command line interface (CLI) is the reference implementation of the API client library and provides the easiest way to interact with your `kv2` server instance. The CLI exposes all the same functionality of the API alongside some usability enhancements for usage in a terminal environment.

The CLI manual, generated from the help text, is available in the [cli-manual](cli-manual/kv2.md) directory.

## Installation

You can download the latest CLI executable for your platform from [the Releases page](https://github.com/hugginsio/kv2/releases/latest) or install it with Homebrew:

```sh
homebrew install hugginsio/tap/kv2
```

## Configuration

Currently, the `kv2` CLI uses a single environment variable for configuration:

- `KV2_SERVER_URL`, the URL of the server instance (for example, `https://kv2.yak-bebop.ts.net`)
