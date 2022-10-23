package source

import (
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	messengermocks "github.com/theobitoproject/kankuro/pkg/messenger/mocks"
	"github.com/theobitoproject/kankuro/pkg/protocol"
	sourcemocks "github.com/theobitoproject/kankuro/pkg/source/mocks"
)

var _ = Describe("SourceRunner", func() {
	var sourceRunner SourceRunner
	var err error

	var mockSource *sourcemocks.MockSource

	var mockMessenger *messengermocks.MockMessenger
	var mockPrivateMessenger *messengermocks.MockPrivateMessenger
	var mockConfigParser *messengermocks.MockConfigParser

	var mockCtrl gomock.Controller

	BeforeEach(func() {
		mockCtrl = *gomock.NewController(GinkgoT())

		mockSource = sourcemocks.NewMockSource(&mockCtrl)

		mockMessenger = messengermocks.NewMockMessenger(&mockCtrl)
		mockPrivateMessenger = messengermocks.NewMockPrivateMessenger(&mockCtrl)
		mockConfigParser = messengermocks.NewMockConfigParser(&mockCtrl)

		sourceRunner = NewSourceRunner(
			mockSource,
			mockMessenger,
			mockPrivateMessenger,
			mockConfigParser,
			nil,
		)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("When running start", func() {
		JustBeforeEach(func() {
			err = sourceRunner.Start()
		})

		Context("when getting main command fails", func() {
			BeforeEach(func() {
				mockConfigParser.
					EXPECT().
					GetMainCommand().
					Return(protocol.Cmd(""), fmt.Errorf("error getting main command")).
					Times(1)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})

		Context("when getting main command succeeds", func() {
			Context("when main command is empty", func() {
				BeforeEach(func() {
					mockConfigParser.
						EXPECT().
						GetMainCommand().
						Return(protocol.Cmd(""), nil).
						Times(1)
				})

				It("should return an error", func() {
					Expect(err).ToNot(BeNil())
				})
			})

			Context("when main command is NOT empty", func() {
				Context("with spec command", func() {
					BeforeEach(func() {
						mockConfigParser.
							EXPECT().
							GetMainCommand().
							Return(protocol.CmdSpec, nil).
							Times(1)
					})

					Context("when source spec method fails", func() {
						BeforeEach(func() {
							mockSource.
								EXPECT().
								Spec(mockMessenger, mockConfigParser).
								Return(
									&protocol.ConnectorSpecification{},
									fmt.Errorf("error running source spec"),
								).
								Times(1)

							mockMessenger.
								EXPECT().
								// TODO: expect WriteLog to be called with a string
								// for the second parameter
								WriteLog(protocol.LogLevelError, gomock.Any()).
								Return(nil).
								Times(1)
						})

						It("should return an error", func() {
							Expect(err).ToNot(BeNil())
						})
					})

					Context("when source spec method succeeds", func() {
						var mockConnectorSpecification *protocol.ConnectorSpecification

						BeforeEach(func() {
							mockConnectorSpecification = &protocol.ConnectorSpecification{
								DocumentationURL: "some url",
								ChangeLogURL:     "some url",
							}

							mockSource.
								EXPECT().
								Spec(mockMessenger, mockConfigParser).
								Return(
									mockConnectorSpecification,
									nil,
								).
								Times(1)
						})

						Context("when writing spec fails", func() {
							BeforeEach(func() {
								mockPrivateMessenger.
									EXPECT().
									WriteSpec(mockConnectorSpecification).
									Return(fmt.Errorf("error writing spec")).
									Times(1)

								mockMessenger.
									EXPECT().
									// TODO: expect WriteLog to be called with a string
									// for the second parameter
									WriteLog(protocol.LogLevelError, gomock.Any()).
									Return(nil).
									Times(1)
							})

							It("should return an error", func() {
								Expect(err).ToNot(BeNil())
							})
						})

						Context("when writing spec succeeds", func() {
							BeforeEach(func() {
								mockPrivateMessenger.
									EXPECT().
									WriteSpec(mockConnectorSpecification).
									Return(nil).
									Times(1)
							})

							It("should NOT return an error", func() {
								Expect(err).To(BeNil())
							})
						})
					})
				})

				Context("with check command", func() {
					BeforeEach(func() {
						mockConfigParser.
							EXPECT().
							GetMainCommand().
							Return(protocol.CmdCheck, nil).
							Times(1)
					})

					Context("when source check method fails", func() {
						BeforeEach(func() {
							mockSource.
								EXPECT().
								Check(mockMessenger, mockConfigParser).
								Return(fmt.Errorf("error running source check")).
								Times(1)

							mockMessenger.
								EXPECT().
								// TODO: expect WriteLog to be called with a string
								// for the second parameter
								WriteLog(protocol.LogLevelError, gomock.Any()).
								Return(nil).
								Times(1)
						})

						Context("when writing connection stat fails", func() {
							BeforeEach(func() {
								mockPrivateMessenger.
									EXPECT().
									WriteConnectionStat(protocol.CheckStatusFailed).
									Return(fmt.Errorf("error writing connection stat")).
									Times(1)

								mockMessenger.
									EXPECT().
									// TODO: expect WriteLog to be called with a string
									// for the second parameter
									WriteLog(protocol.LogLevelError, gomock.Any()).
									Return(nil).
									Times(1)
							})

							It("should return an error", func() {
								Expect(err).ToNot(BeNil())
							})
						})

						Context("when writing connection stat succeeds", func() {
							BeforeEach(func() {
								mockPrivateMessenger.
									EXPECT().
									WriteConnectionStat(protocol.CheckStatusFailed).
									Return(nil).
									Times(1)
							})

							It("should NOT return an error", func() {
								Expect(err).To(BeNil())
							})
						})
					})

					Context("when source check method succeeds", func() {
						BeforeEach(func() {
							mockSource.
								EXPECT().
								Check(mockMessenger, mockConfigParser).
								Return(nil).
								Times(1)
						})

						Context("when writing connection stat fails", func() {
							BeforeEach(func() {
								mockPrivateMessenger.
									EXPECT().
									WriteConnectionStat(protocol.CheckStatusSuccess).
									Return(fmt.Errorf("error writing connection stat")).
									Times(1)

								mockMessenger.
									EXPECT().
									// TODO: expect WriteLog to be called with a string
									// for the second parameter
									WriteLog(protocol.LogLevelError, gomock.Any()).
									Return(nil).
									Times(1)
							})

							It("should return an error", func() {
								Expect(err).ToNot(BeNil())
							})
						})

						Context("when writing connection stat succeeds", func() {
							BeforeEach(func() {
								mockPrivateMessenger.
									EXPECT().
									WriteConnectionStat(protocol.CheckStatusSuccess).
									Return(nil).
									Times(1)
							})

							It("should NOT return an error", func() {
								Expect(err).To(BeNil())
							})
						})
					})
				})

				Context("with discover command", func() {
					BeforeEach(func() {
						mockConfigParser.
							EXPECT().
							GetMainCommand().
							Return(protocol.CmdDiscover, nil).
							Times(1)
					})

					Context("when source discover method fails", func() {
						BeforeEach(func() {
							mockSource.
								EXPECT().
								Discover(mockMessenger, mockConfigParser).
								Return(
									&protocol.Catalog{},
									fmt.Errorf("error running source discover"),
								).
								Times(1)

							mockMessenger.
								EXPECT().
								// TODO: expect WriteLog to be called with a string
								// for the second parameter
								WriteLog(protocol.LogLevelError, gomock.Any()).
								Return(nil).
								Times(1)
						})

						It("should return an error", func() {
							Expect(err).ToNot(BeNil())
						})
					})

					Context("when source discover method succeeds", func() {
						var mockCatalog *protocol.Catalog

						BeforeEach(func() {
							mockCatalog = &protocol.Catalog{
								Streams: []protocol.Stream{
									{
										Name:      "some stream name",
										Namespace: "some namespace",
									},
								},
							}

							mockSource.
								EXPECT().
								Discover(mockMessenger, mockConfigParser).
								Return(
									mockCatalog,
									nil,
								).
								Times(1)
						})

						Context("when writing catalog fails", func() {
							BeforeEach(func() {
								mockPrivateMessenger.
									EXPECT().
									WriteCatalog(mockCatalog).
									Return(fmt.Errorf("error writing catalog")).
									Times(1)

								mockMessenger.
									EXPECT().
									// TODO: expect WriteLog to be called with a string
									// for the second parameter
									WriteLog(protocol.LogLevelError, gomock.Any()).
									Return(nil).
									Times(1)
							})

							It("should return an error", func() {
								Expect(err).ToNot(BeNil())
							})
						})

						Context("when writing catalog succeeds", func() {
							BeforeEach(func() {
								mockPrivateMessenger.
									EXPECT().
									WriteCatalog(mockCatalog).
									Return(nil).
									Times(1)
							})

							It("should NOT return an error", func() {
								Expect(err).To(BeNil())
							})
						})
					})
				})

				Context("with read command", func() {
					var configuredCatalog protocol.ConfiguredCatalog

					BeforeEach(func() {
						mockConfigParser.
							EXPECT().
							GetMainCommand().
							Return(protocol.CmdRead, nil).
							Times(1)
					})

					Context("when unmarshaling catalog path fails", func() {
						BeforeEach(func() {
							mockConfigParser.
								EXPECT().
								UnmarshalCatalogPath(&configuredCatalog).
								Return(fmt.Errorf("error unmarshaling catalog path")).
								Times(1)

							mockMessenger.
								EXPECT().
								// TODO: expect WriteLog to be called with a string
								// for the second parameter
								WriteLog(protocol.LogLevelError, gomock.Any()).
								Return(nil).
								Times(1)
						})

						It("should return an error", func() {
							Expect(err).ToNot(BeNil())
						})
					})

					Context("when unmarshaling catalog path succeeds", func() {
						BeforeEach(func() {
							mockConfigParser.
								EXPECT().
								UnmarshalCatalogPath(&configuredCatalog).
								DoAndReturn(func(v interface{}) error {
									// TODO: solve "argument v is overwritten before first use"
									v = protocol.ConfiguredCatalog{
										Streams: []protocol.ConfiguredStream{
											{
												Stream: protocol.Stream{
													Name: "some name",
												},
												SyncMode: "some sync",
											},
										},
									}
									return nil
								}).
								Times(1)
						})

						Context("when source read method fails", func() {
							BeforeEach(func() {
								mockSource.
									EXPECT().
									Read(&configuredCatalog, mockMessenger, mockConfigParser, nil).
									Return(fmt.Errorf("error running source read")).
									Times(1)

								mockMessenger.
									EXPECT().
									// TODO: expect WriteLog to be called with a string
									// for the second parameter
									WriteLog(protocol.LogLevelError, gomock.Any()).
									Return(nil).
									Times(1)
							})

							It("should return an error", func() {
								Expect(err).ToNot(BeNil())
							})
						})

						Context("when source read method succeeds", func() {
							BeforeEach(func() {
								mockSource.
									EXPECT().
									Read(&configuredCatalog, mockMessenger, mockConfigParser, nil).
									Return(nil).
									Times(1)
							})

							It("should NOT return an error", func() {
								Expect(err).To(BeNil())
							})
						})
					})
				})

				Context("when invalid command", func() {
					BeforeEach(func() {
						mockConfigParser.
							EXPECT().
							GetMainCommand().
							Return(protocol.Cmd("hello"), nil).
							Times(1)
					})

					It("should return an error", func() {
						Expect(err).ToNot(BeNil())
					})
				})
			})
		})
	})
})
