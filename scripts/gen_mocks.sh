#!/bin/bash

# protocol
mockgen -package=mocks -source=pkg/messenger/channel_hub.go -destination=pkg/messenger/mocks/channel_hub_mock.go
mockgen -package=mocks -source=pkg/messenger/config_parser.go -destination=pkg/messenger/mocks/config_parser_mock.go
mockgen -package=mocks -source=pkg/messenger/message_writer.go -destination=pkg/messenger/mocks/message_writer_mock.go
mockgen -package=mocks -source=pkg/messenger/private_message_writer.go -destination=pkg/messenger/mocks/private_message_writer_mock.go

# source
mockgen -package=mocks -source=pkg/source/source.go -destination=pkg/source/mocks/source_mock.go
