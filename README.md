# feedback

Echo your feedback into `/dev/null` and have it actually go somewhere!

Simply delete `/dev/null` and replace it with a named pipe and then run
`feedback <pipe>`. Any message of the format `<1 - 10> <text>` will be printed
to stdout.
