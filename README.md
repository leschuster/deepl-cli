# DeepL-CLI (unofficial)

<p>
<a href="https://github.com/leschuster/deepl-cli/blob/main/LICENSE"><img src="https://img.shields.io/github/license/leschuster/deepl-cli" alt="license badge" /></a>
<a href="https://github.com/leschuster/deepl-cli/releases"><img src="https://img.shields.io/github/v/release/leschuster/deepl-cli" alt="release badge" /></a>
</p>

This application leverages the power of [DeepL](https://www.deepl.com) to provide seamless translations directly from your terminal.

![Demo Gif](./.github/assets/demo.gif)

## 🚀 Getting Started

1. **Requirements**:

When you first launch the application, you will be asked to enter an [API Key](https://www.deepl.com/en/pro#developer). DeepL provides both free and paid API keys.

This key will be saved in your system's keyring. If you're on macOS or Windows, everything should work out of the box. On Linux, make sure that [GNOME Keyring](https://wiki.gnome.org/Projects/GnomeKeyring) is installed.

2. **Installation**:

Make sure that you have at least Golang 1.23.0 installed. Then run the following command:

```bash
go install github.com/leschuster/deepl-cli/cmd/deepl-cli@latest
```

Alternatively, you can download the appropriate binary for your system in the release section.

3. **Usage**:

Run `deepl-cli` in your terminal.

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
