#!/bin/bash

# protocol
mockgen -package=mocks -source=protocol/config_parser.go -destination=protocol/mocks/config_parser_mock.go
mockgen -package=mocks -source=protocol/messenger.go -destination=protocol/mocks/messenger_mock.go
mockgen -package=mocks -source=protocol/private_messenger.go -destination=protocol/mocks/private_messenger_mock.go

# source
mockgen -package=mocks -source=source/source.go -destination=source/mocks/source_mock.go
