medplum-gen
===========

Repo containing generated code for the Medplum API.

* **Uses**
  * Medplum's OpenAPI spec: https://api.medplum.com/openapi.json
  * OpenAPI Generator: https://openapi-generator.tech/

## Usage

* `make help` to list available targets.
* `make generate` to run all generators.
* `make clean` to remove all generated code artifacts.

## Structure

OpenAPI spec files are saved under `./specs/*`. We do this so that we can easily
rollback to a different version if needed.

Current version of the spec file generated code is using is saved under `./version`.

## Flow

When Medplum updates their OpenAPI spec, download the `openapi.json` file and
save it under `./specs/0_0_0.json` (replace 0's with actual version specified in
in `openapi` field).

Update `./version` to reflect the new version.

Run `make generate` to generate new code.

## `1_0_5.json` || 06.26.2024

* Removed undefined definitions in spec file

```bash
#/components/schemas/AccessPolicyIpAccessRule
#/components/schemas/AccessPolicyResource
#/components/schemas/AgentChannel
#/components/schemas/AgentSetting
#/components/schemas/BulkDataExportDeleted
#/components/schemas/BulkDataExportError
#/components/schemas/BulkDataExportOutput
#/components/schemas/ProjectLink
#/components/schemas/ProjectMembershipAccess
#/components/schemas/ProjectSetting
#/components/schemas/ProjectSite
#/components/schemas/Resource
#/components/schemas/UserConfigurationMenu
#/components/schemas/UserConfigurationOption
#/components/schemas/UserConfigurationSearch
```
