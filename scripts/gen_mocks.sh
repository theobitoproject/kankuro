#!/bin/bash

# protocol
mockgen -package=mocks -source=pkg/messenger/channel_hub.go -destination=pkg/messenger/mocks/channel_hub_mock.go
mockgen -package=mocks -source=pkg/messenger/config_parser.go -destination=pkg/messenger/mocks/config_parser_mock.go
mockgen -package=mocks -source=pkg/messenger/messenger.go -destination=pkg/messenger/mocks/messenger_mock.go
mockgen -package=mocks -source=pkg/messenger/private_messenger.go -destination=pkg/messenger/mocks/private_messenger_mock.go

# source
mockgen -package=mocks -source=pkg/source/source.go -destination=pkg/source/mocks/source_mock.go
