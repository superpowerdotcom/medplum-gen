medplum-gen
===========

Repo containing generated code for the Medplum API.

* **Uses**
  * Medplum's OpenAPI spec: https://api.medplum.com/openapi.json
  * OpenAPI Generator: https://openapi-generator.tech/

## Usage

* `make help` to list available targets.
* `make generate` to run all generators.
* `make generate/go` to run just the Go generator.
* `make clean` to remove all generated code artifacts.

## Structure

OpenAPI spec files are saved under `./specs/*`. We do this so that we can easily
rollback to a different version if needed.

Current version of the spec file generated code is using is saved under `./version`.

## Flow

When Medplum updates their OpenAPI spec, download the `openapi.json` file and
save it under `./specs/0_0_0.json` (replace 0's with actual `openapi` version
specified in the openapi.json file).

Update `./version` to reflect the new version.

Run `make generate` to generate new code.