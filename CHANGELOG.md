# Changelog

## [Unreleased]
### Added
- Initial release of the CHANGELOG.md file to document changes made in each release.
- Added comments and documentation to the code to improve readability and maintainability.
- Refactored the code to follow best practices and design patterns.
- Added unit tests to ensure the code is working as expected and to catch any potential bugs.
- Used a linter to enforce coding standards and improve code quality.
- Added more test cases to the existing CI pipeline to cover edge cases and improve test coverage.
- Implemented a code coverage tool to measure the effectiveness of the tests.
- Added a deployment step to the CI pipeline to automatically deploy the application to a staging or production environment.
- Improved the `README.md` file to provide more detailed instructions on how to use the application and its features.
- Created a `CONTRIBUTING.md` file to guide new contributors on how to contribute to the project.
- Created a `CHANGELOG.md` file to document the changes made in each release.
- Added more examples and use cases to the documentation to help users understand how to use the application effectively.

## [2.8.3] - 2023-08-15
### Added
- Added support for new providers: Blackbox AI, Duckduckgo, Groq, KoboldAI, Ollama, OpenAI, and Phind.
- Added support for image generation using BlackBoxAi.
- Added new flags for generating and executing shell commands, generating code, and generating images from text.
- Added support for logging conversations to a file.
- Added support for setting a preprompt.
- Added support for setting a custom OpenAI API endpoint URL.
- Added support for setting various parameters such as temperature, top_p, and max response length.
- Added support for updating the program using the `-u` flag.
- Added support for viewing the changelog using the `-cl` flag.

### Changed
- Refactored the code to improve readability and maintainability.
- Updated the `README.md` file to provide more detailed instructions on how to use the application and its features.
- Updated the usage examples to include the new flag parsing function and its usage.

### Fixed
- Fixed various bugs and issues reported by users.
