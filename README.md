# Encrypt Data

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ralvescosta_encrypt-data&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=ralvescosta_encrypt-data)

## Usage

- You need to download the binary with work better for you, chouse one of the binary in the [Releases](https://github.com/ralvescosta/encrypt-data/releases).

- Update the **input/data.json** with your data. 

  The "public_key" keyword needed to be a valid RSA Public Key.

  Example:

  ```json
  {
    "public_key": "-----BEGIN PUBLIC KEY-----\nsomething\nsomething\nsomething\nsomething\nsomething\n-----END PUBLIC KEY-----"
  }
  ```

  The "payload" keyword will be encrypted, make ensure you have written the correct values.

- Execute the binary

- The encrypted data will be written inside the **output** folder