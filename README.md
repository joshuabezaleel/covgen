## Covenant Generator (Covgen)

Generate a Markdown file of Code of Conduct for your Go project.

## Installation

```
go get github.com/joshuabezaleel/covgen
```

## Usage

In the root of your Go project folder:

```
covgen --email <your_email_address> --version <1.4/2.0>
```

You can leave the version empty and it will default to the latest version which is [2.0](https://www.contributor-covenant.org/version/2/0/code_of_conduct).

The code of conduct in this repository was generated with covenant generator.

## Prior Art

This project is highly inspired by the kind efforts of simonv3's [covenant-generator](https://github.com/simonv3/covenant-generator), @bitandbang's [tweet](https://twitter.com/bitandbang/status/1082331715471925250), and swyx's [blogpost](https://www.swyx.io/writing/oss-repo-setup/).

## License

[MIT](LICENSE)