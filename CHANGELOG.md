# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.2.0] - 2022-06-11

### Added

- Support for [JSON file format](https://www.json.org/json-en.html)
- Support for [YAML file format](https://yaml.org/)
- Support for [TOML file format](https://toml.io/en/)
- Support for [WATSON file format](https://github.com/genkami/watson)

- SetResourceType function

- FileType enum

### Changed

- CreateXMLFile function is now CreateResourceFile and takes a FileType parameter
- DeleteResourceFile function is now DeleteResourceFile
- LoadValues function now takes a FileType parameter

## [1.1.0] - 2022-06-06

### Added

- Support for XML file format

- CreateXMLFile, DeleteXMLFile, LoadValues, NewString, NewStringArray, NewQuantityString, SetFewThreshold, GetString, GetStringArray, GetQuantityString functions